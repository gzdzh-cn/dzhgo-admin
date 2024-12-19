package space

import (
	_ "dzhgo/addons/space/controller"
	_ "dzhgo/addons/space/middleware"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func NewInit() {
	var (
		ctx = gctx.GetInitCtx()
	)
	g.Log().Debug(ctx, "module space init start ...")
	g.Log().Debug(ctx, "module space init finished ...")
}
