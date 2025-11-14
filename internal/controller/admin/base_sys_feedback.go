package admin

import (
	logic "dzhgo/internal/logic/sys"

	"github.com/gzdzh-cn/dzhcore"
)

type BaseSysFeedbackController struct {
	*dzhcore.Controller
}

func init() {
	var baseSysFeedbackController = &BaseSysFeedbackController{
		&dzhcore.Controller{
			Prefix:  "/admin/base/sys/feedback",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: logic.NewBaseSysFeedbackService(),
		},
	}

	// 注册路由
	dzhcore.AddController(baseSysFeedbackController)
}
