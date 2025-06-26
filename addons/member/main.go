package member

import (
	"dzhgo/addons/member/model"
	"dzhgo/internal"
	baseModel "dzhgo/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gzdzh-cn/dzhcore"

	_ "dzhgo/addons/member/controller"
	_ "dzhgo/addons/member/funcs"
	_ "dzhgo/addons/member/middleware"
	_ "dzhgo/addons/member/packed"
)

var (
	ctx = gctx.GetInitCtx()
)

func init() {
	dzhcore.AddAddon(&memberAddon{Version: internal.Version, Name: "member"})
}

type memberAddon struct {
	Version string
	Name    string
}

func (a *memberAddon) GetName() string {
	return a.Name
}

func (a *memberAddon) GetVersion() string {
	return a.Version
}

func (a *memberAddon) NewInit() {
	g.Log().Debug(ctx, "------------ addon member init start ...")
	g.Log().Debugf(ctx, "member version:%v", internal.Version)

	dzhcore.FillInitData(ctx, "member", &model.MemberManage{})
	dzhcore.FillInitData(ctx, "member", &baseModel.BaseSysMenu{})
	g.Log().Debug(ctx, "addon member init end ...")

}
