package admin

import (
	"context"
	v1 "dzhgo/internal/api/admin_v1"
	logic "dzhgo/internal/logic/sys"
	"dzhgo/internal/service"

	"github.com/gzdzh-cn/dzhcore"
)

type BaseSysDepartmentController struct {
	*dzhcore.Controller
}

func init() {
	var baseSysDepartmentController = &BaseSysDepartmentController{
		&dzhcore.Controller{
			Prefix:  "/admin/base/sys/department",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: logic.NewsBaseSysDepartmentService(),
		},
	}
	// 注册路由
	dzhcore.AddController(baseSysDepartmentController)
}

// Order 排序部门
func (c *BaseSysDepartmentController) Order(ctx context.Context, req *v1.OrderReq) (res *dzhcore.BaseRes, err error) {
	err = service.BaseSysDepartmentService().Order(ctx)
	res = dzhcore.Ok(nil)
	return
}
