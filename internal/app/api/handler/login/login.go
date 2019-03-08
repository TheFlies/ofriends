package login

import (
	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
	"github.com/TheFlies/ofriends/internal/pkg/myjwt"
	"github.com/TheFlies/ofriends/internal/pkg/respond"
	"github.com/globalsign/mgo/bson"
	"net/http"
	"time"
)

type (
	service interface {
		GetUserByeUserName(username string) (*types.Usercache, error)
		AddUser(user *types.Usercache) error
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
	err := LDAPAuthentication(usename, password)
	if err != nil {
		respond.JSON(w, http.StatusUnauthorized, "login fails please check uou username and passwrod")
		return
	}
	user, err := h.srv.GetUserByeUserName(usename)
	if err != nil {
		user.Id = bson.NewObjectId()
		err := h.srv.AddUser(user)
		if err != nil {
			respond.JSON(w, http.StatusInternalServerError, err)
			return
		}
	}

	token, err := myjwt.CreateToken(user.Username, user.Fullname, user.Userrole, time.Now().Unix())
	if err != nil {
		respond.JSON(w, http.StatusInternalServerError, err)
		return
	}
	responmap := make(map[string]interface{})
	responmap["userinformation"] = user
	responmap["token"] = token
	respond.JSON(w, http.StatusAccepted, responmap)

}
func (h *LoginHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	usename := r.FormValue("username")
	u, err := h.srv.GetUserByeUserName(usename)
	if err != nil {
		respond.JSON(w, http.StatusAccepted, err)
		return
	}
	respond.JSON(w, http.StatusAccepted, u)
	return
}
