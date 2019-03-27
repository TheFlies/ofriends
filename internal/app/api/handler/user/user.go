package user

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/TheFlies/ofriends/internal/app/api/handler/login"
	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
	"github.com/TheFlies/ofriends/internal/pkg/marshal"
	"github.com/TheFlies/ofriends/internal/pkg/respond"
)

type (
	UserHandler struct {
		srv           UserService
		logger        glog.Logger
		localLoginSrv login.LoginService
		ldapLoginSrv  login.LoginService
	}
	UserService interface {
		GetByName(username string) (*types.User, error)
		AddUser(u *types.User) (string, error)
		CheckExistence(username string) bool
		ChangePassword(username string, oldPassword string, newPassword string) error
		SetNewPassword(username string, password string) error
		UpdateUser(u *types.User) error
	}
)

func NewUserHandler(s UserService, localLogin login.LoginService, ldapLogin login.LoginService) *UserHandler {
	return &UserHandler{
		srv:           s,
		logger:        glog.New().WithField("package", "user"),
		localLoginSrv: localLogin,
		ldapLoginSrv:  ldapLogin,
	}
}
func (u *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	user, err := u.srv.GetByName(username)
	user.Password = " "
	if err != nil {
		respond.JSON(w, http.StatusNotFound, map[string]string{"status": "404", "message": "user is not found"})
		return
	}
	respond.JSON(w, http.StatusOK, user)
	return
}
func (u *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	u.logger.Infof("Registering user")
	requestData := types.User{}
	err := marshal.ParseRequest(r, &requestData)
	if err != nil {
		u.logger.Errorf("can't parse data form http request err: %v", err)
		respond.JSON(w, http.StatusBadRequest, map[string]string{"status": "400", "message": "can't get content form your request"})
	}
	err = requestData.Validate()
	if err != nil {
		u.logger.Errorf("some field of new user is malformed, &v", err)
		respond.JSON(w, http.StatusBadRequest, map[string]string{"status": "400", "message": fmt.Sprintf("input must type correctly input %v", err)})
	}
	userID, err := u.srv.AddUser(&requestData)
	if err != nil {
		respond.JSON(w, http.StatusInternalServerError, map[string]string{"status": "500", "message": "have an error when register for you"})
		return
	}
	respond.JSON(w, http.StatusOK, map[string]string{"status": "200", "userID": userID})
	return
}
func (u *UserHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	var varList = mux.Vars(r)
	user := varList["username"]
	requestData := types.User{}
	err := marshal.ParseRequest(r, &requestData)
	if err != nil {
		u.logger.Errorf("can't parse data form http request err: %v", err)
		respond.JSON(w, http.StatusBadRequest, map[string]string{"status": "400", "message": "can't get content form your request"})
		return
	}
	err = requestData.Validate()
	if err != nil {
		u.logger.Errorf("some field of new user is malformed, &v", err)
		respond.JSON(w, http.StatusBadRequest, map[string]string{"status": "400", "message": fmt.Sprintf("input must type correctly input %v", err)})
	}
	u.logger.Infof("old pass %v ", requestData.OldPassword)
	u.logger.Infof("new pass %v ", requestData.Password)
	err = u.srv.ChangePassword(user, requestData.OldPassword, requestData.Password)
	if err != nil {
		respond.JSON(w, http.StatusInternalServerError, map[string]string{"status": "500", "message": "change password fail,your username or password incorrect"})
		return
	}
	respond.JSON(w, http.StatusOK, map[string]string{"status": "200", "message": "changing your password is successful"})
	return
}
func (u *UserHandler) SetPassword(w http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)
	userName := vars["username"]
	u.logger.Infof("userName := %v", userName)
	requestData := types.User{}
	err := marshal.ParseRequest(r, &requestData)
	if err != nil {
		u.logger.Errorf("can't get json form http request err:&v", err)
		respond.JSON(w, http.StatusBadRequest, map[string]string{"status": "400", "message": "can't get content form your request"})
		return
	}
	err = requestData.Validate()
	if err != nil {
		u.logger.Errorf("some field of new user is malformed, &v", err)
		respond.JSON(w, http.StatusBadRequest, map[string]string{"status": "400", "message": fmt.Sprintf("input must type correctly input %v", err)})
	}
	err = u.srv.SetNewPassword(userName, requestData.Password)
	if err != nil {
		respond.JSON(w, http.StatusInternalServerError, map[string]string{"status": "500", "message": "setting your password fail,your username or password incorrect"})
		return
	}
	respond.JSON(w, http.StatusOK, map[string]string{"status": "200", "message": "setting your password is successful"})
	return

}
func (u *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	requestData := types.User{}
	err := marshal.ParseRequest(r, &requestData)
	if err != nil {
		u.logger.Errorf("can't get json form http request err:&v", err)
		respond.JSON(w, http.StatusBadRequest, map[string]string{"status": "400", "message": "can't get content form your request"})
		return
	}
	err = requestData.Validate()
	if err != nil {
		u.logger.Errorf("some field of new user is malformed, &v", err)
		respond.JSON(w, http.StatusBadRequest, map[string]string{"status": "400", "message": fmt.Sprintf("input must type correctly input %v", err)})
	}
	u.logger.Infof("update for username : %v", requestData.Username)
	err = u.srv.UpdateUser(&requestData)
	if err != nil {
		u.logger.Errorf("can't update err: %v", err)
		respond.JSON(w, http.StatusInternalServerError, map[string]string{"status": "500", "message": "updating user information is failure"})

		return
	}
	respond.JSON(w, http.StatusOK, map[string]string{"status": "200", "message": "updating user information is successful"})
	return

}
