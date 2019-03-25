package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/TheFlies/ofriends/internal/app/api/handler/friend"
	"github.com/TheFlies/ofriends/internal/app/api/handler/gift"
	"github.com/TheFlies/ofriends/internal/app/api/handler/index"
	"github.com/TheFlies/ofriends/internal/app/api/handler/login"
	"github.com/TheFlies/ofriends/internal/app/db"
	"github.com/TheFlies/ofriends/internal/app/dbauth"
	"github.com/TheFlies/ofriends/internal/app/friend"
	"github.com/TheFlies/ofriends/internal/app/gift"
	"github.com/TheFlies/ofriends/internal/app/ldap"
	"github.com/TheFlies/ofriends/internal/app/user"
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
	switch conns.Databases.Type {
	case db.TypeMongoDB:
		friendRepo = friend.NewMongoRepository(conns.Databases.MongoDB)
		userRepo = user.NewUserMongoRepositoty(conns.Databases.MongoDB)
		giftRepo = gift.NewMongoRepository(conns.Databases.MongoDB)
	default:
		return nil, fmt.Errorf("database type not supported: %s", conns.Databases.Type)
	}

	friendLogger := logger.WithField("package", "friend")
	friendSrv := friend.NewService(friendRepo, friendLogger)
	friendHandler := friendhandler.New(friendSrv, friendLogger)

	giftLogger := logger.WithField("package", "gift")
	giftSrv := gift.NewService(giftRepo, giftLogger)
	giftHandler := gifthandler.New(giftSrv, giftLogger)

	indexWebHandler := indexhandler.New()

	userlogger := logger.WithField("package", "user")
	DBLoginLongger := logger.WithField("package", "dbauth")
	userService := user.NewUserService(userRepo, userlogger)
	LDAPLoginService := ldap.New(userService)
	DBLoginService := dbauth.NewDBAuthentication(DBLoginLongger, userService)
	athenList := []login.LoginService{LDAPLoginService, DBLoginService}
	loginHandeler := login.NewLoginHandeler(userService, athenList, userlogger)

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
		// services
		{
			path:    "friends/{id:[a-z0-9-\\-]+}",
			method:  get,
			handler: friendHandler.Get,
		},
		{
			path:    "/friends",
			method:  post,
			handler: friendHandler.Create,
		},
		{
			path:    "/friends/{id:[a-z0-9-\\-]+}",
			method:  put,
			handler: friendHandler.Update,
		},
		{
			path:    "/friends",
			method:  get,
			handler: friendHandler.GetAll,
		},
		{
			path:    "/friends{id:[a-z0-9-\\-]+}",
			method:  delete,
			handler: friendHandler.Delete,
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
			handler: loginHandeler.Authenticate,
		},
		{
			path:    "/api/v1/user/{username}",
			method:  get,
			handler: loginHandeler.GetUser,
		},
		{
			path:    "/api/v1/register",
			method:  post,
			handler: loginHandeler.Register,
		},
		{
			path:    "/api/v1/password/change",
			method:  post,
			handler: loginHandeler.ChangeLocalPassword,
		},
		{
			path:    "/api/v1/user/setlocalpassword",
			method:  post,
			handler: loginHandeler.SetLocalPassword,
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
