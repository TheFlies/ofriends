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
		UseId          string
		UserFullName   string
		Deliverycenter []string
	}
	JWTConf struct {
		PrivateKey string `envconfig:"PRIVATE_KEY"`
		Duration   string `envconfig:"JWT_DURATION"`
	}
)

func New(standard jwt.StandardClaims, uid string, ufullname string, dc []string) *Customeclaims {
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
		UseId:          uid,
		UserFullName:   ufullname,
		Deliverycenter: dc,
	}
}

func CreateToken(uid string, ufullname string, dc []string) (string, error) {
	logger := glog.New().WithField("package", "jwt")
	var conf JWTConf
	envconfig.Load(&conf)
	logger.Infof("environment list [%v]", conf)
	confduration, err := time.ParseDuration(conf.Duration)
	logger.Infof("duration time is : %v", confduration)
	if err != nil {
		logger.Errorf("can't parser duration form server's environment variable: %v", err)
	}
	duration := time.Now().Add(confduration).Unix()
	payload := New(jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: duration,
	},
		uid, ufullname, dc)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	ss, err := token.SignedString([]byte(conf.PrivateKey))
	if err != nil {
		logger.Errorf("can't create token with errors : %v", err)
		return "", err
	}
	return ss, nil
}
func GetPayload(token string) Customeclaims {
	logger := glog.New().WithField("package", "jwt")
	var conf JWTConf
	envconfig.Load(&conf)
	logger.Infof("environment list [%v]", conf)
	reslt, err := jwt.ParseWithClaims(token, &Customeclaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(conf.PrivateKey), nil
	})
	if err != nil {
		logger.Errorf("can not parse claims form token errors: %v", err)
	}
	logger.Infof("get token success ")
	payload, _ := reslt.Claims.(*Customeclaims)
	return *payload
}
func CheckValib(token string) bool {
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
	timenow := time.Now()
	logger := glog.New().WithField("package", "jwt")
	logger.Infof("time now at :", timenow.Format(time.ANSIC))
	var conf JWTConf
	envconfig.Load(&conf)
	reslt, _ := jwt.ParseWithClaims(token, &Customeclaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(conf.PrivateKey), nil
	})
	payload, ok := reslt.Claims.(*Customeclaims)
	if ok && payload.VerifyExpiresAt(timenow.Unix(), false) {
		return true
	}
	return false
}
