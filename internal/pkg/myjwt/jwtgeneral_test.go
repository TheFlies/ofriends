package myjwt

import (
	"github.com/dgrijalva/jwt-go"
	"reflect"
	"testing"
)

const (
	PRI_KEY               = "AllYourBase"
	TRUE_TOKEN            = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJVc2UiLCJleHAiOjE1NTE0MjM1ODIsImlhdCI6MTU1MTQxOTk4MiwiVXNlaWQiOiJrMTIxMiIsIlVzZWZ1bGxuYW1lIjoia2hhZ24iLCJVc2Vyb2xlIjoiZW5nIn0.0MCpBuDz9nvAUjIF5zfBR0kTF1jEH7eQoD-wll006Ac"
	IN_CORRECT_SIGN_TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJVc2UiLCJleHAiOjE1NTE0MDYzMjcsImlhdCI6MTU1MTQwNjIwNywiVXNlaWQiOiJrMTIxMiIsIlVzZWZ1bGxuYW1lIjoia2hhZ24iLCJVc2Vyb2xlIjoiZW5nIn0.tbzd0JvpwQhQZvz-UOM2w9Nt30j8yO7HMJ65vsZB3h8"
)
// in orther time this test case will fail
// becasue this token wiil expired
func TestCheckToken(t *testing.T) {

	trueclaims := Customeclaims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  1551410104,
			Audience:  "Use",
			ExpiresAt: 1551410224,
		}, Useid: "k1212",
		Usefullname: "khagn",
		Userole:     "eng",
	}
	falseclaims := Customeclaims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  1551350357,
			Audience:  "Use",
			ExpiresAt: 1551352157,
		}, Useid: "33",
		Usefullname: "4444",
		Userole:     "pm",
	}
	testmatrix := []struct {
		token                    string
		expectpayload            Customeclaims
		isvalib                  bool
		expectvaluechecktowclaim bool
		expectcheckisvalib       bool
	}{
		{TRUE_TOKEN, trueclaims, true, true, true},
		{TRUE_TOKEN, falseclaims, true, false, true},
		{IN_CORRECT_SIGN_TOKEN, trueclaims, false, true, false},
	}

	for i, testcase := range testmatrix {
		payload, ok := CheckToken(testcase.token, PRI_KEY)
		if reflect.DeepEqual(payload, testcase.expectpayload) != testcase.expectvaluechecktowclaim && ok != testcase.expectcheckisvalib {
			t.Errorf("The test case: %d get payload from token is fail, payload is geted :%v, \n expect token %v \n", i, payload, testcase.expectpayload)

		}
	}

}
func TestGetPayload(t *testing.T) {
	trueclaims := Customeclaims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  1551410104,
			Audience:  "Use",
			ExpiresAt: 1551410224,
		}, Useid: "k1212",
		Usefullname: "khagn",
		Userole:     "eng",
	}
	re := GetPayload(IN_CORRECT_SIGN_TOKEN, PRI_KEY)
	if !reflect.DeepEqual(re,trueclaims){
		t.Errorf("fails to get payload")
	}
}
func TestCheckExp(t *testing.T) {
	// test expir
	initoken2 := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJVc2UiLCJleHAiOjE1NTEzNTIxMjgsImlhdCI6MTU1MTM1MjEyMywiVXNlaWQiOiJrMTIxMiIsIlVzZWZ1bGxuYW1lIjoia2hhZ24iLCJVc2Vyb2xlIjoiZW5nIn0.MXH6BTzDq3IwPNrnQh1EZV8G4eMGr8u58hiokC5JfBk"
	payload2, ok := CheckToken(initoken2, PRI_KEY)

	if ok {
		t.Error("Fails test exp token")
		t.Error(payload2)
	}
}
func TestCreateToken(t *testing.T) {
	token, err := CreateToken(1551350357, "Use", 1551352157, PRI_KEY, "k1212", "khagn", "eng")
	expecttoken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJVc2UiLCJleHAiOjE1NTEzNTIxNTcsImlhdCI6MTU1MTM1MDM1NywiVXNlaWQiOiJrMTIxMiIsIlVzZWZ1bGxuYW1lIjoia2hhZ24iLCJVc2Vyb2xlIjoiZW5nIn0.dtWYMFdW75wsNfPk68qGH_gYre9WI7nv5pVnz2OrGBE"
	if err != nil {
		t.Error("Fail to create token")
	}
	if token != expecttoken {
		t.Errorf("the result = %s, but expect = : %s", token, expecttoken)
	}
}
