package api

import (
	"fmt"
	"net/http"

	"github.com/TheFlies/ofriends/internal/app/api/handler/login"
	"github.com/TheFlies/ofriends/internal/app/giftassociate"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/TheFlies/ofriends/internal/app/activity"
	activityhandler "github.com/TheFlies/ofriends/internal/app/api/handler/activity"
	friendhandler "github.com/TheFlies/ofriends/internal/app/api/handler/friend"
	gifthandler "github.com/TheFlies/ofriends/internal/app/api/handler/gift"
	giftassociatehandler "github.com/TheFlies/ofriends/internal/app/api/handler/giftassociate"
	indexhandler "github.com/TheFlies/ofriends/internal/app/api/handler/index"
	userhandler "github.com/TheFlies/ofriends/internal/app/api/handler/user"
	visithandler "github.com/TheFlies/ofriends/internal/app/api/handler/visit"
	"github.com/TheFlies/ofriends/internal/app/db"
	"github.com/TheFlies/ofriends/internal/app/dbauth"
	"github.com/TheFlies/ofriends/internal/app/friend"
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
	get    = http.MethodGet
	post   = http.MethodPost
	put    = http.MethodPut
	delete = http.MethodDelete
)

// Init init all handlers
func Init(conns *InfraConns) (http.Handler, error) {
	logger := glog.New()
	var friendRepo friend.Repository
	var userRepo user.UserRepository
	var giftRepo gift.Repository
	var visitRepo visit.Repository
	var actRepo activity.Repository
	var giftAssociateRepo giftassociate.Repository

	switch conns.Databases.Type {
	case db.TypeMongoDB:
		friendRepo = friend.NewMongoRepository(conns.Databases.MongoDB)
		userRepo = user.NewUserMongoRepository(conns.Databases.MongoDB)
		giftRepo = gift.NewMongoRepository(conns.Databases.MongoDB)
		visitRepo = visit.NewMongoRepository(conns.Databases.MongoDB)
		actRepo = activity.NewMongoRepository(conns.Databases.MongoDB)
		friendRepo = friend.NewMongoRepository(conns.Databases.MongoDB)
		userRepo = user.NewUserMongoRepository(conns.Databases.MongoDB)
		giftAssociateRepo = giftassociate.NewMongoRepository(conns.Databases.MongoDB)
	default:
		return nil, fmt.Errorf("database type not supported: %s", conns.Databases.Type)
	}

	friendLogger := logger.WithField("package", "friend")
	friendSrv := friend.NewService(friendRepo, friendLogger)
	friendHandler := friendhandler.New(friendSrv, friendLogger)

	giftLogger := logger.WithField("package", "gift")
	giftSrv := gift.NewService(giftRepo, giftLogger)
	giftHandler := gifthandler.New(giftSrv, giftLogger)

	visitLogger := logger.WithField("package", "visit")
	visitSrv := visit.NewService(visitRepo, visitLogger)
	visitHandler := visithandler.New(visitSrv, visitLogger)

	actLogger := logger.WithField("package", "activity")
	actSrv := activity.NewService(actRepo, actLogger)
	actHandler := activityhandler.New(actSrv, actLogger)

	giftAssociateLogger := logger.WithField("package", "giftassociate")
	giftAssociateSrv := giftassociate.NewService(giftAssociateRepo, giftAssociateLogger)
	giftAssociateHandler := giftassociatehandler.New(giftAssociateSrv, giftAssociateLogger)

	indexWebHandler := indexhandler.New()

	userLogger := logger.WithField("package", "user")
	dbLoginLogger := logger.WithField("package", "dbauth")
	userService := user.NewUserService(userRepo, userLogger)
	ldapLoginService := ldapauth.New(userService)
	localLoginService := dbauth.NewDBAuthentication(dbLoginLogger, userService)

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
		// Friend services
		{
			path:    "/friends/{id:[a-z0-9-\\-]+}",
			method:  get,
			handler: friendHandler.Get,
		},
		{
			path:    "/friends",
			method:  post,
			handler: friendHandler.Create,
		},
		{
			path:    "/friends",
			method:  put,
			handler: friendHandler.Update,
		},
		{
			path:    "/friends",
			method:  get,
			handler: friendHandler.GetAll,
		},
		{
			path:    "/friends/{id:[a-z0-9-\\-]+}",
			method:  delete,
			handler: friendHandler.Delete,
		},

		// Visit services
		{
			path:    "/visits/{id:[a-z0-9-\\-]+}",
			method:  get,
			handler: visitHandler.Get,
		},
		{
			path:    "/visits",
			method:  post,
			handler: visitHandler.Create,
		},
		{
			path:    "/visits",
			method:  put,
			handler: visitHandler.Update,
		},
		{
			path:    "/visits",
			method:  get,
			handler: visitHandler.GetAll,
		},
		{
			path:    "/visits/{id:[a-z0-9-\\-]+}",
			method:  delete,
			handler: visitHandler.Delete,
		},
		{
			path:    "/visits/friend/{id:[a-z0-9-\\-]+}",
			method:  get,
			handler: visitHandler.GetByFriendID,
		},
		// Activity services
		{
			path:    "/activity/{id:[a-z0-9-\\-]+}",
			method:  get,
			handler: actHandler.Get,
		},
		{
			path:    "/activity",
			method:  post,
			handler: actHandler.Create,
		},
		{
			path:    "/activity",
			method:  put,
			handler: actHandler.Update,
		},
		{
			path:    "/activity",
			method:  get,
			handler: actHandler.GetAll,
		},
		{
			path:    "/activity/{id:[a-z0-9-\\-]+}",
			method:  delete,
			handler: actHandler.Delete,
		},
		{
			path:    "/activity/visit/{id:[a-z0-9-\\-]+}",
			method:  get,
			handler: actHandler.GetByVisitID,
		},
		// Gift services
		{
			path:    "/gifts/{id:[a-z0-9-\\-]+}",
			method:  get,
			handler: giftHandler.Get,
		},

		{
			path:    "/gifts",
			method:  get,
			handler: giftHandler.GetAll,
		},

		{
			path:    "/gifts",
			method:  post,
			handler: giftHandler.Create,
		},

		{
			path:    "/gifts",
			method:  put,
			handler: giftHandler.Update,
		},

		{
			path:    "/gifts/{id:[a-z0-9-\\-]+}",
			method:  delete,
			handler: giftHandler.Delete,
		},
		{
			path:    "/api/v1/login",
			method:  post,
			handler: loginHandler.Authenticate,
		},
		{
			path:    "/api/v1/register",
			method:  post,
			handler: userHandler.Register,
		}, {
			path:    "/api/v1/getme",
			method:  get,
			handler: userHandler.GetMe,
		},
		{
			path:    "/api/v1/user/{username}",
			method:  get,
			handler: userHandler.GetUser,
		},
		{
			path:    "/api/v1/user/{username}",
			method:  post,
			handler: userHandler.SetPassword,
		},
		{
			path:    "/api/v1/user/{username}",
			method:  put,
			handler: userHandler.ChangePassword,
		},
		{
			path:    "/api/v1/user/{username}",
			method:  put,
			handler: userHandler.Update,
		},
		// Gift Associate services
		{
			path:    "/gifts/associates/{id:[a-z0-9-\\-]+}",
			method:  get,
			handler: giftAssociateHandler.Get,
		},

		{
			path:    "/gifts/associates/visits/{visitID:[a-z0-9-\\-]+}",
			method:  get,
			handler: giftAssociateHandler.GetByVisitID,
		},

		{
			path:    "/gifts/associates",
			method:  get,
			handler: giftAssociateHandler.GetAll,
		},

		{
			path:    "/gifts/associates",
			method:  post,
			handler: giftAssociateHandler.Create,
		},

		{
			path:    "/gifts/associates",
			method:  put,
			handler: giftAssociateHandler.Update,
		},

		{
			path:    "/gifts/associates/{id:[a-z0-9-\\-]+}",
			method:  delete,
			handler: giftAssociateHandler.Delete,
		},
	}

	loggingMW := middleware.Logging(logger.WithField("package", "middleware"))
	r := mux.NewRouter()
	r.Use(middleware.EnableCORS)
	r.Use(middleware.RequestID)
	r.Use(middleware.StatusResponseWriter)
	r.Use(loggingMW)
	r.Use(handlers.CompressHandler)
	r.Use(middleware.Security)

	for _, rt := range routes {
		h := rt.handler
		for i := len(rt.middlewares) - 1; i >= 0; i-- {
			h = rt.middlewares[i](h)
		}
		r.Path(rt.path).Methods(rt.method).HandlerFunc(h)
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
