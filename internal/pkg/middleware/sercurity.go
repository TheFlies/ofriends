package middleware

import (
	"fmt"
	"github.com/TheFlies/ofriends/internal/pkg/config/env"
	"github.com/TheFlies/ofriends/internal/pkg/jwt"
	"github.com/TheFlies/ofriends/internal/pkg/respond"
	"github.com/TheFlies/ofriends/internal/pkg/utils"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
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
		fmt.Println(r.RequestURI)
		if utils.Iscontains(ignorlist, r.RequestURI) {
			h.ServeHTTP(w, r)
			return
		}
		fmt.Println(conf.PrivateKey)
		thebearertoken := r.Header.Get("Authorization")
		tokenarry := strings.Split(thebearertoken, " ")
		if len(tokenarry) != 2 {
			respond.JSON(w, http.StatusUnauthorized, "Can find your token pls adding bearetoken beore query")
			return
		}
		jwttoken := strings.Split(thebearertoken, " ")[1]
		logrus.Info("the token : %v\n", jwttoken)
		notExpire := jwt.CheckExp(jwttoken)
		if notExpire && jwt.Checkvalib(jwttoken) {
			payload := jwt.GetPayload(jwttoken)
			logrus.Info("the payload is %v\n", payload)
			h.ServeHTTP(w, r)
			return
		}
		respond.JSON(w, http.StatusUnauthorized, "Your token not valib or your session is expire plse login agin")
		return
	})
}