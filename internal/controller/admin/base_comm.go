package admin

import (
	"context"
	v1 "dzhgo/internal/api/admin_v1"
	"dzhgo/internal/common"
	"dzhgo/internal/service"

	"github.com/gzdzh-cn/dzhcore"

	"github.com/gogf/gf/v2/frame/g"
)

type BaseCommController struct {
	*dzhcore.ControllerSimple
}

func init() {
	var baseCommController = &BaseCommController{
		ControllerSimple: &dzhcore.ControllerSimple{
			Prefix: "/admin/base/comm",
		},
	}
	// 注册路由
	dzhcore.AddControllerSimple(baseCommController)
}

// 会员数据
func (c *BaseCommController) Person(ctx context.Context, req *v1.BaseCommPersonReq) (res *dzhcore.BaseRes, err error) {
	admin := common.GetAdmin(ctx)
	data, err := service.BaseSysUserService().Person(admin.UserId)
	res = dzhcore.Ok(data)
	return
}

// 权限菜单
func (c *BaseCommController) Permmenu(ctx context.Context, req *v1.BaseCommPermmenuReq) (res *dzhcore.BaseRes, err error) {

	admin := common.GetAdmin(ctx)
	data := service.BaseSysPermsService().Permmenu(ctx, admin.RoleIds)
	res = dzhcore.Ok(data)
	return
}

// 退出登录
func (c *BaseCommController) Logout(ctx context.Context, req *v1.BaseCommLogoutReq) (res *dzhcore.BaseRes, err error) {

	err = service.BaseSysLoginService().Logout(ctx)
	res = dzhcore.Ok(nil)
	return
}

// 上传模式
func (c *BaseCommController) UploadMode(ctx context.Context, req *v1.BaseCommUploadModeReq) (res *dzhcore.BaseRes, err error) {
	data, err := dzhcore.File().GetMode()
	res = dzhcore.Ok(data)
	return
}

// 上传
func (c *BaseCommController) Upload(ctx context.Context, req *v1.BaseCommUploadReq) (res *dzhcore.BaseRes, err error) {
	data, err := dzhcore.File().Upload(ctx)
	res = dzhcore.Ok(data)
	return
}

// 更新个人信息
func (c *BaseCommController) PersonUpdate(ctx g.Ctx, req *v1.PersonUpdateReq) (res *dzhcore.BaseRes, err error) {

	_, err = service.BaseSysUserService().ServiceUpdate(ctx, &dzhcore.UpdateReq{})
	if err != nil {
		return
	}

	res = dzhcore.Ok(nil)
	return
}
