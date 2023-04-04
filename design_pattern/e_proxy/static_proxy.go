package e_proxy

import (
	"github.com/labstack/gommon/log"
	"time"
)

// IUser IUser
type IUser interface {
	Login(username, password string)
}

// User 用户
type User struct {
}

// Login 登录接口
func (u *User) Login(username, password string) error {
	// 不实现具体登录逻辑
	return nil
}

// UserProxy 用户代理
type UserProxy struct {
	user *User
}

func NewUserProxy(user *User) *UserProxy {
	return &UserProxy{
		user: user,
	}
}

// Login 代理类的登录实现，这里只做了接口耗时的处理逻辑
func (up *UserProxy) Login(username, password string) error {
	before := time.Now()

	// 这里调用原本的业务逻辑
	err := up.user.Login(username, password)
	if err != nil {
		return err
	}

	log.Printf("用户登录耗时：%s", time.Now().Sub(before))
	return nil
}
