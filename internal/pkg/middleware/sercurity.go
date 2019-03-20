package middleware

import (
	"github.com/TheFlies/ofriends/internal/pkg/glog"
	"net/http"
	"strings"

	"github.com/TheFlies/ofriends/internal/pkg/config/env"
	"github.com/TheFlies/ofriends/internal/pkg/jwt"
	"github.com/TheFlies/ofriends/internal/pkg/respond"
)

type (
	JWTConf struct {
		PrivateKey string `envconfig:"PRIVATE_KEY"`
	}
)

func Security(h http.Handler) http.Handler {
	var conf JWTConf
	envconfig.Load(&conf)
	logger := glog.New().WithField("package", "middleware")
	var ignorlist []string
	ignorlist = append(ignorlist, "/api/v1/login")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, a := range ignorlist {
			if a == r.RequestURI {
				h.ServeHTTP(w, r)
				return
			}
		}
		thebearertoken := r.Header.Get("Authorization")
		tokenarry := strings.Split(thebearertoken, " ")
		if len(tokenarry) != 2 {
			logger.Infof("the request isn't contain bearer token")
			respond.JSON(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
		JWTToken := strings.Split(thebearertoken, " ")[1]
		logger.Infof("checking token")
		JWTGeneration := jwt.NewTokenGeneration()
		notExpire := JWTGeneration.CheckExp(JWTToken)
		if notExpire && JWTGeneration.CheckValib(JWTToken) {
			payload := JWTGeneration.GetPayload(JWTToken)
			logger.Infof("the payload owner is %v", payload.UserFullname)
			h.ServeHTTP(w, r)
			return
		}

		respond.JSON(w, http.StatusUnauthorized, "Unauthorized")
		return
	})
}
