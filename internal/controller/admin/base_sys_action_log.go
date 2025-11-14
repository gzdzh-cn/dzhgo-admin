package admin

import (
	logic "dzhgo/internal/logic/sys"

	"github.com/gzdzh-cn/dzhcore"
)

type BaseSysActionLogController struct {
	*dzhcore.Controller
}

func init() {
	var baseSysActionLogController = &BaseSysActionLogController{
		&dzhcore.Controller{
			Prefix:  "/admin/base/sys/actionLog",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: logic.NewBaseSysActionLogService(),
		},
	}

	// 注册路由
	dzhcore.AddController(baseSysActionLogController)
}
