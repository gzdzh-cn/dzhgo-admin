package admin

import (
	logic "dzhgo/internal/logic/sys"
	"github.com/gzdzh-cn/dzhcore"
)

type BaseSysAddonsTypeController struct {
	*dzhcore.Controller
}

func init() {
	var baseSysAddonsTypeController = &BaseSysAddonsTypeController{
		&dzhcore.Controller{
			Prefix:  "/admin/base/sys/addonsType",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: logic.NewsBaseSysAddonsTypeService(),
		},
	}
	// 注册路由
	dzhcore.RegisterController(baseSysAddonsTypeController)
}
