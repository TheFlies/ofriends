package middleware

import (
	"github.com/TheFlies/ofriends/internal/pkg/myjwt"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

func Authent(h http.Handler) http.Handler {
	//TODO change the location and value of private key
	PRIVATE_KEY := "AllYourBase"

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		thebearertoken := r.Header.Get("Authorization")
		jwttoken := strings.Split(thebearertoken, " ")[1]
		logrus.Info("the token : %v\n", jwttoken)
		unixtimenow := time.Now().Unix()
		notExpire := myjwt.CheckExp(jwttoken, PRIVATE_KEY, unixtimenow)
		if notExpire {
			payload := myjwt.GetPayload(jwttoken, PRIVATE_KEY)
			logrus.Info("the payload is %v\n", payload)
			w.Write([]byte(payload.Usefullname))
			h.ServeHTTP(w, r)
		}
		//TODO changne the return information for client
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("you not have"))
	})
}
