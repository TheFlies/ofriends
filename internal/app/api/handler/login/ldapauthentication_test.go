package login

import (
	"fmt"
	"os"
	"testing"
)

func TestLDAPAuthentication(t *testing.T) {
	os.Setenv("LDAP_ADDR", "***REMOVED***")
	os.Setenv("LDAP_PORT", "389")
	os.Setenv("LDAP_FIRST_BIND_USERNAME", "***REMOVED***")
	os.Setenv("LDAP_FIRST_BIND_USERNAME_PASSWORD", "***REMOVED***")
	os.Setenv("BASEDN", "ou=TMAuser,dc=tma,dc=com,dc=vn")
	user, err := LDAPAuthentication("dhnhan", "fZ31iXyO%##")
	fmt.Println(user)
	if err != nil {
		t.Error(err)
	}

}
