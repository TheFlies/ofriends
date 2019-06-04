package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/robfig/cron"

	"github.com/TheFlies/ofriends/internal/app/activity"
	"github.com/TheFlies/ofriends/internal/app/actvisitassoc"
	activityhandler "github.com/TheFlies/ofriends/internal/app/api/handler/activity"
	actvisitassochandler "github.com/TheFlies/ofriends/internal/app/api/handler/actvisitassoc"
	customerhandler "github.com/TheFlies/ofriends/internal/app/api/handler/customer"
	cusvisitassochandler "github.com/TheFlies/ofriends/internal/app/api/handler/cusvisitassoc"
	gifthandler "github.com/TheFlies/ofriends/internal/app/api/handler/gift"
	giftassociatehandler "github.com/TheFlies/ofriends/internal/app/api/handler/giftassociate"
	indexhandler "github.com/TheFlies/ofriends/internal/app/api/handler/index"
	"github.com/TheFlies/ofriends/internal/app/api/handler/login"
	userhandler "github.com/TheFlies/ofriends/internal/app/api/handler/user"
	visithandler "github.com/TheFlies/ofriends/internal/app/api/handler/visit"
	timelinehandler "github.com/TheFlies/ofriends/internal/app/api/handler/timeline"
	"github.com/TheFlies/ofriends/internal/app/customer"
	"github.com/TheFlies/ofriends/internal/app/cusvisitassoc"
	"github.com/TheFlies/ofriends/internal/app/timeline"
	"github.com/TheFlies/ofriends/internal/app/db"
	"github.com/TheFlies/ofriends/internal/app/dbauth"
	"github.com/TheFlies/ofriends/internal/app/email"
	"github.com/TheFlies/ofriends/internal/app/gift"
	"github.com/TheFlies/ofriends/internal/app/giftassociate"
	"github.com/TheFlies/ofriends/internal/app/ldapauth"
	"github.com/TheFlies/ofriends/internal/app/user"
	"github.com/TheFlies/ofriends/internal/app/visit"
	envconfig "github.com/TheFlies/ofriends/internal/pkg/config/env"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
	"github.com/TheFlies/ofriends/internal/pkg/health"
	"github.com/TheFlies/ofriends/internal/pkg/middleware"
)

type (
	// InfraConns holds infrastructure services connections like MongoDB, Redis, Kafka,...
	InfraConns struct {
		Databases db.Connections
	}

	middlewareFunc = func(http.HandlerFunc) http.HandlerFunc
	route          struct {
		path        string
		method      string
		handler     http.HandlerFunc
		middlewares []middlewareFunc
	}
	CronConfig struct {
		CronTAB string `envconfig:"CRON_TAB" default:"30 9 * * *"` // 9h30 daily will send mail
	}
)

const (
	get       = http.MethodGet
	post      = http.MethodPost
	put       = http.MethodPut
	delete    = http.MethodDelete
	roleAdmin = 3
	roleUser  = 2
	none      = 1
)

// Init init all handlers
func Init(conns *InfraConns) (http.Handler, error) {
	logger := glog.New()
	var conf CronConfig
	var customerRepo customer.Repository
	var userRepo user.UserRepository
	var giftRepo gift.Repository
	var visitRepo visit.Repository
	var actRepo activity.Repository
	var giftAssociateRepo giftassociate.Repository
	var actVisitAssocRepo actvisitassoc.Repository
	var cusVisitAssocRepo cusvisitassoc.Repository

	switch conns.Databases.Type {
	case db.TypeMongoDB:
		customerRepo = customer.NewMongoRepository(conns.Databases.MongoDB)
		userRepo = user.NewUserMongoRepository(conns.Databases.MongoDB)
		giftRepo = gift.NewMongoRepository(conns.Databases.MongoDB)
		visitRepo = visit.NewMongoRepository(conns.Databases.MongoDB)
		actRepo = activity.NewMongoRepository(conns.Databases.MongoDB)
		customerRepo = customer.NewMongoRepository(conns.Databases.MongoDB)
		userRepo = user.NewUserMongoRepository(conns.Databases.MongoDB)
		giftAssociateRepo = giftassociate.NewMongoRepository(conns.Databases.MongoDB)
		actVisitAssocRepo = actvisitassoc.NewMongoRepository(conns.Databases.MongoDB)
		cusVisitAssocRepo = cusvisitassoc.NewMongoRepository(conns.Databases.MongoDB)
	default:
		return nil, fmt.Errorf("database type not supported: %s", conns.Databases.Type)
	}
	envconfig.Load(&conf)
	actVisitAssocLogger := logger.WithField("package", "actvisitassociate")
	actVisitAssocSrv := actvisitassoc.NewService(actVisitAssocRepo, actVisitAssocLogger)
	actVisitAssocHandler := actvisitassochandler.New(actVisitAssocSrv, actVisitAssocLogger)

	giftAssociateLogger := logger.WithField("package", "giftassociate")
	giftAssociateSrv := giftassociate.NewService(giftAssociateRepo, giftAssociateLogger)
	giftAssociateHandler := giftassociatehandler.New(giftAssociateSrv, giftAssociateLogger)

	cusVisitAssocLogger := logger.WithField("package", "cusvisitassociate")
	cusVisitAssocSrv := cusvisitassoc.NewService(cusVisitAssocRepo, giftAssociateRepo, cusVisitAssocLogger)
	cusVisitAssocHandler := cusvisitassochandler.New(cusVisitAssocSrv, cusVisitAssocLogger)

	actLogger := logger.WithField("package", "activity")
	actSrv := activity.NewService(actRepo, actVisitAssocSrv, actLogger)
	actHandler := activityhandler.New(actSrv, actLogger)

	customerLogger := logger.WithField("package", "customer")
	customerSrv := customer.NewService(customerRepo, cusVisitAssocSrv, customerLogger)
	customerHandler := customerhandler.New(customerSrv, customerLogger)

	giftLogger := logger.WithField("package", "gift")
	giftSrv := gift.NewService(giftRepo, giftAssociateSrv, giftLogger)
	giftHandler := gifthandler.New(giftSrv, giftLogger)

	visitLogger := logger.WithField("package", "visit")
	visitSrv := visit.NewService(visitRepo, cusVisitAssocRepo, actVisitAssocRepo, visitLogger)
	emailLogger := logger.WithField("package", "email")
	emailService, err := email.NewSendMailService(visitRepo, customerRepo, cusVisitAssocRepo, emailLogger)
	if err != nil {
		return nil, err
	}
	visitHandler := visithandler.New(visitSrv, emailService, visitLogger)
	visitSrv := visit.NewService(visitRepo, cusVisitAssocSrv, actVisitAssocSrv, visitLogger)
	visitHandler := visithandler.New(visitSrv, visitLogger)

	indexWebHandler := indexhandler.New()

	userLogger := logger.WithField("package", "user")
	dbLoginLogger := logger.WithField("package", "dbauth")
	userService := user.NewUserService(userRepo, userLogger)
	ldapLoginService := ldapauth.New(userService)
	localLoginService := dbauth.NewDBAuthentication(dbLoginLogger, userService)

	timelineLogger := logger.WithField("package", "timeline")
	timelineService := timeline.NewService(visitRepo, customerRepo, cusVisitAssocRepo, giftAssociateRepo,
										  giftRepo, actVisitAssocRepo, actRepo, timelineLogger)
	timelineHandler := timelinehandler.New(timelineService, timelineLogger)

	loginHandler := login.NewLoginHandler(localLoginService, ldapLoginService, userLogger)
	userHandler := userhandler.NewUserHandler(userService, localLoginService, ldapLoginService)

	routes := []route{
		// infra
		{
			path:    "/readiness",
			method:  get,
			handler: health.Readiness().ServeHTTP,
		},
		// web
		{
			path:    "/",
			method:  get,
			handler: indexWebHandler.Index,
		},
		// Customer services
		{
			path:        "/customers/{id:[a-z0-9-\\-]+}",
			method:      get,
			handler:     customerHandler.Get,
			middlewares: []middlewareFunc{middleware.Authentication, middleware.Authorization(roleUser)},
		},
		{
			path:        "/customers",
			method:      post,
			handler:     customerHandler.Create,
			middlewares: []middlewareFunc{middleware.Authentication, middleware.Authorization(roleUser)},
		},
		{
			path:        "/customers",
			method:      put,
			handler:     customerHandler.Update,
			middlewares: []middlewareFunc{middleware.Authentication, middleware.Authorization(roleUser)},
		},
		{
			path:        "/customers",
			method:      get,
			handler:     customerHandler.GetAll,
			middlewares: []middlewareFunc{middleware.Authentication, middleware.Authorization(roleUser)},
		},
		{
			path:        "/customers/{id:[a-z0-9-\\-]+}",
			method:      delete,
			handler:     customerHandler.Delete,
			middlewares: []middlewareFunc{middleware.Authentication, middleware.Authorization(roleUser)},
		},

		// Visit services
		{
			path:        "/visits/{id:[a-z0-9-\\-]+}",
			method:      get,
			handler:     visitHandler.Get,
			middlewares: []middlewareFunc{middleware.Authentication, middleware.Authorization(roleUser)},
		},
		{
			path:        "/visits",
			method:      post,
			handler:     visitHandler.Create,
			middlewares: []middlewareFunc{middleware.Authentication, middleware.Authorization(roleUser)},
		},
		{
			path:        "/visits",
			method:      put,
			handler:     visitHandler.Update,
			middlewares: []middlewareFunc{middleware.Authentication, middleware.Authorization(roleUser)},
		},
		{
			path:        "/visits",
			method:      get,
			handler:     visitHandler.GetAll,
			middlewares: []middlewareFunc{middleware.Authentication, middleware.Authorization(roleUser)},
		},
		{
			path:        "/visits/{id:[a-z0-9-\\-]+}",
			method:      delete,
			handler:     visitHandler.Delete,
			middlewares: []middlewareFunc{middleware.Authentication, middleware.Authorization(roleUser)},
		},
		{
			path:        "/visits/customer/{id:[a-z0-9-\\-]+}",
			method:      get,
			handler:     visitHandler.GetByCustomerID,
			middlewares: []middlewareFunc{middleware.Authentication, middleware.Authorization(roleUser)},
		},
		{
			path:        "/visits/email/{visitID:[a-z0-9-\\-]+}",
			method:      post,
			handler:     visitHandler.SendMail,
			middlewares: []middlewareFunc{middleware.Authentication, middleware.Authorization(roleUser)},
		},
		// Activity services
		{
			path:        "/activities/{id:[a-z0-9-\\-]+}",
			method:      get,
			handler:     actHandler.Get,
			middlewares: []middlewareFunc{middleware.Authentication, middleware.Authorization(roleUser)},
		},
		{
			path:        "/activities",
			method:      post,
			handler:     actHandler.Create,
			middlewares: []middlewareFunc{middleware.Authentication, middleware.Authorization(roleUser)},
		},
		{
			path:        "/activities",
			method:      put,
			handler:     actHandler.Update,
			middlewares: []middlewareFunc{middleware.Authentication, middleware.Authorization(roleUser)},
		},
		{
			path:        "/activities",
			method:      get,
			handler:     actHandler.GetAll,
			middlewares: []middlewareFunc{middleware.Authentication, middleware.Authorization(roleUser)},
		},
		{
			path:        "/activities/{id:[a-z0-9-\\-]+}",
			method:      delete,
			handler:     actHandler.Delete,
			middlewares: []middlewareFunc{middleware.Authentication, middleware.Authorization(roleUser)},
		},
		// Gift services
		{
			path:        "/gifts/{id:[a-z0-9-\\-]+}",
			method:      get,
			handler:     giftHandler.Get,
			middlewares: []middlewareFunc{middleware.Authentication, middleware.Authorization(roleUser)},
		},

		{
			path:        "/gifts",
			method:      get,
			handler:     giftHandler.GetAll,
			middlewares: []middlewareFunc{middleware.Authentication, middleware.Authorization(roleUser)},
		},

		{
			path:        "/gifts",
			method:      post,
			handler:     giftHandler.Create,
			middlewares: []middlewareFunc{middleware.Authentication, middleware.Authorization(roleUser)},
		},

		{
			path:        "/gifts",
			method:      put,
			handler:     giftHandler.Update,
			middlewares: []middlewareFunc{middleware.Authentication, middleware.Authorization(roleUser)},
		},

		{
			path:        "/gifts/{id:[a-z0-9-\\-]+}",
			method:      delete,
			handler:     giftHandler.Delete,
			middlewares: []middlewareFunc{middleware.Authentication, middleware.Authorization(roleUser)},
		},
		{
			path:    "/login",
			method:  post,
			handler: loginHandler.Authenticate,
		},
		{
			path:        "/register",
			method:      post,
			handler:     userHandler.Register,
			middlewares: []middlewareFunc{middleware.Authentication, middleware.Authorization(roleAdmin)},
		},
		{
			path:        "/getme",
			method:      get,
			handler:     userHandler.GetMe,
			middlewares: []middlewareFunc{middleware.Authentication, middleware.Authorization(roleUser)},
		},
		{
			path:        "/users",
			method:      get,
			handler:     userHandler.FindAll,
			middlewares: []middlewareFunc{middleware.Authentication, middleware.Authorization(roleUser)},
		},
		{
			path:        "/users",
			method:      put,
			handler:     userHandler.Update,
			middlewares: []middlewareFunc{middleware.Authentication, middleware.Authorization(roleUser)},
		},
		{
			path:        "/users/{username}",
			method:      get,
			handler:     userHandler.GetUser,
			middlewares: []middlewareFunc{middleware.Authentication, middleware.Authorization(roleUser)},
		},
		{
			path:        "/users/{username}",
			method:      post,
			handler:     userHandler.SetPassword,
			middlewares: []middlewareFunc{middleware.Authentication, middleware.Authorization(roleUser)},
		},
		{
			path:        "/users/{username}",
			method:      put,
			handler:     userHandler.ChangePassword,
			middlewares: []middlewareFunc{middleware.Authentication, middleware.Authorization(roleUser)},
		},

		// Gift Associate services
		{
			path:    "/giftassociates/{id:[a-z0-9-\\-]+}",
			method:  get,
			handler: giftAssociateHandler.Get,
		},

		{
			path:    "/giftassociates",
			method:  get,
			handler: giftAssociateHandler.GetAll,
		},

		{
			path:    "/giftassociates",
			method:  post,
			handler: giftAssociateHandler.Create,
		},

		{
			path:    "/giftassociates",
			method:  put,
			handler: giftAssociateHandler.Update,
		},

		{
			path:    "/giftassociates/{id:[a-z0-9-\\-]+}",
			method:  delete,
			handler: giftAssociateHandler.Delete,
		},

		{
			path:    "/giftassociates",
			method:  delete,
			handler: giftAssociateHandler.DeleteByCusVisitAssocID,
		},
		// Activity Associate services
		{
			path:    "/actvisitassocs/{id:[a-z0-9-\\-]+}",
			method:  get,
			handler: actVisitAssocHandler.Get,
		},

		{
			path:    "/actvisitassocs",
			method:  get,
			handler: actVisitAssocHandler.GetActVisitAssocs,
		},

		{
			path:    "/actvisitassocs",
			method:  post,
			handler: actVisitAssocHandler.Create,
		},

		{
			path:    "/actvisitassocs",
			method:  put,
			handler: actVisitAssocHandler.Update,
		},

		{
			path:    "/actvisitassocs/{id:[a-z0-9-\\-]+}",
			method:  delete,
			handler: actVisitAssocHandler.Delete,
		},

		{
			path:    "/actvisitassocs",
			method:  delete,
			handler: actVisitAssocHandler.DeleteByVisitID,
		},
		// Customer visit Associate services
		{
			path:    "/cusvisitassocs/{id:[a-z0-9-\\-]+}",
			method:  get,
			handler: cusVisitAssocHandler.Get,
		},

		{
			path:    "/cusvisitassocs",
			method:  get,
			handler: cusVisitAssocHandler.GetAll,
		},

		{
			path:    "/cusvisitassocs",
			method:  post,
			handler: cusVisitAssocHandler.Create,
		},

		{
			path:    "/cusvisitassocs",
			method:  put,
			handler: cusVisitAssocHandler.Update,
		},

		{
			path:    "/cusvisitassocs/{id:[a-z0-9-\\-]+}",
			method:  delete,
			handler: cusVisitAssocHandler.Delete,
		},

		{
			path:    "/cusvisitassocs",
			method:  delete,
			handler: cusVisitAssocHandler.DeleteByVisitID,
		},
		{
			path:    "/timeline?daytime={dayTime:[0-9]+}",
			method:  get,
			handler: timelineHandler.FindTimelineByDay,
		},
	}

	loggingMW := middleware.Logging(logger.WithField("package", "middleware"))
	r := mux.NewRouter()
	r.Use(middleware.EnableCORS)
	r.Use(middleware.RequestID)
	r.Use(middleware.StatusResponseWriter)
	r.Use(loggingMW)
	r.Use(handlers.CompressHandler)
	// Add version
	subRouter := r.PathPrefix("/v1").Subrouter()

	for _, rt := range routes {
		h := rt.handler
		for i := len(rt.middlewares) - 1; i >= 0; i-- {
			h = rt.middlewares[i](h)
		}
		subRouter.Path(rt.path).Methods(rt.method).HandlerFunc(h)
	}

	// even not found, return index so that VueJS does its job
	r.NotFoundHandler = middleware.RequestID(loggingMW(http.HandlerFunc(indexWebHandler.Index)))

	// static resources
	static := []struct {
		prefix string
		dir    string
	}{
		{
			prefix: "/",
			dir:    "web/",
		},
	}
	for _, s := range static {
		h := http.StripPrefix(s.prefix, http.FileServer(http.Dir(s.dir)))
		r.PathPrefix(s.prefix).Handler(middleware.StaticCache(h, 3600*24))
	}

	scheduler := cron.New()
	//read about cron tab here https://crontab.guru/ and https://en.wikipedia.org/wiki/Cron
	_, err = scheduler.AddFunc(conf.CronTAB, func() {
		logger.Infof("sending email")
		err := emailService.Send(context.Background(), time.Now())
		if err != nil {
			logger.Errorf("in %v don't have visit to send email", time.Now())
		}
	})
	if err != nil {
		logger.Errorf("can't create cron job, the job is not run service still run %v", err)
		return r, nil
	}
	scheduler.Start()
	return r, nil
}

// Close close all underlying connections
func (c *InfraConns) Close() {
	c.Databases.Close()
}
