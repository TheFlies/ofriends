package middleware

import (
	"net/http"
	"sort"
	"strings"

	"github.com/TheFlies/ofriends/internal/pkg/glog"
	"github.com/TheFlies/ofriends/internal/pkg/jwt"
	"github.com/TheFlies/ofriends/internal/pkg/respond"
)

const (
	none      = 1
	roleUser  = 2
	roleAdmin = 3
)

func Authorization(roles ...int) func(http.HandlerFunc) http.HandlerFunc {
	l := glog.New().WithField("package", "middleware")
	if len(roles) == 0 {
		l.Errorf("role list must not blank")
		return nil
	}
	sort.SliceIsSorted(roles, func(i, j int) bool {
		return i > j
	})
	l.Infof("list role %v", roles)
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if ok := RoleCheck(r, roles...); !ok {
				l.Infof("authorization fail")
				respond.JSON(w, http.StatusUnauthorized, map[string]string{"status": "401", "message": "authorization"})
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
func RoleCheck(r *http.Request, roles ...int) bool {
	theBearerToken := r.Header.Get("Authorization")
	tokenArray := strings.Split(theBearerToken, " ")
	if len(tokenArray) != 2 {
		return false
	}
	jwtToken := strings.Split(theBearerToken, " ")[1]
	JWTGeneration := jwt.NewTokenGeneration()
	notExpire := JWTGeneration.CheckExp(jwtToken)
	if !notExpire || !JWTGeneration.CheckValib(jwtToken) {
		return false
	}
	payload := JWTGeneration.GetPayload(jwtToken)
	jwtRolePriority := payload.Priority
	l := glog.New().WithField("package", "middleware")
	l.Infof("jwt priority geted : %v", jwtRolePriority)
	for _, lv := range roles {
		switch lv {
		case roleAdmin:
			l.Infof("checking in role admin role %v", roleAdmin)
			if jwtRolePriority < roleAdmin {
				return false
			}
		case roleUser:
			l.Infof("checking in role user")
			if jwtRolePriority < roleUser {
				return false
			}
		default:
			l.Infof("check default")
			return false
		}
	}
	return true
}
