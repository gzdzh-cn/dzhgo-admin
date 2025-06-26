package admin

import (
	"dzhgo/addons/space/service"

	"github.com/gzdzh-cn/dzhcore"
)

type SpaceTypeController struct {
	*dzhcore.Controller
}

func init() {
	var spaceTypeController = &SpaceTypeController{
		&dzhcore.Controller{
			Prefix:  "/admin/space/type",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewSpaceTypeService(),
		},
	}
	// 注册路由
	dzhcore.AddController(spaceTypeController)
}
