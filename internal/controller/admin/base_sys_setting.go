package admin

import (
	logic "dzhgo/internal/logic/sys"

	"github.com/gzdzh-cn/dzhcore"
)

type BaseSysSettingController struct {
	*dzhcore.Controller
}

func init() {
	var baseSysSettingController = &BaseSysSettingController{
		&dzhcore.Controller{
			Prefix:  "/admin/base/sys/setting",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: logic.NewsBaseSysSettingService(),
		},
	}
	// 注册路由
	dzhcore.AddController(baseSysSettingController)
}
