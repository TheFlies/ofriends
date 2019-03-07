package api

import (
	"fmt"
	"github.com/TheFlies/ofriends/internal/app/api/handler/login"
	"github.com/TheFlies/ofriends/internal/app/usercache"
	"net/http"

	"github.com/TheFlies/ofriends/internal/app/api/handler/friend"
	"github.com/TheFlies/ofriends/internal/app/api/handler/index"
	"github.com/TheFlies/ofriends/internal/app/db"
	"github.com/TheFlies/ofriends/internal/app/friend"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
	"github.com/TheFlies/ofriends/internal/pkg/health"
	"github.com/TheFlies/ofriends/internal/pkg/middleware"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
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
	var usercacheRepo usercache.UsercacheReponsitory
	switch conns.Databases.Type {
	case db.TypeMongoDB:
		friendRepo = friend.NewMongoRepository(conns.Databases.MongoDB)
		usercacheRepo = usercache.NewUsercachMongoReopnsitoty(conns.Databases.MongoDB)

	default:
		return nil, fmt.Errorf("database type not supported: %s", conns.Databases.Type)
	}

	friendLogger := logger.WithField("package", "friend")
	friendSrv := friend.NewService(friendRepo, friendLogger)
	friendHandler := friendhandler.New(friendSrv, friendLogger)
	indexWebHandler := indexhandler.New()
	// usercache
	usercachelogger := logger.WithField("package", "usercache")
	usercacheService := usercache.NewUserCacheService(usercacheRepo, usercachelogger)
	loginhandeler := login.NewLoginHandeler(usercacheService, usercachelogger)
	routes := []route{
		// infra
		{
			path:    "/readiness",
			method:  get,
			handler: health.Readiness().ServeHTTP,
		},
		// services
		{
			path:    "/api/v1/friend/{id:[a-z0-9-\\-]+}",
			method:  get,
			handler: friendHandler.Get,
		},

		// web
		{
			path:    "/",
			method:  get,
			handler: indexWebHandler.Index,
		},
		{
			path:    "/api/v1/login",
			method:  post,
			handler: loginhandeler.Authentication,
		},
		{
			path:    "/api/v1/get",
			method:  post,
			handler: loginhandeler.GetUser,
		},
	}

	loggingMW := middleware.Logging(logger.WithField("package", "middleware"))
	r := mux.NewRouter()
	r.Use(middleware.EnableCORS)
	r.Use(middleware.RequestID)
	r.Use(middleware.StatusResponseWriter)
	r.Use(loggingMW)
	r.Use(handlers.CompressHandler)

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
