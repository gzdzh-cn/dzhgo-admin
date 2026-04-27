package admin

import (
	"context"
	v1 "dzhgo/internal/api/admin_v1"
	"dzhgo/internal/common"
	logic "dzhgo/internal/logic/sys"
	"dzhgo/internal/service"

	"github.com/gzdzh-cn/dzhcore"
)

type BaseSysQuickMenuController struct {
	*dzhcore.Controller
}

func init() {
	var baseSysQuickMenuController = &BaseSysQuickMenuController{
		&dzhcore.Controller{
			Prefix:  "/admin/base/sys/quickMenu",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: logic.NewsBaseSysQuickMenuService(),
		},
	}

	// 注册路由
	dzhcore.AddController(baseSysQuickMenuController)
}

// 获取当前用户的快捷菜单列表
func (c *BaseSysQuickMenuController) QuickMenuList(ctx context.Context, req *v1.QuickMenuListReq) (res *dzhcore.BaseRes, err error) {
	data, err := service.BaseSysQuickMenuService().QuickMenuList(ctx, common.GetAdmin(ctx).RoleIds)
	if err != nil {
		return
	}
	res = dzhcore.Ok(data)
	return
}
