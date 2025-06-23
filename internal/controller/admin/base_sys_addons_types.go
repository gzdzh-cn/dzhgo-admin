package admin

import (
	logic "dzhgo/internal/logic/sys"

	"github.com/gzdzh-cn/dzhcore"
)

type BaseSysAddonsTypesController struct {
	*dzhcore.Controller
}

func init() {
	var baseSysAddonsTypesController = &BaseSysAddonsTypesController{
		&dzhcore.Controller{
			Prefix:  "/admin/base/sys/addonsTypes",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: logic.NewsBaseSysAddonsTypesService(),
		},
	}
	// 注册路由
	dzhcore.AddController(baseSysAddonsTypesController)
}
