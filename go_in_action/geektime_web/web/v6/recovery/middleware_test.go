package recovery

import (
	v6 "gitee.com/geektime-geekbang/geektime-go/go_in_action/geektime_web/web/v6"
	"log"
	"testing"
)

func TestMiddlewareBuilder_Build(t *testing.T) {
	s := v6.NewHTTPServer()
	s.Get("/user", func(ctx *v6.Context) {
		ctx.RespData = []byte("hello, world")
	})

	s.Get("/panic", func(ctx *v6.Context) {
		panic("闲着没事 panic")
	})

	s.Use((&MiddlewareBuilder{
		StatusCode: 500,
		ErrMsg:     "你 Panic 了",
		LogFunc: func(ctx *v6.Context) {
			log.Println(ctx.Req.URL.Path)
		},
	}).Build())

	s.Start(":8081")
}
