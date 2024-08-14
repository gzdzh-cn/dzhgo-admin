package middleware

import (
	"dzhgo/internal/config"
	"github.com/gogf/gf/v2/frame/g"
)

func init() {
	if config.Config.Middleware.Authority.Enable {

		g.Server().BindMiddleware("/app/*/open/*", BaseAuthorityMiddlewareOpen)
		g.Server().BindMiddleware("/app/*/comm/*", BaseAuthorityMiddlewareComm)
		g.Server().BindMiddleware("/admin/*/open/*", BaseAuthorityMiddlewareOpen)
		g.Server().BindMiddleware("/admin/*/comm/*", BaseAuthorityMiddlewareComm)

		g.Server().BindMiddleware("/admin/*", BaseAuthorityMiddleware)
		g.Server().BindMiddleware("/app/*", BaseAuthorityMiddleware)
		g.Server().BindMiddleware("/*", AutoI18n)
		g.Server().BindMiddleware("/*", Exception)

	}
	if config.Config.Middleware.Log.Enable {

		g.Server().BindMiddleware("/admin/*", BaseLog)

	}

}
