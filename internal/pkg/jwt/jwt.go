package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/TheFlies/ofriends/internal/pkg/config/env"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
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
		Duration   string `envconfig:"JWT_DURATION"`
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

func CreateToken(userId string, userFullName string, userRole string) (string, error) {
	logger := glog.New()
	logger.WithField("package", "jwt")
	var conf JWTConf
	envconfig.Load(&conf)
	confduration, err := time.ParseDuration(conf.Duration)
	logger.Infof("duratione time is : %v", confduration)
	if err != nil {
		logger.Errorf("can't parser duraration form server envairoment: &v", err)
	}
	duration := time.Now().Add(confduration).Unix()
	payload := New(jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: duration,
	},
		userId, userFullName, userRole)
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
