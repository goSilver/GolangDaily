package errhdl

import "gitee.com/geektime-geekbang/geektime-go/go_in_action/geektime_web/web/v6"

type MiddlewareBuilder struct {
	// 错误状态码及对应返回页面
	resp map[int][]byte
}

func NewMiddlewareBuilder() *MiddlewareBuilder {
	return &MiddlewareBuilder{
		// 这里可以非常大方，因为在预计中用户会关心的错误码不可能超过 64
		resp: make(map[int][]byte, 64),
	}
}

func (m *MiddlewareBuilder) RegisterError(code int, resp []byte) *MiddlewareBuilder {
	m.resp[code] = resp
	return m
}

func (m *MiddlewareBuilder) Build() v6.Middleware {
	return func(next v6.HandleFunc) v6.HandleFunc {
		return func(ctx *v6.Context) {
			next(ctx)

			// 如果中间件中保存了对应状态码的返回页面则直接返回对应页面
			resp, ok := m.resp[ctx.RespStatusCode]
			if ok {
				ctx.RespData = resp
			}
		}
	}
}
