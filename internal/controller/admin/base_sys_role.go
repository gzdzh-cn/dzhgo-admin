package admin

import (
	logic "dzhgo/internal/logic/sys"

	"github.com/gzdzh-cn/dzhcore"
)

type BaseSysRoleController struct {
	*dzhcore.Controller
}

func init() {
	var baseSysRoleController = &BaseSysRoleController{
		&dzhcore.Controller{
			Prefix:  "/admin/base/sys/role",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: logic.NewsBaseSysRoleService(),
		},
	}
	// 注册路由
	dzhcore.AddController(baseSysRoleController)
}
