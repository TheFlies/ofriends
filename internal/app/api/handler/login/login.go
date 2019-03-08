package login

import (
	"fmt"
	"github.com/TheFlies/ofriends/internal/app/ldap"
	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
	"github.com/TheFlies/ofriends/internal/pkg/jwt"
	"github.com/TheFlies/ofriends/internal/pkg/respond"
	"github.com/globalsign/mgo/bson"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type (
	service interface {
		GetUserByeUserName(username string) (*types.Usercache, error)
		AddUser(user *types.Usercache) error
		CheckUseExist(username string) bool
	}
	LoginHandler struct {
		srv    service
		logger glog.Logger
	}
)

func NewLoginHandeler(s service, l glog.Logger) *LoginHandler {
	return &LoginHandler{
		srv:    s,
		logger: l,
	}
}
func (h *LoginHandler) Authentication(w http.ResponseWriter, r *http.Request) {
	usename := r.FormValue("username")
	password := r.FormValue("password")
	err := ldap.LDAPAuthentication(usename, password)
	resondamp := make(map[string]interface{})
	if err != nil {
		respond.JSON(w, http.StatusUnauthorized, "login fails please check uou username and passwrod")
		return
	}
	dbuser, err := h.srv.GetUserByeUserName(usename)
	if err != nil {
		// data not exist in database
		ldapuser, err := ldap.LDAPQuery(usename)
		if err != nil {
			respond.JSON(w, http.StatusInternalServerError, err)
			return
		}
		ldapuser.Id = bson.NewObjectId()
		err = h.srv.AddUser(&ldapuser)
		if err != nil {
			respond.JSON(w, http.StatusInternalServerError, err)
			return
		}
		token, err := jwt.CreateToken(ldapuser.Username, ldapuser.Fullname, ldapuser.Userrole, time.Now().Unix())
		if err != nil {
			respond.JSON(w, http.StatusInternalServerError, err)
			return
		}
		resondamp["user"] = ldapuser
		resondamp["token"] = token
		respond.JSON(w, http.StatusAccepted, resondamp)
		return
	}

	token, err := jwt.CreateToken(dbuser.Username, dbuser.Fullname, dbuser.Userrole, time.Now().Unix())
	if err != nil {
		respond.JSON(w, http.StatusInternalServerError, err)
		return
	}
	resondamp["user"] = dbuser
	fmt.Println("dbuser")
	fmt.Println(dbuser)
	resondamp["token"] = token
	respond.JSON(w, http.StatusAccepted, resondamp)
	return

}
func (h *LoginHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	usename := r.FormValue("username")
	u, err := h.srv.GetUserByeUserName(usename)
	logrus.Infof("username %v", u)
	if err != nil {
		respond.JSON(w, http.StatusInternalServerError, err)
		return
	}
	respond.JSON(w, http.StatusAccepted, u)
	return
}
