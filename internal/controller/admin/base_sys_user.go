package admin

import (
	"context"
	v1 "dzhgo/internal/api/admin_v1"
	logic "dzhgo/internal/logic/sys"
	"dzhgo/internal/service"

	"github.com/gzdzh-cn/dzhcore"
)

type BaseSysUserController struct {
	*dzhcore.Controller
}

func init() {
	var baseSysUserController = &BaseSysUserController{
		&dzhcore.Controller{
			Prefix:  "/admin/base/sys/user",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page", "Move"},
			Service: logic.NewsBaseSysUserService(),
		},
	}
	// 注册路由
	dzhcore.AddController(baseSysUserController)
}

// 移动
func (c *BaseSysUserController) Move(ctx context.Context, req *v1.UserMoveReq) (res *dzhcore.BaseRes, err error) {
	err = service.BaseSysUserService().Move(ctx)
	res = dzhcore.Ok(nil)
	return
}
