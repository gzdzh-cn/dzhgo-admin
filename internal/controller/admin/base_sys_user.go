package admin

import (
	"context"
	v1 "dzhgo/internal/api/admin_v1"
	logic "dzhgo/internal/logic/sys"
	"dzhgo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gzdzh-cn/dzhcore"
)

type BaseSysUserController struct {
	*dzhcore.Controller
}

func init() {
	var baseSysUserController = &BaseSysUserController{
		&dzhcore.Controller{
			Prefix:  "/admin/base/sys/user",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page", "Move", "Count", "OnlineCount"},
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

// 获取用户总数
func (c *BaseSysUserController) Count(ctx context.Context, req *v1.UserCountReq) (res *dzhcore.BaseRes, err error) {
	count, err := service.BaseSysUserService().Count(ctx)
	if err != nil {
		return
	}
	res = dzhcore.Ok(g.Map{"count": count})
	return
}

// 获取在线用户数
func (c *BaseSysUserController) OnlineCount(ctx context.Context, req *v1.UserOnlineCountReq) (res *dzhcore.BaseRes, err error) {
	count, err := service.BaseSysUserService().OnlineCount(ctx)
	if err != nil {
		return
	}
	res = dzhcore.Ok(g.Map{"count": count})
	return
}
