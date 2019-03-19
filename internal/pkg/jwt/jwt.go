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
		UseID          string
		UserFullname   string
		DeliveryCenter []string
	}
	JWTConf struct {
		PrivateKey string `envconfig:"PRIVATE_KEY"`
		Duration   string `envconfig:"JWT_DURATION"`
	}
)

func New(standard jwt.StandardClaims, uid string, fullname string, deliverycenter []string) *Customeclaims {
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
		UseID:          uid,
		UserFullname:   fullname,
		DeliveryCenter: deliverycenter,
	}
}

func CreateToken(uid string, fullname string, deliverycenter []string) (string, error) {
	logger := glog.New().WithField("package", "jwt")
	var conf JWTConf
	envconfig.Load(&conf)
	logger.Infof("environment list [%v]", conf)
	confDuration, err := time.ParseDuration(conf.Duration)
	logger.Infof("duration time is : %v", confDuration)
	if err != nil {
		logger.Errorf("can't parser duration form server's environment variable: %v", err)
	}
	duration := time.Now().Add(confDuration).Unix()
	payload := New(jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: duration,
	},
		uid, fullname, deliverycenter)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	JWTToken, err := token.SignedString([]byte(conf.PrivateKey))
	if err != nil {
		logger.Errorf("can't create token with errors : %v", err)
		return "", err
	}
	return JWTToken, nil
}
func GetPayload(token string) Customeclaims {
	logger := glog.New().WithField("package", "jwt")
	var conf JWTConf
	envconfig.Load(&conf)
	logger.Infof("environment list [%v]", conf)
	result, err := jwt.ParseWithClaims(token, &Customeclaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(conf.PrivateKey), nil
	})
	if err != nil {
		logger.Errorf("can not parse claims form token errors: %v", err)
	}
	logger.Infof("get token success ")
	payload, _ := result.Claims.(*Customeclaims)
	return *payload
}
func CheckValib(token string) bool {
	var conf JWTConf
	envconfig.Load(&conf)
	result, err := jwt.ParseWithClaims(token, &Customeclaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(conf.PrivateKey), nil
	})
	if err != nil {
		return false
	}
	return result.Valid
}

// This method will check a token is expired with current time
// Return true if the token still no expired
func CheckExp(token string) bool {
	timeNow := time.Now()
	logger := glog.New().WithField("package", "jwt")
	logger.Infof("time now at :", timeNow.Format(time.ANSIC))
	var conf JWTConf
	envconfig.Load(&conf)
	result, _ := jwt.ParseWithClaims(token, &Customeclaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(conf.PrivateKey), nil
	})
	payload, ok := result.Claims.(*Customeclaims)
	if ok && payload.VerifyExpiresAt(timeNow.Unix(), false) {
		return true
	}
	return false
}
