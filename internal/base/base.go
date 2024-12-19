package base

import (
	"dzhgo/internal/config"
	"dzhgo/internal/model"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gzdzh-cn/dzhcore"

	"github.com/gogf/gf/v2/os/gctx"

	_ "dzhgo/internal/controller"
	_ "dzhgo/internal/funcs"
	_ "dzhgo/internal/middleware"

	_ "dzhgo/internal/packed"
)

var (
	ctx = gctx.GetInitCtx()
)

func NewInit() {
	glog.Debug(ctx, "------------ base NewFilter")
	glog.Debug(ctx, "------------ base init start ...")
	glog.Debugf(ctx, "------------ base version:%v", config.Version)

	dzhcore.FillInitData(ctx, "base", &model.BaseSysMenu{})
	dzhcore.FillInitData(ctx, "base", &model.BaseSysUser{})
	dzhcore.FillInitData(ctx, "base", &model.BaseSysUserRole{})
	dzhcore.FillInitData(ctx, "base", &model.BaseSysRole{})
	dzhcore.FillInitData(ctx, "base", &model.BaseSysRoleMenu{})
	dzhcore.FillInitData(ctx, "base", &model.BaseSysDepartment{})
	dzhcore.FillInitData(ctx, "base", &model.BaseSysRoleDepartment{})
	dzhcore.FillInitData(ctx, "base", &model.BaseSysSetting{})
	dzhcore.FillInitData(ctx, "base", &model.BaseSysAddons{})
	dzhcore.FillInitData(ctx, "base", &model.BaseSysAddonsType{})

	glog.Debug(ctx, "------------ base init finished ...")
}
