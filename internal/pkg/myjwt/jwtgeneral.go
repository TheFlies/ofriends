package myjwt

import (
	"fmt"
	"github.com/TheFlies/ofriends/internal/pkg/config/env"
	"github.com/dgrijalva/jwt-go"
	"log"
	"strconv"
	"strings"
	"time"
)

type (
	Customeclaims struct {
		jwt.StandardClaims
		Useid       string
		Usefullname string
		Userole     string
	}
	JWTConf struct {
		PrivateKey string `envconfig:"PRIVATE_KEY"`
		TimeOut    string `envconfig:"TIME_OUT"`
	}
)

func New(standard jwt.StandardClaims, uuseid string, usefullname string, useroole string) *Customeclaims {
	return &Customeclaims{
		StandardClaims: jwt.StandardClaims{
			Audience:  standard.Audience,
			ExpiresAt: standard.ExpiresAt,
			Id:        standard.Id,
			IssuedAt:  standard.IssuedAt,
			Issuer:    standard.Issuer,
			NotBefore: standard.NotBefore,
			Subject:   standard.Subject,
		},
		Useid:       uuseid,
		Usefullname: usefullname,
		Userole:     useroole,
	}
}

func CreateToken(userid string, userfullnme string, userrole string, issueat int64) (string, error) {
	var conf JWTConf
	envconfig.Load(&conf)
	fmt.Println(conf.PrivateKey)
	rawtimeout := strings.Split(conf.TimeOut, ":")
	if len(rawtimeout) != 3 {
		log.Fatalf("too many agrumant fails to get time out form env %v", rawtimeout)
	}

	hour, err := strconv.Atoi(rawtimeout[0])
	if err != nil {
		log.Fatalf("can't cover value %v to int ", rawtimeout[0])
	}
	minutes, err := strconv.Atoi(rawtimeout[1])
	if err != nil {
		log.Fatalf("can't cover value %v to int ", rawtimeout[1])
	}
	second, err := strconv.Atoi(rawtimeout[2])
	if err != nil {
		log.Fatalf("can't cover value %v to int ", rawtimeout[2])
	}
	timeout := time.Now().Local().Add(time.Hour*time.Duration(hour) +
		time.Minute*time.Duration(minutes) +
		time.Second*time.Duration(second)).Unix()
	payload := New(jwt.StandardClaims{
		IssuedAt:  issueat,
		ExpiresAt: timeout,
	},
		userid, userfullnme, userrole)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	ss, err := token.SignedString([]byte(conf.PrivateKey))
	if err != nil {
		return "", err
	}
	return ss, nil
}
func GetPayload(token string) Customeclaims {
	var conf JWTConf
	envconfig.Load(&conf)
	reslt, _ := jwt.ParseWithClaims(token, &Customeclaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(conf.PrivateKey), nil
	})
	payload, _ := reslt.Claims.(*Customeclaims)
	return *payload
}
func Checkvalib(token string) bool {
	var conf JWTConf
	envconfig.Load(&conf)
	reslt, err := jwt.ParseWithClaims(token, &Customeclaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(conf.PrivateKey), nil
	})
	if err != nil {
		return false
	}
	return reslt.Valid
}

// This method will check a token is expired with current time
// Return true if the token still no expired
func CheckExp(token string) bool {
	timenow := time.Now().Unix()
	var conf JWTConf
	envconfig.Load(&conf)
	reslt, _ := jwt.ParseWithClaims(token, &Customeclaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(conf.PrivateKey), nil
	})
	payload, ok := reslt.Claims.(*Customeclaims)
	if ok && payload.VerifyExpiresAt(timenow, false) {
		return true
	}
	return false
}
