package login

import (
	"github.com/TheFlies/ofriends/internal/app/email"
	"net/http"

	"github.com/TheFlies/ofriends/internal/pkg/glog"
	"github.com/TheFlies/ofriends/internal/pkg/respond"
)

type (
	LoginHandler struct {
		logger        glog.Logger
		localLoginSrv LoginService
		ldapLoginSrv  LoginService
		sendMailSrv   email.MailService
	}
	LoginService interface {
		Authenticate(username string, password string) (interface{}, error)
	}
)

func NewLoginHandler(localLogin LoginService, ldapLogin LoginService, emailSrv email.MailService, l glog.Logger) *LoginHandler {
	return &LoginHandler{
		logger:        l,
		localLoginSrv: localLogin,
		ldapLoginSrv:  ldapLogin,
		sendMailSrv:   emailSrv,
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
func (h *LoginHandler) SendMail(w http.ResponseWriter, r *http.Request) {
	if err := h.sendMailSrv.Send(r.Context()); err != nil {
		respond.JSON(w, 500, err)
		return
	}
	respond.JSON(w, 200, "seen")
	return
}
