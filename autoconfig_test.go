package go_session_auto_config

import (
	"github.com/zj-kenzhou/go-session"
	"github.com/zj-kenzhou/go-session/model"
	"testing"
)

func TestAutoConfig(t *testing.T) {
	token, err := session.Login(model.LoginModel{
		LoginId:   "123",
		Ip:        "127.0.0.1",
		Device:    "pc",
		UserAgent: "test",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
}
