package myjwt

import (
	"testing"
	"time"
)

const (
	PRI_KEY = "AllYourBase"
)

func TestCheckExp(t *testing.T) {
	// test expir
	initoken2 := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJVc2UiLCJleHAiOjE1NTEzNTIxMjgsImlhdCI6MTU1MTM1MjEyMywiVXNlaWQiOiJrMTIxMiIsIlVzZWZ1bGxuYW1lIjoia2hhZ24iLCJVc2Vyb2xlIjoiZW5nIn0.MXH6BTzDq3IwPNrnQh1EZV8G4eMGr8u58hiokC5JfBk"
	ok := CheckExp(initoken2, PRI_KEY, time.Now().Unix())
	if ok {
		t.Error("Fails test exp token")

	}
}
func TestCreateToken(t *testing.T) {
	token, err := CreateToken("k1212", "k1212", "eng", 1551350357, 1551352157, PRI_KEY)
	expecttoken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJVc2UiLCJleHAiOjE1NTEzNTIxNTcsImlhdCI6MTU1MTM1MDM1NywiVXNlaWQiOiJrMTIxMiIsIlVzZWZ1bGxuYW1lIjoia2hhZ24iLCJVc2Vyb2xlIjoiZW5nIn0.dtWYMFdW75wsNfPk68qGH_gYre9WI7nv5pVnz2OrGBE"
	if err != nil {
		t.Error("Fail to create token")
	}
	if token != expecttoken {
		t.Errorf("the result = %s, but expect = : %s", token, expecttoken)
	}
}
