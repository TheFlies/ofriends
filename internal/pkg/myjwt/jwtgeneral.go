package myjwt

import (
	"github.com/dgrijalva/jwt-go"
)

type (
	Customeclaims struct {
		jwt.StandardClaims
		Useid       string
		Usefullname string
		Userole     string
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

func CreateToken(issuedAt int64, audience string, timeout int64, serectkey string, useid string, usefullname string, userole string) (string, error) {
	payload := New(jwt.StandardClaims{
		IssuedAt:  issuedAt,
		Audience:  audience,
		ExpiresAt: timeout,
	},
		useid, usefullname, userole)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	ss, err := token.SignedString([]byte(serectkey))
	if err != nil {
		return "", err
	}
	return ss, nil
}
func GetPayload(token string, serectkey string) Customeclaims {
	reslt, _ := jwt.ParseWithClaims(token, &Customeclaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(serectkey), nil
	})
	payload, _ := reslt.Claims.(*Customeclaims)
	return *payload
}

// This method will check a token is expired with current time
// Return true if the token still no expired
func CheckExp(token string, serectkey string, timenow int64) bool {
	reslt, _ := jwt.ParseWithClaims(token, &Customeclaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(serectkey), nil
	})
	payload, ok := reslt.Claims.(*Customeclaims)
	if ok && payload.VerifyExpiresAt(timenow, false) {
		return true
	}
	return false
}
