package ldap

import (
	"fmt"
	"log"
	"strings"

	"gopkg.in/ldap.v3"

	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/pkg/config/env"
)

type (
	LDAPconf struct {
		LdapAddr              string `envconfig:"LDAP_ADDR"`
		LdapPort              int    `envconfig:"LDAP_PORT"`
		LdapFirstBindUsername string `envconfig:"LDAP_FIRST_BIND_USERNAME"`
		LdapFirstBindPassword string `envconfig:"LDAP_FIRST_BIND_PASSWORD"`
		BaseDN                string `envconfig:"BASEDN"`
	}
)

func LDAPAuthentication(username string, password string) error {
	var conf LDAPconf
	envconfig.Load(&conf)
	ldapconn, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", conf.LdapAddr, conf.LdapPort))
	if err != nil {
		log.Fatal(err)
	}
	defer ldapconn.Close()
	var sb strings.Builder
	sb.WriteString("Cn=")
	sb.WriteString(username)
	sb.WriteString(",")
	sb.WriteString(conf.BaseDN)
	err = ldapconn.Bind(sb.String(), password)
	if err != nil {
		return err
	}
	return nil
}
func LDAPQuery(username string) (types.Usercache, error) {
	var conf LDAPconf
	envconfig.Load(&conf)
	ldapconn, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", conf.LdapAddr, conf.LdapPort))
	if err != nil {
		return types.Usercache{}, err
	}
	defer ldapconn.Close()
	err = ldapconn.Bind(conf.LdapFirstBindUsername, conf.LdapFirstBindPassword)
	if err != nil {
		return types.Usercache{}, err
	}
	searchRequest := ldap.NewSearchRequest(
		conf.BaseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=*)(cn=%s))", username),
		[]string{"dn", "mail", "memberOf", "description", "displayName"},
		nil,
	)
	searchrequest, err := ldapconn.Search(searchRequest)
	if err != nil {
		fmt.Println("search fails")
		return types.Usercache{}, err
	}
	emailaddr := searchrequest.Entries[0].GetAttributeValues("mail")
	role := searchrequest.Entries[0].GetAttributeValues("description")
	fulname := searchrequest.Entries[0].GetAttributeValues("displayName")
	return types.Usercache{
		Username: username,
		Fullname: fulname[0],
		Userrole: role[0],
		Email:    emailaddr[0],
	}, nil
}
