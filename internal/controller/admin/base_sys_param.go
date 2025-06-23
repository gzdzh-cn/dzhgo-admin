package admin

import (
	v1 "dzhgo/internal/api/admin_v1"
	logic "dzhgo/internal/logic/sys"
	"dzhgo/internal/service"

	"github.com/gzdzh-cn/dzhcore"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type BaseSysParamController struct {
	*dzhcore.Controller
}

func init() {
	var baseSysParamController = &BaseSysParamController{
		&dzhcore.Controller{
			Prefix:  "/admin/base/sys/param",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: logic.NewsBaseSysParamService(),
		},
	}
	// 注册路由
	dzhcore.AddController(baseSysParamController)
}

// Html 根据配置参数key获取网页内容(富文本)
func (c *BaseSysParamController) Html(ctx g.Ctx, req *v1.BaseSysParamHtmlReq) (res *dzhcore.BaseRes, err error) {

	r := ghttp.RequestFromCtx(ctx)
	r.Response.WriteExit(service.BaseSysParamService().HtmlByKey(req.Key))
	return
}
