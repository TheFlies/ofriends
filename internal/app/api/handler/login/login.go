package login

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
	"github.com/TheFlies/ofriends/internal/pkg/respond"
)

type (
	service interface {
		GetByName(username string) (*types.Usercache, error)
		AddUser(user *types.Usercache) error
		CheckExistence(username string) bool
		LDAPAuthenticate(username string, password string) (interface{}, error)
		//DatabaseAuthenticate (usename string,password string) bool

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
func (h *LoginHandler) Authenticate(w http.ResponseWriter, r *http.Request) {
	usename := r.FormValue("username")
	password := r.FormValue("password")
	h.logger.Infof("login with username : %v,password:****** ", usename)
	respondmap, err := h.srv.LDAPAuthenticate(usename, password)
	if err != nil {
		respond.JSON(w, http.StatusInternalServerError, err)
		return
	}
	respond.JSON(w, http.StatusAccepted, respondmap)

}
func (h *LoginHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	usename := r.FormValue("username")
	u, err := h.srv.GetByName(usename)
	logrus.Infof("username %v", u)
	if err != nil {
		respond.JSON(w, http.StatusInternalServerError, err)
		return
	}
	respond.JSON(w, http.StatusAccepted, u)
	return
}
