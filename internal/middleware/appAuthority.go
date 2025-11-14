package middleware

import (
	"dzhgo/internal/config"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gzdzh-cn/dzhcore"
)

// 本类接口无需权限验证
func AppAuthorityMiddlewareOpen(r *ghttp.Request) {
	r.SetCtxVar("AuthOpen", true)
	r.Middleware.Next()
}

// 本类接口无需权限验证,只需登录验证
func AppAuthorityMiddlewareComm(r *ghttp.Request) {
	r.SetCtxVar("AuthComm", true)
	r.Middleware.Next()
}

// 其余接口需登录验证同时需要权限验证
func AppAuthorityMiddleware(r *ghttp.Request) {

	var (
		statusCode = 200
		ctx        = r.GetCtx()
	)

	// 无需登录验证
	AuthOpen := r.GetCtxVar("AuthOpen", false)
	if AuthOpen.Bool() {
		r.Middleware.Next()
		return
	}

	tokenString := r.GetHeader("Authorization")
	token, err := jwt.ParseWithClaims(tokenString, &dzhcore.AppClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(config.Cfg.Jwt.Secret), nil
	})
	if err != nil {
		g.Log().Error(ctx, "BaseAuthorityMiddleware", err)
		statusCode = 401
		r.Response.WriteStatusExit(statusCode, g.Map{
			"code":    1001,
			"message": "登陆失效～",
		})
	}
	if !token.Valid {
		// g.Log().Error(ctx, "BaseAuthorityMiddleware", "token invalid")
		statusCode = 401
		r.Response.WriteStatusExit(statusCode, g.Map{
			"code":    1001,
			"message": "登陆失效～",
		})
	}

	member := token.Claims.(*dzhcore.AppClaims)
	//g.Log().Debugf(ctx, "admin %v", gconv.String(admin))
	// 将用户信息放入上下文
	r.SetCtxVar("member", member)

	cachetoken, _ := dzhcore.CacheManager.Get(ctx, "member:token:"+gconv.String(member.MemberId))
	rtoken := cachetoken.String()

	if tokenString != rtoken {
		statusCode = 401
		r.Response.WriteStatusExit(statusCode, g.Map{
			"code":    1001,
			"message": "登陆失效～",
		})
	}

	r.Middleware.Next()
}
