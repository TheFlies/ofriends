package login

import (
	"github.com/TheFlies/ofriends/internal/app/friend"
	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
	"github.com/TheFlies/ofriends/internal/pkg/myjwt"
	"github.com/TheFlies/ofriends/internal/pkg/respond"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type (

	// Handler is index web handler
	LoginHandeler struct {
		usercacherepo friend.UseCacheRepository
		logger        glog.Logger
	}
)

func New(repo friend.UseCacheRepository, l glog.Logger) *LoginHandeler {
	return &LoginHandeler{
		usercacherepo: repo,
		logger:        l,
	}
}
func (web *LoginHandeler) Login(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("web/login.html")
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	if _, err := io.Copy(w, f); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
}
func (web *LoginHandeler) Authentication(w http.ResponseWriter, r *http.Request) {
	Primarykey := "AllYorbase"
	//TODO ldap will query username role and fullname
	usename := r.FormValue("usename")
	password := r.FormValue("password")
	user := types.Usercache{Username: usename, Userrole: "eng", Fullname: "full name"}
	if ok := LDAPAuthentication(usename, password); ok {
		//Check user in database, if the user did't exist, get user information in ldap database and save them into mongodb
		_, err := web.usercacherepo.FindUserByUserName(usename)
		if err != nil {
			if err := web.usercacherepo.InserUser(&user); err != nil {
				log.Fatal("insert a user to cache database is fails")
				log.Fatal(err)
				respond.Error(w, err, http.StatusInternalServerError)
			}
		}
		timeout := time.Now().Local().Add(time.Hour*time.Duration(0) +
			time.Minute*time.Duration(60) +
			time.Second*time.Duration(0)).Unix()
		token, err := myjwt.CreateToken(time.Now().Unix(), "", timeout, Primarykey, usename, "fullname", "eng")
		if err != nil {
			log.Fatal("creating token is fails")
			respond.Error(w, err, http.StatusInternalServerError)
		}
		respondmap := make(map[string]string)
		respondmap["token"] = token
		respondmap["usename"] = usename
		respond.JSON(w, http.StatusAccepted, respondmap)
	}
}
