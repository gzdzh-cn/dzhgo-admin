package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

func Exception(r *ghttp.Request) {

	r.Middleware.Next()

}
