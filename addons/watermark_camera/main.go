package watermark_camera

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	_ "dzhgo/addons/watermark_camera/controller"
	_ "dzhgo/addons/watermark_camera/funcs"
	_ "dzhgo/addons/watermark_camera/middleware"
	_ "dzhgo/addons/watermark_camera/packed"
)

func NewInit() {
	var (
		ctx = gctx.GetInitCtx()
	)

	g.Log().Debug(ctx, "addon watermark_camera init start ...")
	g.Log().Debugf(ctx, "watermark_camera version:%v", Version)

	g.Log().Debug(ctx, "addon watermark_camera init finished ...")

	g.Log().Debug(ctx, "module task init finished ...")

}
