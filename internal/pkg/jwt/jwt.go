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
		UserName string
	}
	JWTConf struct {
		PrivateKey string `envconfig:"PRIVATE_KEY"`
		Duration   string `envconfig:"JWT_DURATION"`
	}
	TokenGeneration struct {
		logger glog.Logger
		conf   JWTConf
	}
)

func NewTokenGeneration() *TokenGeneration {
	logger := glog.New().WithField("package", "jwt")
	var conf JWTConf
	envconfig.Load(&conf)
	return &TokenGeneration{
		logger: logger,
		conf:   conf,
	}
}
func New(standard jwt.StandardClaims, uid string) *Customeclaims {
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
		UserName: uid,
	}
}
func (tkg *TokenGeneration) CreateToken(uid string) (string, error) {

	tkg.logger.Infof("environment list [%v]", tkg.conf)
	confDuration, err := time.ParseDuration(tkg.conf.Duration)
	tkg.logger.Infof("duration time is : %v", confDuration)
	if err != nil {
		tkg.logger.Errorf("can't parser duration form server's environment variable: %v", err)
	}
	duration := time.Now().Add(confDuration).Unix()
	payload := New(jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: duration,
	}, uid)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	JWTToken, err := token.SignedString([]byte(tkg.conf.PrivateKey))
	if err != nil {
		tkg.logger.Errorf("can't create token with errors : %v", err)
		return "", err
	}
	return JWTToken, nil
}
func (tkg *TokenGeneration) GetPayload(token string) Customeclaims {

	tkg.logger.Infof("environment list [%v]", tkg.conf)
	result, err := jwt.ParseWithClaims(token, &Customeclaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(tkg.conf.PrivateKey), nil
	})
	if err != nil {
		tkg.logger.Errorf("can not parse claims form token errors: %v", err)
	}
	tkg.logger.Infof("get token success ")
	payload, _ := result.Claims.(*Customeclaims)
	return *payload
}
func (tkg TokenGeneration) CheckValib(token string) bool {
	result, err := jwt.ParseWithClaims(token, &Customeclaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(tkg.conf.PrivateKey), nil
	})
	if err != nil {
		return false
	}
	return result.Valid
}

// This method will check a token is expired with current time
// Return true if the token still no expired
func (tkg *TokenGeneration) CheckExp(token string) bool {
	timeNow := time.Now()
	tkg.logger.Infof("time now at :", timeNow.Format(time.ANSIC))
	var conf JWTConf
	envconfig.Load(&conf)
	result, err := jwt.ParseWithClaims(token, &Customeclaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(conf.PrivateKey), nil
	})
	if err != nil {
		tkg.logger.Errorf("%v", err)
		return false
	}
	payload, ok := result.Claims.(*Customeclaims)
	if ok && payload.VerifyExpiresAt(timeNow.Unix(), false) {
		tkg.logger.Infof("Token not expire")
		return true
	}
	tkg.logger.Infof("Token is expire")
	return false
}
