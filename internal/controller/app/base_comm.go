package app

import (
	"github.com/gzdzh-cn/dzhcore"
)

type BaseCommController struct {
	*dzhcore.ControllerSimple
}

func init() {
	var baseCommController = &BaseCommController{
		&dzhcore.ControllerSimple{
			Prefix: "/app/base/comm",
		},
	}
	// 注册路由
	dzhcore.RegisterControllerSimple(baseCommController)
}
