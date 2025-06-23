package middleware

import (
	"dzhgo/internal/config"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func init() {
	var ctx = gctx.GetInitCtx()
	//路由鉴权开启
	if config.Cfg.Middleware.Authority.Enable {

		if config.Cfg.Middleware.Cors {
			g.Log().Info(ctx, "开启跨域")
			g.Server().BindMiddleware("/app/*", AddonAuthorityMiddleware)
			g.Server().BindMiddleware("/admin/*", AddonAuthorityMiddleware)
		}

		g.Server().BindMiddleware("/app/*/open/*", BaseAuthorityMiddlewareOpen) // 开放接口
		g.Server().BindMiddleware("/app/*/comm/*", BaseAuthorityMiddlewareComm) // 需登录接口
		g.Server().BindMiddleware("/admin/*/open/*", BaseAuthorityMiddlewareOpen)
		g.Server().BindMiddleware("/admin/*/comm/*", BaseAuthorityMiddlewareComm)

		g.Server().BindMiddleware("/admin/*", BaseAuthorityMiddleware)
		g.Server().BindMiddleware("/app/*", BaseAuthorityMiddleware)
		g.Server().BindMiddleware("/admin/*", AutoI18n)  //
		g.Server().BindMiddleware("/admin/*", Exception) //异常抛出捕获
	}

	//请求日志记录到数据库开启
	if config.Cfg.Middleware.Log.Enable {
		g.Server().BindMiddleware("/admin/*", BaseLog)
		g.Server().BindMiddleware("/app/*", BaseLog)
	}

	//请求日志运行明细开启
	if config.Cfg.Middleware.RunLogger.Enable {
		g.Server().BindMiddleware("/admin/*", RunLog) //请求日志明细
		g.Server().BindMiddleware("/app/*", RunLog)   //请求日志明细
	}

}
