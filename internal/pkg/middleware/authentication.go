package middleware

import (
	"fmt"
	"github.com/TheFlies/ofriends/internal/pkg/config/env"
	"github.com/TheFlies/ofriends/internal/pkg/myjwt"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type (
	JWTConf struct {
		PrivateKey string `envconfig:"PRIVATE_KEY"`
	}
)

func AuthenticateMiddelware(h http.Handler) http.Handler {
	//TODO change the location and value of private key
	var conf JWTConf
	envconfig.Load(&conf)
	var ignorlist []string
	ignorlist = append(ignorlist, "/api/v1/login")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)
		if Contains(ignorlist, r.RequestURI) {
			h.ServeHTTP(w, r)
			return
		}
		fmt.Println(conf.PrivateKey)
		thebearertoken := r.Header.Get("Authorization")
		tokenarry := strings.Split(thebearertoken, " ")
		if len(tokenarry) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Can find your token pls adding bearetoken beore query"))
			return
		}
		jwttoken := strings.Split(thebearertoken, " ")[1]
		logrus.Info("the token : %v\n", jwttoken)
		notExpire := myjwt.CheckExp(jwttoken)
		if notExpire && myjwt.Checkvalib(jwttoken) {
			payload := myjwt.GetPayload(jwttoken)
			logrus.Info("the payload is %v\n", payload)
			w.Write([]byte(payload.Usefullname))
			h.ServeHTTP(w, r)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Your token not valib or your session is expire plse login agin"))
		return
	})
}
