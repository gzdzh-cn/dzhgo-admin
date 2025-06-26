package admin

import (
	"context"
	v1 "dzhgo/internal/api/admin_v1"
	logic "dzhgo/internal/logic/sys"
	"dzhgo/internal/service"

	"github.com/gzdzh-cn/dzhcore"
)

type BaseSysAddonsController struct {
	*dzhcore.Controller
}

func init() {
	var baseSysAddonsController = &BaseSysAddonsController{
		&dzhcore.Controller{
			Prefix:  "/admin/base/sys/addons",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: logic.NewsBaseSysAddonsService(),
		},
	}
	// 注册路由
	dzhcore.AddController(baseSysAddonsController)

}

// 安装卸载插件
func (c *BaseSysAddonsController) InstallUpdateStatus(ctx context.Context, req *v1.InstallUpdateStatusReq) (res *dzhcore.BaseRes, err error) {
	data, err := service.BaseSysAddonsService().InstallUpdateStatus(ctx, req)
	if err != nil {
		return
	}
	res = dzhcore.Ok(data)
	return
}

// 上下架插件
func (c *BaseSysAddonsController) LineUpdateStatus(ctx context.Context, req *v1.LineUpdateStatusReq) (res *dzhcore.BaseRes, err error) {
	data, err := service.BaseSysAddonsService().LineUpdateStatus(ctx, req)
	if err != nil {
		return
	}
	res = dzhcore.Ok(data)
	return
}
