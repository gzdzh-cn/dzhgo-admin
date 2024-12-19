package app

import (
	"github.com/gzdzh-cn/dzhcore"
)

type BaseMemberController struct {
	*dzhcore.Controller
}

//func init() {
//	var baseMemberController = &BaseMemberController{
//		&dzhcore.Controller{
//			Prefix:  "/app/base/member",
//			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
//			Service: logic.NewsBaseSysMemberService(),
//		},
//	}
//	// 注册路由
//	dzhcore.RegisterController(baseMemberController)
//}
