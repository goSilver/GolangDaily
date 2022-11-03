package v6

import (
	"encoding/json"
	"log"
)

type MiddlewareBuilder2 struct {
	logFunc func(accessLog string)
}

func (b *MiddlewareBuilder2) LogFunc2(logFunc func(accessLog string)) *MiddlewareBuilder2 {
	b.logFunc = logFunc
	return b
}

func NewBuilder2() *MiddlewareBuilder2 {
	return &MiddlewareBuilder2{
		logFunc: func(accessLog string) {
			log.Println("B--" + accessLog)
		},
	}
}

type accessLog2 struct {
	Host       string
	Route      string
	HTTPMethod string `json:"http_method"`
	Path       string
}

func (b *MiddlewareBuilder2) Build2() Middleware {
	return func(next HandleFunc) HandleFunc {
		return func(ctx *Context) {
			defer func() {
				l := accessLog2{
					Host:       ctx.Req.Host,
					Route:      ctx.MatchedRoute,
					Path:       ctx.Req.URL.Path,
					HTTPMethod: ctx.Req.Method,
				}
				val, _ := json.Marshal(l)
				b.logFunc(string(val))
			}()
			next(ctx)
		}
	}
}
