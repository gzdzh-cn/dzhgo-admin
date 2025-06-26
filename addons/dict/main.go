package dict

import (
	baseModel "dzhgo/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	"github.com/gzdzh-cn/dzhcore"

	"dzhgo/addons/dict/model"

	_ "dzhgo/addons/dict/controller"
	_ "dzhgo/addons/dict/packed"
)

func init() {
	dzhcore.AddAddon(&dictAddon{Version: Version, Name: "dict"})
}

type dictAddon struct {
	Version string
	Name    string
}

func (a *dictAddon) GetName() string {
	return a.Name
}

func (a *dictAddon) GetVersion() string {
	return a.Version
}

func (a *dictAddon) NewInit() {
	var (
		ctx = gctx.GetInitCtx()
	)
	g.Log().Debug(ctx, "------------ addon dict init start ...")
	g.Log().Debugf(ctx, "dict version:%v", Version)
	dzhcore.FillInitData(ctx, "dict", &model.DictInfo{})
	dzhcore.FillInitData(ctx, "dict", &model.DictType{})
	dzhcore.FillInitData(ctx, "dict", &baseModel.BaseSysMenu{})
	g.Log().Debug(ctx, "------------ addon dict init end ...")
}
