package e_proxy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserProxy_Login(t *testing.T) {
	testCases := []struct {
		name     string
		username string
		password string
		want     string
		wantErr  bool
	}{
		{
			name:     "测试代理模式",
			username: "root",
			password: "root",
			want:     "",
			wantErr:  false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			user := &User{}
			userProxy := NewUserProxy(user)
			err := userProxy.Login(testCase.username, testCase.password)
			assert.Equal(t, testCase.wantErr, err != nil)
		})
	}
}
