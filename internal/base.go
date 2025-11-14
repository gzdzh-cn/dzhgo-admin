package internal

import (
	"dzhgo/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gzdzh-cn/dzhcore"

	dictModel "dzhgo/addons/dict/model"

	_ "dzhgo/internal/controller"
	_ "dzhgo/internal/funcs"
	_ "dzhgo/internal/middleware"

	_ "dzhgo/internal/packed"
)

var (
	ctx = gctx.GetInitCtx()
)

func init() {

}

func NewInit() {
	g.Log().Debug(ctx, "------------ base init start ...")
	g.Log().Debugf(ctx, "base version:%v", Version)

	dzhcore.FillInitData(ctx, "base", &model.BaseSysMenu{})
	dzhcore.FillInitData(ctx, "base", &model.BaseSysUser{})
	dzhcore.FillInitData(ctx, "base", &model.BaseSysUserRole{})
	dzhcore.FillInitData(ctx, "base", &model.BaseSysRole{})
	dzhcore.FillInitData(ctx, "base", &model.BaseSysRoleMenu{})
	dzhcore.FillInitData(ctx, "base", &model.BaseSysDepartment{})
	dzhcore.FillInitData(ctx, "base", &model.BaseSysRoleDepartment{})
	dzhcore.FillInitData(ctx, "base", &model.BaseSysSetting{})
	// dzhcore.FillInitData(ctx, "base", &model.BaseSysAddons{})
	dzhcore.FillInitData(ctx, "base", &model.BaseSysAddonsTypes{})
	dzhcore.FillInitData(ctx, "base", &dictModel.DictType{})
	dzhcore.FillInitData(ctx, "base", &dictModel.DictInfo{})

	g.Log().Debug(ctx, "------------ base init end ...")
}
