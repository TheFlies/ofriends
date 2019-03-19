package ldap

import (
	"fmt"
	"regexp"
	"unicode"

	"github.com/globalsign/mgo/bson"
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
		usersvr service.UserService
		lconn   *ldap.Conn
	}
)

func New(svr service.UserService) *LdapAuthentication {
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
func (l *LdapAuthentication) ConnectToServer() {
	l.log.Infof("connecting to ldap server with addr: %v , port : %v", l.conf.LdapAddr, l.conf.LdapPort)
	connection, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", l.conf.LdapAddr, l.conf.LdapPort))
	if err != nil {
		l.log.Errorf("connect to ldap server fails %v", err)
	}
	l.lconn = connection
}
func (l *LdapAuthentication) Close() {
	err := l.lconn.UnauthenticatedBind(l.conf.LdapFirstBindUsername)
	if err != nil {
		l.log.Errorf("can not unbind user : %s ", l.conf.LdapFirstBindUsername)
	}
	l.log.Infof("unbind user")
	l.lconn.Close()
	l.log.Infof("close a connection to : %v", l.conf.LdapAddr)
}
func (l *LdapAuthentication) Authenticate(username string, password string) (interface{}, error) {
	l.ConnectToServer()
	defer l.Close()
	err := ldapAuthentication(username, password, l.conf, l.log, l.lconn)
	responmap := make(map[string]interface{})
	if err != nil {
		l.log.Errorf("login fail %v ", err)
		return "", err
	}
	l.log.Infof("login success, waiting query user information")
	isExists := l.usersvr.CheckExistence(username)
	if isExists {
		l.log.Infof("query user information form database")
		dbUser, err := l.usersvr.GetByName(username)
		if err != nil {
			l.log.Errorf("get user fail: %v", err)
			return "", err
		}
		jwttoken, err := jwt.CreateToken(dbUser.UserName, dbUser.FullName, dbUser.DeliveryCenter)
		if err != nil {
			l.log.Errorf("get user fail : %v", err)
			return "", err
		}
		l.log.Infof("returning token")
		responmap["userinformation"] = dbUser
		responmap["token"] = jwttoken
		return responmap, nil
	}
	l.log.Infof("query user information in ldap server ")
	ldapuser, err := ldapQuery(username, l.conf, l.log, l.lconn)
	if err != nil {
		l.log.Errorf("get user fail")
		return "", err
	}
	ldapuser.Id = bson.NewObjectId()
	err = l.usersvr.AddUser(&ldapuser)
	if err != nil {
		l.log.Errorf("can not add a user to mongodb %v", err)
	}
	jwttoken, err := jwt.CreateToken(ldapuser.UserName, ldapuser.FullName, ldapuser.DeliveryCenter)
	if err != nil {
		l.log.Errorf("Get token fail : %v", err)
		return "", err
	}
	l.log.Infof("returning token")
	responmap["userinformation"] = ldapuser
	responmap["token"] = jwttoken
	return responmap, nil

}
func ldapAuthentication(username string, password string, cf LDAPconf, l glog.Logger, conn *ldap.Conn) error {
	userdn, err := findUserDN(username, cf, l, conn)
	l.Infof("login with user dn : %v", userdn)
	if err != nil {
		l.Errorf("find user dn fail %v", err)
		return err
	}
	// try login with list dn is queried
	// if all userdn fails, login will fail
	ok := false
	for _, dn := range userdn {
		err = conn.Bind(dn, password)
		if err == nil {
			l.Errorf("login fail with dn: %v", dn)
			ok = true
			break
		}
	}
	if ok {
		return nil
	}
	l.Errorf("all user dn is fail")
	return errors.New("login fails")

}
func ldapQuery(username string, cf LDAPconf, l glog.Logger, conn *ldap.Conn) (types.User, error) {
	l.Infof("query user information form ldap server by user name : %v ", username)
	err := conn.Bind(cf.LdapFirstBindUsername, cf.LdapFirstBindPassword)
	l.Infof("bind  user to query data with user : %v", cf.LdapFirstBindUsername)
	if err != nil {
		l.Errorf("can not bind first user : %v ", err)
		return types.User{}, err
	}
	l.Infof("bind  user success")
	searchRequest := ldap.NewSearchRequest(
		cf.BaseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=*)(cn=%s))", username),
		[]string{"dn", "mail", "memberOf", "description", "displayName", "cn"},
		nil,
	)
	searchresult, err := conn.Search(searchRequest)
	if err != nil {
		l.Errorf("ldap query search fail %v", err)
		return types.User{}, err
	}
	l.Infof("query success")
	if len(searchresult.Entries) == 0 {
		l.Errorf("no user query")
		err := errors.New("user does not exist ")
		return types.User{}, err
	}
	l.Infof("passer user information ")
	for _, entries := range searchresult.Entries {
		if isLower(entries.GetAttributeValue("cn")) {
			emailaddr := entries.GetAttributeValue("mail")
			ds := entries.GetAttributeValue("description")
			// try to get DC name form description value
			dc := getDCFormDescripton(ds, "DC[\\s,\\-]*[0-9]+")
			fulname := entries.GetAttributeValues("displayName")
			var name string
			if len(fulname) == 0 {
				name = ""
			}
			l.Infof("passer user information success ")
			name = fulname[0]
			l.Infof("username: %s, full name : %v , dc: %v, email : %s ", username, name, dc, emailaddr)
			return types.User{
				UserName:       username,
				FullName:       name,
				DeliveryCenter: dc,
				Email:          emailaddr,
			}, nil
		}
	}
	l.Errorf("can't get user information form ldap")
	return types.User{}, errors.New("can't get user information form ldap")
}
func findUserDN(username string, cf LDAPconf, l glog.Logger, conn *ldap.Conn) ([]string, error) {
	l.Infof("find user dn form ldap server by user name : %v ", username)
	err := conn.Bind(cf.LdapFirstBindUsername, cf.LdapFirstBindPassword)
	if err != nil {
		l.Errorf("can't first bind to ldap database to query user dn %v", err)
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
		l.Errorf("search dn for login fail %v ", err)
		return nil, err
	}
	if len(searresult.Entries) == 0 {
		l.Errorf("result is not contain anything")
		err := errors.New("user does not exist")
		return nil, err
	}
	dn := make([]string, len(searresult.Entries))
	for _, stringdn := range searresult.Entries {
		dn = append(dn, stringdn.DN)
	}
	return dn, nil

}

// the method just is used in here
func getDCFormDescripton(s string, r string) []string {
	logger := glog.New().WithField("package", "ldap")
	logger.Infof("the description : %s", s)
	re := regexp.MustCompile(r)
	result := re.FindAllString(s, -1)
	if len(result) == 0 {
		logger.Infof("can not get DC form description string")
		logger.Infof("Default is Wanderer")
		re := []string{"Wanderer"}
		return re
	}
	logger.Infof("get dc success %v ", result)
	return result
}
func isLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) {
			return false
		}
	}
	return true
}
