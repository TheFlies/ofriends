package middleware

import (
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/TheFlies/ofriends/internal/pkg/config/env"
	"github.com/TheFlies/ofriends/internal/pkg/jwt"
	"github.com/TheFlies/ofriends/internal/pkg/respond"
)

type (
	JWTConf struct {
		PrivateKey string `envconfig:"PRIVATE_KEY"`
	}
)

func Sercurity(h http.Handler) http.Handler {
	var conf JWTConf
	envconfig.Load(&conf)
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
			respond.JSON(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
		jwttoken := strings.Split(thebearertoken, " ")[1]
		logrus.Info("the token : %v\n", jwttoken)
		notExpire := jwt.CheckExp(jwttoken)
		if notExpire && jwt.Checkvalib(jwttoken) {
			payload := jwt.GetPayload(jwttoken)
			logrus.Info("the payload is %v\n", payload.Usefullname)
			h.ServeHTTP(w, r)
			return
		}
		respond.JSON(w, http.StatusUnauthorized, "Unauthorized")
		return
	})
}
