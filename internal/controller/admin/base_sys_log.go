package admin

import (
	v1 "dzhgo/internal/api/admin_v1"
	logic "dzhgo/internal/logic/sys"
	"dzhgo/internal/service"

	"github.com/gzdzh-cn/dzhcore"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type BaseSysLogController struct {
	*dzhcore.Controller
}

func init() {
	var baseSysLogController = &BaseSysLogController{
		&dzhcore.Controller{
			Prefix:  "/admin/base/sys/log",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: logic.NewsBaseSysLogService(),
		},
	}
	// 注册路由
	dzhcore.AddController(baseSysLogController)
}

// 设置保留天数
func (c *BaseSysLogController) SetKeep(ctx g.Ctx, req *v1.SetKeepReq) (res *dzhcore.BaseRes, err error) {
	err = service.BaseSysConfService().UpdateValue("logKeep", gconv.String(req.Value))
	return
}

// 获取保留天数
func (c *BaseSysLogController) GetKeep(ctx g.Ctx, req *v1.GetKeepReq) (res *dzhcore.BaseRes, err error) {
	res = dzhcore.Ok(service.BaseSysConfService().GetValue("logKeep"))
	return
}

// 清空日志
func (c *BaseSysLogController) Clear(ctx g.Ctx, req *v1.ClearReq) (res *dzhcore.BaseRes, err error) {
	err = service.BaseSysLogService().Clear(true)
	return
}
