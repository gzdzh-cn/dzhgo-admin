package middleware

import (
	logic "dzhgo/internal/logic/sys"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
)

func BaseLog(r *ghttp.Request) {
	var (
		ctx               = r.GetCtx()
		BaseSysLogService = logic.NewsBaseSysLogService()
		path              = r.URL.Path
	)

	//忽略正则规则和指定的请求
	ignorePathSlice := g.Cfg().MustGet(ctx, "modules.base.middleware.authority.ignorePath").String()
	ignoreReg := g.Cfg().MustGet(ctx, "modules.base.middleware.authority.ignoreReg").String()
	ignorePath := garray.NewStrArrayFrom(gstr.Split(ignorePathSlice, ","))

	if !ignorePath.Contains(path) && !gregex.IsMatch(ignoreReg, []byte(path)) {
		BaseSysLogService.Record(ctx)
	}

	r.Middleware.Next()

}
