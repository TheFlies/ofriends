package login

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/TheFlies/ofriends/internal/app/service"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
	"github.com/TheFlies/ofriends/internal/pkg/respond"
)

type (
	LoginHandler struct {
		srv      service.UserService
		logger   glog.Logger
		loginSvr service.Loginservice
	}
)

func NewLoginHandeler(s service.UserService, loginsvr service.Loginservice, l glog.Logger) *LoginHandler {
	return &LoginHandler{
		srv:      s,
		logger:   l,
		loginSvr: loginsvr,
	}
}
func (h *LoginHandler) Authenticate(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	h.logger.Infof("login with username : %v,password:****** ", username)
	respondMap, err := h.loginSvr.Authenticate(username, password)
	if err != nil {
		respond.JSON(w, http.StatusInternalServerError, "login fail")
		return
	}
	respond.JSON(w, http.StatusAccepted, respondMap)

}
func (h *LoginHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	u, err := h.srv.GetByName(username)
	logrus.Infof("username %v", u)
	if err != nil {
		respond.JSON(w, http.StatusInternalServerError, err)
		return
	}
	respond.JSON(w, http.StatusAccepted, u)
	return
}
