package login

import (
	"net/http"

	"github.com/TheFlies/ofriends/internal/pkg/glog"
	"github.com/TheFlies/ofriends/internal/pkg/respond"
)

type (
	LoginHandler struct {
		logger        glog.Logger
		localLoginSrv LoginService
		ldapLoginSrv  LoginService
	}
	LoginService interface {
		Authenticate(username string, password string) (interface{}, error)
	}
)

func NewLoginHandler(localLogin LoginService, ldapLogin LoginService, l glog.Logger) *LoginHandler {
	return &LoginHandler{
		logger:        l,
		localLoginSrv: localLogin,
		ldapLoginSrv:  ldapLogin,
	}
}
func (h *LoginHandler) Authenticate(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	h.logger.Infof("login with Username : %v,Password:****** ", username)
	respondValue, err := h.localLoginSrv.Authenticate(username, password)
	if err != nil {
		respondValue, err = h.ldapLoginSrv.Authenticate(username, password)
		if err != nil {
			respond.JSON(w, http.StatusForbidden, map[string]string{"status": "403", "message": "your username and password incorrect"})
			return
		}
	}
	respond.JSON(w, http.StatusAccepted, respondValue)
	return
}
