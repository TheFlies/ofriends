package api

import (
	"fmt"
	"github.com/TheFlies/ofriends/internal/app/email"
	"net/http"

	"github.com/TheFlies/ofriends/internal/app/actvisitassoc"
	"github.com/TheFlies/ofriends/internal/app/api/handler/login"
	"github.com/TheFlies/ofriends/internal/app/cusvisitassoc"
	"github.com/TheFlies/ofriends/internal/app/giftassociate"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/TheFlies/ofriends/internal/app/activity"
	activityhandler "github.com/TheFlies/ofriends/internal/app/api/handler/activity"
	actvisitassochandler "github.com/TheFlies/ofriends/internal/app/api/handler/actvisitassoc"
	customerhandler "github.com/TheFlies/ofriends/internal/app/api/handler/customer"
	cusvisitassochandler "github.com/TheFlies/ofriends/internal/app/api/handler/cusvisitassoc"
	gifthandler "github.com/TheFlies/ofriends/internal/app/api/handler/gift"
	giftassociatehandler "github.com/TheFlies/ofriends/internal/app/api/handler/giftassociate"
	indexhandler "github.com/TheFlies/ofriends/internal/app/api/handler/index"
	userhandler "github.com/TheFlies/ofriends/internal/app/api/handler/user"
	visithandler "github.com/TheFlies/ofriends/internal/app/api/handler/visit"
	"github.com/TheFlies/ofriends/internal/app/customer"
	"github.com/TheFlies/ofriends/internal/app/db"
	"github.com/TheFlies/ofriends/internal/app/dbauth"
	"github.com/TheFlies/ofriends/internal/app/gift"
	"github.com/TheFlies/ofriends/internal/app/ldapauth"
	"github.com/TheFlies/ofriends/internal/app/user"
	"github.com/TheFlies/ofriends/internal/app/visit"
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
	actSrv := activity.NewService(actRepo, actVisitAssocRepo, actLogger)
	actHandler := activityhandler.New(actSrv, actLogger)

	customerLogger := logger.WithField("package", "customer")
	customerSrv := customer.NewService(customerRepo, cusVisitAssocRepo, customerLogger)
	customerHandler := customerhandler.New(customerSrv, customerLogger)

	giftLogger := logger.WithField("package", "gift")
	giftSrv := gift.NewService(giftRepo, giftAssociateRepo, giftLogger)
	giftHandler := gifthandler.New(giftSrv, giftLogger)

	visitLogger := logger.WithField("package", "visit")
	visitSrv := visit.NewService(visitRepo, cusVisitAssocRepo, actVisitAssocRepo, visitLogger)
	visitHandler := visithandler.New(visitSrv, visitLogger)

	indexWebHandler := indexhandler.New()

	userLogger := logger.WithField("package", "user")
	dbLoginLogger := logger.WithField("package", "dbauth")
	userService := user.NewUserService(userRepo, userLogger)
	ldapLoginService := ldapauth.New(userService)
	localLoginService := dbauth.NewDBAuthentication(dbLoginLogger, userService)

	emailLogger := logger.WithField("package", "email")
	emailServie := email.NewSendMailService(visitRepo, customerRepo, cusVisitAssocRepo, emailLogger)

	loginHandler := login.NewLoginHandler(localLoginService, ldapLoginService, emailServie, userLogger)
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
			path:    "/email",
			method:  get,
			handler: loginHandler.SendMail,
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
			path:    "/giftassociates?cusvisitassocid={assignID:[a-z0-9-\\-]+}",
			method:  get,
			handler: giftAssociateHandler.GetByCusVisitAssocID,
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
			path:    "/giftassociates?cusvisitassocid={assignID:[a-z0-9-\\-]+}",
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
			path:    "/actvisitassocs?visitid={visitID:[a-z0-9-\\-]+}",
			method:  get,
			handler: actVisitAssocHandler.GetByVisitID,
		},

		{
			path:    "/actvisitassocs",
			method:  get,
			handler: actVisitAssocHandler.GetAll,
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
			path:    "/actvisitassocs?visitid={visitID:[a-z0-9-\\-]+}",
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
			path:    "/cusvisitassocs?visitid={visitID:[a-z0-9-\\-]+}",
			method:  get,
			handler: cusVisitAssocHandler.GetByVisitID,
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
			path:    "/cusvisitassocs?visitid={visitID:[a-z0-9-\\-]+}",
			method:  delete,
			handler: cusVisitAssocHandler.DeleteByVisitID,
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
	subrouter := r.PathPrefix("/v1").Subrouter()

	for _, rt := range routes {
		h := rt.handler
		for i := len(rt.middlewares) - 1; i >= 0; i-- {
			h = rt.middlewares[i](h)
		}
		subrouter.Path(rt.path).Methods(rt.method).HandlerFunc(h)
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

	return r, nil
}

// Close close all underlying connections
func (c *InfraConns) Close() {
	c.Databases.Close()
}
