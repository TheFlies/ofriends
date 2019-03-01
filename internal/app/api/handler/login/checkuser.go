package login

import (
	"github.com/TheFlies/ofriends/internal/pkg/myjwt"
	"log"
	"time"
)

const ()

func Authenticationuser(usename string, password string) (string, bool) {
	primaryley := "AllYourBase"
	// this funtion still not implemnt now.
	// TODO impletment with ldap authentication.
	// Thi funton will check usename and password if both is vlaib will return the token and true value
	// the token timeout is 60minus
	issusedtime := time.Now().Unix()
	exprietime := time.Now().Local().Add(time.Hour*time.Duration(0) +
		time.Minute*time.Duration(60) +
		time.Second*time.Duration(0)).Unix()
	token, err := myjwt.CreateToken(issusedtime, "Use", exprietime, primaryley, usename, "khagn", "eng")
	if err != nil {
		log.Fatal("Fail to create token ")
		return "", false
	}
	return token, true
}
