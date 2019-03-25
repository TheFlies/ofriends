package login

import (
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"

	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
	"github.com/TheFlies/ofriends/internal/pkg/respond"
)

type (
	LoginHandler struct {
		srv      UserService
		logger   glog.Logger
		loginSvr []LoginService
	}
	UserService interface {
		GetByName(username string) (*types.User, error)
		AddUser(user *types.User) error
		CheckExistence(username string) bool
		ChangeUserPassword(user *types.User) error
	}
	LoginService interface {
		Authenticate(username string, password string) (interface{}, error)
	}
)

func NewLoginHandeler(s UserService, loginsvr []LoginService, l glog.Logger) *LoginHandler {
	return &LoginHandler{
		srv:      s,
		logger:   l,
		loginSvr: loginsvr,
	}
}
func (h *LoginHandler) Authenticate(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("Login with %d method ", len(h.loginSvr))
	Username := r.FormValue("username")
	Password := r.FormValue("password")
	h.logger.Infof("login with Username : %v,Password:****** ", Username)
	for _, loginsvr := range h.loginSvr {
		respondMap, err := loginsvr.Authenticate(Username, Password)
		if err == nil {
			respond.JSON(w, http.StatusAccepted, respondMap)
			return
		}
	}
	respond.JSON(w, http.StatusUnauthorized, "login fail")
	return

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
func (h *LoginHandler) Register(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	rawPassword := r.FormValue("password")
	fullName := r.FormValue("fullname")
	email := r.FormValue("email")
	deliveryCenter := r.FormValue("deliverycenter")
	dcs := []string{deliveryCenter}
	bytes, err := bcrypt.GenerateFromPassword([]byte(rawPassword), 14)
	if err != nil {
		h.logger.Errorf("password encryption is fail err: %v", err)
		respond.JSON(w, http.StatusUnauthorized, "register fail")
		return
	}
	password := string(bytes)
	user := types.User{
		ID:             bson.NewObjectId(),
		Username:       username,
		FullName:       fullName,
		Email:          email,
		DeliveryCenter: dcs,
		Password:       password,
	}
	err = h.srv.AddUser(&user)
	if err != nil {
		h.logger.Infof("can't save new user into database err: %v ", err)
		respond.JSON(w, http.StatusInternalServerError, "add user fail")
		return
	}
	respond.JSON(w, http.StatusCreated, "save user successfully")
}
func (h *LoginHandler) ChangeLocalPassword(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	newPassword := r.FormValue("newpassword")
	_, err := h.loginSvr[1].Authenticate(username, password)
	if err != nil {
		h.logger.Errorf("user authentication is fail err:%v ", err)
		respond.JSON(w, http.StatusUnauthorized, "you must type correctly your user name and password before you changes password")
		return
	}
	user, err := h.srv.GetByName(username)
	if err != nil {
		h.logger.Errorf("can't get user form database  err %v", err)
		respond.JSON(w, http.StatusInternalServerError, "can't get user form database")
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), 14)
	if err != nil {
		h.logger.Errorf("encryption your password is fail err:%v", err)
		respond.JSON(w, http.StatusInternalServerError, "have a error when update password")
		return
	}
	user.Password = string(hashPassword)
	h.logger.Infof("newPassword  : %v encrypt pass : %v", newPassword, string(hashPassword))
	err = h.srv.ChangeUserPassword(user)
	if err != nil {
		h.logger.Errorf("can't update password err: %v,", err)
		respond.JSON(w, http.StatusInternalServerError, "the password isn't change ")
		return
	}
	respond.JSON(w, http.StatusOK, "your password been changed")
	return
}
func (h *LoginHandler) SetLocalPassword(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	localPassword := r.FormValue("localpassword")
	_, err := h.loginSvr[0].Authenticate(username, password)
	if err != nil {
		h.logger.Errorf("user ldap authentication is fail err:%v ", err)
		respond.JSON(w, http.StatusUnauthorized, "you must type correctly your user name and password before change password")
		return
	}
	user, err := h.srv.GetByName(username)
	if err != nil {
		h.logger.Infof("can't get user form database err:%v", err)
		respond.JSON(w, http.StatusInternalServerError, "can't find user form database")
		return
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(localPassword), 14)
	h.logger.Infof("hasspass :%s", string(hashPassword))
	h.logger.Infof("local pass : %v", localPassword)
	if err != nil {
		h.logger.Errorf("encryption your password is fail err:%v", err)
		respond.JSON(w, http.StatusInternalServerError, "have a error when set local password")
		return
	}
	user.Password = string(hashPassword)
	err = h.srv.ChangeUserPassword(user)
	if err != nil {
		h.logger.Errorf("setting local password is fail")
		respond.JSON(w, http.StatusInternalServerError, "have a error when set local password")
		return
	}
	h.logger.Infof("setting local password is fail")
	respond.JSON(w, http.StatusOK, "setting local password is successful")
	return

}
