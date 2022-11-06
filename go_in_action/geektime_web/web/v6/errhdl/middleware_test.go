package errhdl

import (
	"bytes"
	"gitee.com/geektime-geekbang/geektime-go/go_in_action/geektime_web/web/v6"
	"html/template"
	"testing"
)

func TestNewMiddlewareBuilder(t *testing.T) {
	s := v6.NewHTTPServer()
	s.Get("/user", func(ctx *v6.Context) {
		ctx.RespData = []byte("hello, world")
	})

	page := `
<html>
	<h1>404 NOT FOUND</h1>
</html>
`
	tpl, err := template.New("404").Parse(page)
	if err != nil {
		t.Fatal(err)
	}
	buffer := &bytes.Buffer{}
	err = tpl.Execute(buffer, nil)
	if err != nil {
		t.Fatal(err)
	}
	// 这里引入错误页面中间件，如果返回状态码=404，则返回固定页面
	s.Use(NewMiddlewareBuilder().RegisterError(404, buffer.Bytes()).Build())
	s.Start(":8081")
}
