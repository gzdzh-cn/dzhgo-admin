package middleware

import (
	"dzhgo/internal/config"

	"github.com/gogf/gf/v2/frame/g"
)

func init() {

	var s = g.Server()
	//路由鉴权开启
	if config.Cfg.Middleware.Authority.Enable {

		if config.Cfg.Middleware.Cors {
			s.BindMiddleware("/app/*", AddonAuthorityMiddleware)
			s.BindMiddleware("/admin/*", AddonAuthorityMiddleware)
		}

		s.BindMiddleware("/app/*/open/*", BaseAuthorityMiddlewareOpen) // 开放接口
		s.BindMiddleware("/app/*/comm/*", BaseAuthorityMiddlewareComm) // 需登录接口
		s.BindMiddleware("/admin/*/open/*", BaseAuthorityMiddlewareOpen)
		s.BindMiddleware("/admin/*/comm/*", BaseAuthorityMiddlewareComm)

		s.BindMiddleware("/admin/*", BaseAuthorityMiddleware)
		s.BindMiddleware("/app/*", BaseAuthorityMiddleware)
		s.BindMiddleware("/admin/*", AutoI18n)  //
		s.BindMiddleware("/admin/*", Exception) //异常抛出捕获
	}

	//请求日志记录到数据库开启
	if config.Cfg.Middleware.Log.Enable {
		s.BindMiddleware("/admin/*", BaseLog)
		s.BindMiddleware("/app/*", BaseLog)
	}

}
