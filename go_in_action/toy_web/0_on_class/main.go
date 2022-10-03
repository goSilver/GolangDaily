package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是主页")
}

func user(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是用户")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是创建用户")
}

func order(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是订单")
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	req := &signUpReq{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "read body failed:%v", err)
		return
	}
	err = json.Unmarshal(body, req)
	if err != nil {
		fmt.Fprintf(w, "deserialized failed:%v", err)
		return
	}
	// 返回一个虚拟的user id，表示注册成功
	fmt.Fprintf(w, "%d", err)
}

func SignUpV2(w http.ResponseWriter, r *http.Request) {
	c := NewContext(w, r)
	req := &signUpReq{}
	err := c.ReadJson(req)
	// 返回一个虚拟的user id，表示注册成功
	fmt.Fprintf(w, "%d", err)
}

func main() {
	server := NewSdkHttpServer("my-test-server")
	server.Route("/", home)
	server.Route("/user", user)
	server.Route("/user/create", createUser)
	server.Route("/order", order)

	server.Start("8080")
	//http.HandleFunc("/", home)
	//http.HandleFunc("/user", user)
	//http.HandleFunc("/user/create", createUser)
	//http.HandleFunc("/order", order)
	//http.ListenAndServe(":8080", nil)
}

// Server 是http server 的顶级抽象
type Server interface {
	// Route 设计一个路由，命中该路由的会执行handleFunc的代码
	Route(pattern string, handlerFunc http.HandlerFunc)
	Start(address string) error
}

type sdkHttpServer struct {
	Name string
}

func (s *sdkHttpServer) Route(pattern string, handlerFunc http.HandlerFunc) {
	panic("implement me")
}

func (s *sdkHttpServer) Start(address string) error {
	panic("implement me")
}

func NewSdkHttpServer(name string) Server {
	return &sdkHttpServer{
		Name: name,
	}
}

type signUpReq struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		W: w,
		R: r,
	}
}

func (c *Context) ReadJson(data interface{}) error {
	body, err := io.ReadAll(c.R.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, data)
}

func (c *Context) WriteJson(status int, data interface{}) error {
	bs, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = c.W.Write(bs)
	if err != nil {
		return err
	}
	c.W.WriteHeader(status)
	return nil
}

func (c *Context) OkJson(data interface{}) error {
	return c.WriteJson(http.StatusOK, data)
}

func (c *Context) SystemErrJson(data interface{}) error {
	return c.WriteJson(http.StatusInternalServerError, data)
}

func (c *Context) BadRequestJson(data interface{}) error {
	return c.WriteJson(http.StatusBadRequest, data)
}
