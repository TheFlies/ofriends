package ldap

import (
	"fmt"
	"regexp"

	"github.com/globalsign/mgo/bson"
	"github.com/nasa9084/go-strlib"
	"github.com/pkg/errors"
	"gopkg.in/ldap.v3"

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
		lconn   *ldap.Conn
	}
)

func New(svr service.Userservice) *LdapAuthentication {
	var ldapconf LDAPconf
	log := glog.New()
	l := log.WithField("package", "ldap")
	envconfig.Load(&ldapconf)
	return &LdapAuthentication{
		log:     l,
		conf:    ldapconf,
		usersvr: svr,
	}
}
func (l *LdapAuthentication) Conection() {
	conection, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", l.conf.LdapAddr, l.conf.LdapPort))
	if err != nil {
		l.log.Errorf("connection fails")
	}
	l.lconn = conection

}
func (l *LdapAuthentication) Authenticate(username string, password string) (interface{}, error) {
	l.Conection()
	defer l.lconn.Close()
	err := ldapauthentication(username, password, l.conf, l.log, l.lconn)
	responmap := make(map[string]interface{})
	if err != nil {
		l.log.Errorf("login fail %v", err)
		return "", err
	}
	l.log.Infof("login success, waiting query user information")
	isExists := l.usersvr.CheckExistence(username)
	if isExists {
		l.log.Infof("query user information form database")
		dbUser, err := l.usersvr.GetByName(username)
		if err != nil {
			l.log.Errorf("get user fail")
			return "", err
		}
		jwttoken, err := jwt.CreateToken(dbUser.Username, dbUser.Fullname, dbUser.DeliveryCenter)
		if err != nil {
			l.log.Errorf("get user fail")
			return "", err
		}
		responmap["userinformation"] = dbUser
		responmap["token"] = jwttoken
		return responmap, nil
	}
	ldapuser, err := ldapquery(username, l.conf, l.log, l.lconn)
	l.log.Infof("query user information in ldap ")
	if err != nil {
		l.log.Errorf("get user fail")
		return "", err
	}
	ldapuser.Id = bson.NewObjectId()
	l.usersvr.AddUser(&ldapuser)
	jwttoken, err := jwt.CreateToken(ldapuser.Username, ldapuser.Fullname, ldapuser.DeliveryCenter)
	if err != nil {
		l.log.Errorf("Get user fail")
		return "", err
	}
	responmap["userinformation"] = ldapuser
	responmap["token"] = jwttoken
	return responmap, nil

}
func ldapauthentication(username string, password string, cf LDAPconf, l glog.Logger, conn *ldap.Conn) error {
	userdn, err := finduserdn(username, cf, l, conn)
	l.Infof("userdn= %v", userdn)
	if err != nil {
		return err
	}
	ok := false
	for _, dn := range userdn {
		err = conn.Bind(dn, password)
		if err == nil {
			ok = true
			break
		}
	}
	if ok {
		return nil
	}
	return errors.New("login fails")

}
func ldapquery(username string, cf LDAPconf, l glog.Logger, conn *ldap.Conn) (types.User, error) {
	err := conn.Bind(cf.LdapFirstBindUsername, cf.LdapFirstBindPassword)
	if err != nil {
		return types.User{}, err
	}
	searchRequest := ldap.NewSearchRequest(
		cf.BaseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=*)(cn=%s))", username),
		[]string{"dn", "mail", "memberOf", "description", "displayName", "cn"},
		nil,
	)
	searchresult, err := conn.Search(searchRequest)
	if err != nil {
		l.Errorf("ldap query search fail")
		return types.User{}, err
	}
	if len(searchresult.Entries) == 0 {
		l.Errorf("no user query")
		err := errors.New("user does not exist ")
		return types.User{}, err
	}
	for _, entries := range searchresult.Entries {
		if strlib.IsLower(entries.GetAttributeValue("cn")) {
			emailaddr := entries.GetAttributeValue("mail")
			ds := entries.GetAttributeValue("description")
			// try to get DC name form description value
			dc := getdcformdescripton(ds, "DC[0-9]+")
			fulname := entries.GetAttributeValues("displayName")
			var name string
			if len(fulname) == 0 {
				name = ""
			}
			name = fulname[0]
			return types.User{
				Username:       username,
				Fullname:       name,
				DeliveryCenter: dc,
				Email:          emailaddr,
			}, nil
		}
	}
	return types.User{}, errors.New("can't get user information form ldap")
}
func finduserdn(username string, cf LDAPconf, l glog.Logger, conn *ldap.Conn) ([]string, error) {
	err := conn.Bind(cf.LdapFirstBindUsername, cf.LdapFirstBindPassword)
	if err != nil {
		l.Errorf("can't first bind to ldap database to query user dn ")
		return nil, err
	}
	l.Infof("query user by basedn : %s", cf.BaseDN)
	searchRequest := ldap.NewSearchRequest(
		cf.BaseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=*)(cn=%s))", username),
		[]string{"dn"},
		nil,
	)
	searresult, err := conn.Search(searchRequest)
	if err != nil {
		l.Errorf("search dn for login fail")
		return nil, err
	}
	if len(searresult.Entries) == 0 {
		l.Errorf("query fail")
		err := errors.New("user does not exist")
		return nil, err
	}
	dn := make([]string, len(searresult.Entries))
	for _, stringdn := range searresult.Entries {
		dn = append(dn, stringdn.DN)
	}
	return dn, nil

}
func getdcformdescripton(s string, r string) string {
	re := regexp.MustCompile(r)
	result := re.FindAllString(s, 1)
	if len(result) != 1 {
		return "wanderer"
	}
	return result[0]
}
