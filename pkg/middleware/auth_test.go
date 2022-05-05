package middleware

import (
	"testing"
)

func TestNewToken(t *testing.T) {
	u := User{
		UserID:   1,
		Username: "xantares",
		Password: "qwer",
		Email:    "aaa@ff.com",
		Avatar:   "noavatar",
	}

	token, err := NewToken(u)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	t.Log(token)
	println(token)
}
