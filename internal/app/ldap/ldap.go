package ldap

import (
	"fmt"
	"regexp"
	"unicode"

	"github.com/globalsign/mgo/bson"
	"github.com/pkg/errors"
	"gopkg.in/ldap.v3"

	"github.com/TheFlies/ofriends/internal/app/api/handler/login"
	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/pkg/config/env"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
	"github.com/TheFlies/ofriends/internal/pkg/jwt"
)

type (
	LDAPconf struct {
		LDAPAddr              string `envconfig:"LDAP_ADDR"`
		LDAPPort              int    `envconfig:"LDAP_PORT"`
		LDAPFirstBindUsername string `envconfig:"LDAP_FIRST_BIND_USERNAME"`
		LDAPFirstBindPassword string `envconfig:"LDAP_FIRST_BIND_PASSWORD"`
		BaseDN                string `envconfig:"BASEDN"`
	}
	LdapAuthentication struct {
		log      glog.Logger
		conf     LDAPconf
		userSvr  login.UserService
		LDAPConn *ldap.Conn
	}
)

func New(svr login.UserService) *LdapAuthentication {
	var LDAPConf LDAPconf
	log := glog.New()
	l := log.WithField("package", "ldap")
	envconfig.Load(&LDAPConf)
	return &LdapAuthentication{
		log:     l,
		conf:    LDAPConf,
		userSvr: svr,
	}
}
func (l *LdapAuthentication) ConnectToServer() {
	l.log.Infof("connecting to ldap server with addr: %v , port : %v", l.conf.LDAPAddr, l.conf.LDAPPort)
	connection, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", l.conf.LDAPAddr, l.conf.LDAPPort))
	if err != nil {
		l.log.Errorf("connect to ldap server fails %v", err)
	}
	l.LDAPConn = connection
}
func (l *LdapAuthentication) Close() {
	err := l.LDAPConn.UnauthenticatedBind(l.conf.LDAPFirstBindUsername)
	if err != nil {
		l.log.Errorf("can not unbind user : %s ", l.conf.LDAPFirstBindUsername)
	}
	l.log.Infof("unbind user")
	l.LDAPConn.Close()
	l.log.Infof("close a connection to : %v", l.conf.LDAPAddr)
}
func (l *LdapAuthentication) Authenticate(username string, password string) (interface{}, error) {
	l.ConnectToServer()
	defer l.Close()
	err := ldapAuthentication(username, password, l.conf, l.log, l.LDAPConn)
	JWTGeneration := jwt.NewTokenGeneration()
	respondMap := make(map[string]interface{})
	if err != nil {
		l.log.Errorf("login fail %v ", err)
		return "", err
	}
	l.log.Infof("login success, waiting query user information")
	isExists := l.userSvr.CheckExistence(username)
	if isExists {
		l.log.Infof("query user information form database")
		DBUser, err := l.userSvr.GetByName(username)
		if err != nil {
			l.log.Errorf("get user fail: %v", err)
			return "", err
		}

		JWTToken, err := JWTGeneration.CreateToken(DBUser.Username, DBUser.FullName, DBUser.DeliveryCenter)
		if err != nil {
			l.log.Errorf("get user fail : %v", err)
			return "", err
		}
		returnUser := types.User{
			ID:             bson.NewObjectId(),
			Username:       DBUser.Username,
			FullName:       DBUser.FullName,
			Email:          DBUser.Email,
			DeliveryCenter: DBUser.DeliveryCenter,
		}
		l.log.Infof("returning token")
		respondMap["userInformation"] = returnUser
		respondMap["token"] = JWTToken
		return respondMap, nil
	}
	l.log.Infof("query user information in ldap server ")
	LDAPUser, err := ldapQuery(username, l.conf, l.log, l.LDAPConn)
	if err != nil {
		l.log.Errorf("get user fail")
		return "", err
	}
	LDAPUser.ID = bson.NewObjectId()
	err = l.userSvr.AddUser(&LDAPUser)
	if err != nil {
		l.log.Errorf("can not add a user to mongodb %v", err)
	}
	JWTToken, err := JWTGeneration.CreateToken(LDAPUser.Username, LDAPUser.FullName, LDAPUser.DeliveryCenter)
	if err != nil {
		l.log.Errorf("get token fail : %v", err)
		return "", err
	}
	l.log.Infof("returning token")
	respondMap["userInformation"] = LDAPUser
	respondMap["token"] = JWTToken
	return respondMap, nil

}
func ldapAuthentication(username string, password string, cf LDAPconf, l glog.Logger, conn *ldap.Conn) error {
	userDN, err := findUserDN(username, cf, l, conn)
	l.Infof("login with user dn : %v and lens : %d", userDN, len(userDN))
	if err != nil {
		l.Errorf("find user dn fail %v", err)
		return err
	}

	//try login with list dn is queried
	//if all userDN fails, login will fail
	for i, dn := range userDN {
		l.Infof("authenticating with id: %d dn : %v", i, dn)
		if dn == "" {
			continue
		}
		err = conn.Bind(dn, password)
		if err == nil {
			l.Infof("dn login %v pass: %v", dn, password)
			l.Errorf("login successfully with dn: %v", dn)
			return nil
		}
	}

	l.Errorf("all user dn is fail")
	return errors.New("login fails")

}
func ldapQuery(username string, cf LDAPconf, l glog.Logger, conn *ldap.Conn) (types.User, error) {
	l.Infof("query user information form ldap server by user name : %v ", username)
	err := conn.Bind(cf.LDAPFirstBindUsername, cf.LDAPFirstBindPassword)
	l.Infof("bind  user to query data with user : %v", cf.LDAPFirstBindUsername)
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
	searchResult, err := conn.Search(searchRequest)
	if err != nil {
		l.Errorf("ldap query search fail %v", err)
		return types.User{}, err
	}
	l.Infof("query success")
	if len(searchResult.Entries) == 0 {
		l.Errorf("no user query")
		err := errors.New("user does not exist ")
		return types.User{}, err
	}
	l.Infof("passer user information ")
	for _, entries := range searchResult.Entries {
		if isUserInfoCN(entries.GetAttributeValue("cn")) {
			emailAddress := entries.GetAttributeValue("mail")
			description := entries.GetAttributeValue("description")
			// try to get DC name form description value
			deliveryCenter := getDCFormDescription(description, "DC[\\s,\\-]*[0-9]+", l)
			fulName := entries.GetAttributeValues("displayName")
			var name string
			if len(fulName) == 0 {
				name = ""
			}
			l.Infof("passer user information success ")
			name = fulName[0]
			l.Infof("username: %s, full name : %v , delivery center: %v, email : %s ", username, name, deliveryCenter, emailAddress)
			return types.User{
				Username:       username,
				FullName:       name,
				DeliveryCenter: deliveryCenter,
				Email:          emailAddress,
			}, nil
		}
	}
	l.Errorf("can't get user information form ldap")
	return types.User{}, errors.New("can't get user information form ldap")
}
func findUserDN(username string, cf LDAPconf, l glog.Logger, conn *ldap.Conn) ([]string, error) {
	l.Infof("find user distinguishedName form ldap server by user name : %v ", username)
	err := conn.Bind(cf.LDAPFirstBindUsername, cf.LDAPFirstBindPassword)
	if err != nil {
		l.Errorf("can't first bind to ldap database to query user distinguishedName %v", err)
		return nil, err
	}
	l.Infof("query user by base dn : %s", cf.BaseDN)
	searchRequest := ldap.NewSearchRequest(
		cf.BaseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=*)(cn=%s))", username),
		[]string{"distinguishedName"},
		nil,
	)
	searchResult, err := conn.Search(searchRequest)
	if err != nil {
		l.Errorf("search distinguishedName for login fail %v ", err)
		return nil, err
	}
	if len(searchResult.Entries) == 0 {
		l.Errorf("result is not contain anything")
		err := errors.New("user does not exist")
		return nil, err
	}
	l.Infof("len of distinguishedName: %d", len(searchResult.Entries))
	distinguishedName := make([]string, len(searchResult.Entries))

	for _, dn := range searchResult.Entries {
		l.Infof("dn is appended %v ", dn.DN)
		distinguishedName = append(distinguishedName, dn.DN)
	}
	return distinguishedName, nil

}

// the method just is used in here
func getDCFormDescription(s string, r string, l glog.Logger) []string {
	l.Infof("the description : %s", s)
	re := regexp.MustCompile(r)
	result := re.FindAllString(s, -1)
	if len(result) == 0 {
		l.Infof("can not get DC form description string")
		l.Infof("Default is Wanderer")
		re := []string{"Wanderer"}
		return re
	}
	l.Infof("get dc success %v ", result)
	return result
}
func isUserInfoCN(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) {
			return false
		}
	}
	return true
}
