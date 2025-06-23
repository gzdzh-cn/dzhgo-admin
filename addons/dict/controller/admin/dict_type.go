package admin

import (
	logic "dzhgo/addons/dict/logic/sys"

	"github.com/gzdzh-cn/dzhcore"
)

type DictTypeController struct {
	*dzhcore.Controller
}

func init() {
	var dictTypeController = &DictTypeController{
		&dzhcore.Controller{
			Prefix:  "/admin/dict/type",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: logic.NewsDictTypeService(),
		},
	}
	// 注册路由
	dzhcore.AddController(dictTypeController)
}
