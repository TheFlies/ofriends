package middleware

import (
	"net/http"
	"strings"

	"github.com/TheFlies/ofriends/internal/pkg/config/env"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
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
	var whiteList []string
	whiteList = append(whiteList, "/api/v1/login")
	whiteList = append(whiteList, "/api/v1/register")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, a := range whiteList {
			if a == r.RequestURI {
				h.ServeHTTP(w, r)
				return
			}
		}
		theBearerToken := r.Header.Get("Authorization")
		tokenArray := strings.Split(theBearerToken, " ")
		if len(tokenArray) != 2 {
			logger.Infof("the request isn't contain bearer token")
			respond.JSON(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
		jwtToken := strings.Split(theBearerToken, " ")[1]
		logger.Infof("checking token")
		JWTGeneration := jwt.NewTokenGeneration()
		notExpire := JWTGeneration.CheckExp(jwtToken)
		if notExpire && JWTGeneration.CheckValib(jwtToken) {
			payload := JWTGeneration.GetPayload(jwtToken)
			logger.Infof("the payload owner is %v", payload.UserFullname)
			logger.Infof("the request is continue treatment")
			h.ServeHTTP(w, r)
			return
		}
		respond.JSON(w, http.StatusUnauthorized, "Unauthorized")
		return
	})
}
