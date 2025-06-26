package admin

import (
	logic "dzhgo/addons/member/logic/sys"

	"github.com/gzdzh-cn/dzhcore"
)

type MemberManageController struct {
	*dzhcore.Controller
}

func init() {
	var memberManageController = &MemberManageController{
		&dzhcore.Controller{
			Prefix:  "/admin/member/manage",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: logic.NewsMemberManageService(),
		},
	}
	// 注册路由
	dzhcore.AddController(memberManageController)
}
