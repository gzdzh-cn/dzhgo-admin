package admin

import (
	"context"
	"dzhgo/internal/api/admin_v1"
	logic "dzhgo/internal/logic/sys"
	"dzhgo/internal/service"

	"github.com/gzdzh-cn/dzhcore"
)

type BaseSysNoticeController struct {
	*dzhcore.Controller
}

func init() {
	var baseSysNoticeController = &BaseSysNoticeController{
		&dzhcore.Controller{
			Prefix:  "/admin/base/sys/notice",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: logic.NewBaseSysNoticeService(),
		},
	}

	// 注册路由
	dzhcore.AddController(baseSysNoticeController)
}

// 一键已阅
func (c *BaseSysNoticeController) ReadAll(ctx context.Context, req *admin_v1.NoticeReadAllReq) (res *dzhcore.BaseRes, err error) {

	data, err := service.BaseSysNoticeService().ServiceReadAll(ctx)
	if err != nil {
		return
	}

	return dzhcore.Ok(data), nil
}
