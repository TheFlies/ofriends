package login

import (
	"fmt"
	"os"
	"testing"
)

func TestLDAPAuthentication(t *testing.T) {
	os.Setenv("LDAP_ADDR", "tma-pdc.tma.com.vn")
	os.Setenv("LDAP_PORT", "389")
	os.Setenv("LDAP_FIRST_BIND_USERNAME", "CN=pdkhang,OU=Admin_Privileges,OU=Revoke,DC=tma,DC=com,DC=vn")
	os.Setenv("LDAP_FIRST_BIND_USERNAME_PASSWORD", "9e3#duXc123")
	os.Setenv("BASEDN", "ou=TMAuser,dc=tma,dc=com,dc=vn")
	user, err := LDAPAuthentication("dhnhan", "fZ31iXyO%##")
	fmt.Println(user)
	if err != nil {
		t.Error(err)
	}

}
