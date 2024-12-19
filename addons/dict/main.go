package dict

import (
	baseModel "dzhgo/internal/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"

	"github.com/gzdzh-cn/dzhcore"

	"dzhgo/addons/dict/model"

	_ "dzhgo/addons/dict/controller"
	_ "dzhgo/addons/dict/packed"
)

func NewInit() {
	var (
		ctx = gctx.GetInitCtx()
	)
	glog.Debug(ctx, "------------ dict")

	g.Log().Debug(ctx, "addon dict init start ...")
	g.Log().Debugf(ctx, "dict version:%v", Version)
	dzhcore.FillInitData(ctx, "dict", &model.DictInfo{})
	dzhcore.FillInitData(ctx, "dict", &model.DictType{})
	dzhcore.FillInitData(ctx, "dict", &baseModel.BaseSysMenu{})
	g.Log().Debug(ctx, "addon dict init finished ...")
}
