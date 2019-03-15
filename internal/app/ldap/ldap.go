package ldap

import (
	"fmt"
	"gopkg.in/ldap.v3"
	"log"
	"strings"

	"github.com/TheFlies/ofriends/internal/app/service"
	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/pkg/config/env"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
	"github.com/TheFlies/ofriends/internal/pkg/jwt"
)

type (
	LDAPconf struct {
		LdapAddr              string `envconfig:"LDAP_ADDR"`
		LdapPort              int    `envconfig:"LDAP_PORT"`
		LdapFirstBindUsername string `envconfig:"LDAP_FIRST_BIND_USERNAME"`
		LdapFirstBindPassword string `envconfig:"LDAP_FIRST_BIND_PASSWORD"`
		BaseDN                string `envconfig:"BASEDN"`
	}
	LdapAuthentication struct {
		log     glog.Logger
		conf    LDAPconf
		usersvr service.Userservice
	}
)

func New(svr service.Userservice) *LdapAuthentication {
	var ldapconf LDAPconf
	l := glog.New()
	l.WithField("package", "ldap")
	envconfig.Load(&ldapconf)
	return &LdapAuthentication{
		log:     l,
		conf:    ldapconf,
		usersvr: svr,
	}
}
func (l *LdapAuthentication) Authenticate(username string, password string) (interface{}, error) {
	err := ldapauthentication(username, password, l.conf)
	responmap := make(map[string]interface{})
	if err != nil {
		l.log.Errorf("login fail %v", err)
		return "", err
	} else {
		l.log.Infof("login success, waiting query user information")
		isExists := l.usersvr.CheckExistence(username)
		if isExists {
			l.log.Infof("query user information form database")
			dbUser, err := l.usersvr.GetByName(username)
			if err != nil {
				l.log.Errorf("Get user fail")
				return "", err
			}
			jwttoken, err := jwt.CreateToken(dbUser.Username, dbUser.Fullname, dbUser.Userrole)
			if err != nil {
				l.log.Errorf("Get user fail")
				return "", err
			}
			responmap["userinformation"] = dbUser
			responmap["token"] = jwttoken
		} else {
			ldapuser, err := ldapquery(username, l.conf)
			if err != nil {
				l.log.Errorf("Get user fail")
				return "", err
			}
			jwttoken, err := jwt.CreateToken(ldapuser.Username, ldapuser.Fullname, ldapuser.Userrole)
			if err != nil {
				l.log.Errorf("Get user fail")
				return "", err
			}
			responmap["userinformation"] = ldapuser
			responmap["token"] = jwttoken
		}
		return responmap, nil
	}

}
func ldapauthentication(username string, password string, cf LDAPconf) error {
	ldapconn, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", cf.LdapAddr, cf.LdapPort))
	if err != nil {
		log.Fatal(err)
	}
	defer ldapconn.Close()
	var sb strings.Builder
	sb.WriteString("Cn=")
	sb.WriteString(username)
	sb.WriteString(",")
	sb.WriteString(cf.BaseDN)
	err = ldapconn.Bind(sb.String(), password)
	if err != nil {
		return err
	}
	return nil
}
func ldapquery(username string, cf LDAPconf) (types.User, error) {
	ldapconn, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", cf.LdapAddr, cf.LdapPort))
	if err != nil {
		return types.User{}, err
	}
	defer ldapconn.Close()
	err = ldapconn.Bind(cf.LdapFirstBindUsername, cf.LdapFirstBindPassword)
	if err != nil {
		return types.User{}, err
	}
	searchRequest := ldap.NewSearchRequest(
		cf.BaseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=*)(cn=%s))", username),
		[]string{"dn", "mail", "memberOf", "description", "displayName"},
		nil,
	)
	searchrequest, err := ldapconn.Search(searchRequest)
	if err != nil {
		fmt.Println("search fails")
		return types.User{}, err
	}
	emailaddr := searchrequest.Entries[0].GetAttributeValues("mail")
	role := searchrequest.Entries[0].GetAttributeValues("description")
	fulname := searchrequest.Entries[0].GetAttributeValues("displayName")
	return types.User{
		Username: username,
		Fullname: fulname[0],
		Userrole: role[0],
		Email:    emailaddr[0],
	}, nil
}
