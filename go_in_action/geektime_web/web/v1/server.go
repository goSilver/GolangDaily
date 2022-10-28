package v1

import "net/http"

type HandleFunc func(ctx *Context)

type Server interface {
	http.Handler
	// Start 启动服务器
	// addr是监听地址。如果只指定端口，可以使用":8081"或者"localhost:8082"
	Start(addr string) error
	// AddRoute 注册一个路由
	// method是HTTP方法；path是路径
	AddRoute(method, path string, handler HandleFunc)
	// 不采取该实现，原因：按照规范，一个path对应一个handler更为清晰；多个handler的执行顺序、事务问题；
	//addRoute(method, path string, handler... HandleFunc)
}

// 确保 HTTPServer 肯定实现了 Server 接口
var _ Server = &HTTPServer{}

type HTTPServer struct {
}

func (s *HTTPServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx := &Context{
		Req:  request,
		Resp: writer,
	}
	s.serve(ctx)
}

func (s *HTTPServer) Start(addr string) error {
	return http.ListenAndServe(addr, s)
}

func (s *HTTPServer) Post(path string, handler HandleFunc) {
	s.AddRoute(http.MethodPost, path, handler)
}

func (s *HTTPServer) Get(path string, handler HandleFunc) {
	s.AddRoute(http.MethodGet, path, handler)
}

func (s *HTTPServer) AddRoute(method, path string, handler HandleFunc) {
	panic("implement me")
}

func (s *HTTPServer) serve(ctx *Context) {

}
