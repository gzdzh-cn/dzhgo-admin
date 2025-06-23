package middleware

import (
	logic "dzhgo/internal/logic/sys"
	"dzhgo/internal/utils"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"runtime"
	"time"
)

// 请求日志记录到数据库开启
func BaseLog(r *ghttp.Request) {
	var (
		ctx               = r.GetCtx()
		BaseSysLogService = logic.NewsBaseSysLogService()
		path              = r.URL.Path
	)

	//忽略正则规则和指定的请求
	ignorePathSlice := g.Cfg().MustGet(ctx, "modules.base.middleware.log.ignorePath").String()
	ignoreReg := g.Cfg().MustGet(ctx, "modules.base.middleware.log.ignoreReg").String()
	ignorePath := garray.NewStrArrayFrom(gstr.Split(ignorePathSlice, ","))
	openReg := `^(/admin/.*/open/.*|/app/.*/open/.*)$` //过滤不鉴权的接口
	if !ignorePath.Contains(path) && !gregex.IsMatch(ignoreReg, []byte(path)) && !gregex.IsMatch(openReg, []byte(path)) {
		//g.Log().Infof(ctx, "写入日志,路径:%v", path)
		BaseSysLogService.Record(ctx)
	}

	r.Middleware.Next()

}

// 请求日志运行明细开启
func RunLog(r *ghttp.Request) {
	var (
		startTime     = time.Now() //请求进入时间
		ctx           = r.Context()
		memStatsStart runtime.MemStats
	)
	runtime.ReadMemStats(&memStatsStart)

	r.Middleware.Next()

	//日志打印运行时间
	utils.StdOutLog(ctx, startTime, memStatsStart)
}
