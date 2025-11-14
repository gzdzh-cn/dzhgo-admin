package admin

import (
	"context"
	"dzhgo/internal/common"
	logic "dzhgo/internal/logic/sys"
	"dzhgo/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gzdzh-cn/dzhcore"
)

type BaseSysMenuController struct {
	*dzhcore.Controller
}

func init() {
	var baseSysMenuController = &BaseSysMenuController{
		&dzhcore.Controller{
			Prefix:  "/admin/base/sys/menu",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: logic.NewsBaseSysMenuService(),
		},
	}
	// 注册路由
	dzhcore.AddController(baseSysMenuController)
}

func (c *BaseSysMenuController) Add(ctx context.Context, req *dzhcore.AddReq) (res *dzhcore.BaseRes, err error) {
	res, err = c.Controller.Add(ctx, req)
	if err != nil {
		return
	}
	admin := common.GetAdmin(ctx)
	err = service.BaseSysPermsService().RefreshPerms(ctx, gconv.String(admin.UserId))
	if err != nil {
		return nil, err
	}
	// g.Log().Info(ctx, "刷新权限缓存成功", admin.UserId)
	return
}
