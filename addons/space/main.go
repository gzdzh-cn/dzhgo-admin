package space

import (
	"dzhgo/internal"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gzdzh-cn/dzhcore"
	// _ "dzhgo/addons/space/controller"
	// _ "dzhgo/addons/space/middleware"
)

func init() {
	dzhcore.AddAddon(&spaceAddon{Version: internal.Version, Name: "space"})
}

type spaceAddon struct {
	Version string
	Name    string
}

func (a *spaceAddon) GetName() string {
	return a.Name
}

func (a *spaceAddon) GetVersion() string {
	return a.Version
}

func (a *spaceAddon) NewInit() {
	var (
		ctx = gctx.GetInitCtx()
	)
	g.Log().Debug(ctx, "------------ addon space init start ...")
	g.Log().Debug(ctx, "------------ addon space init end ...")
}
