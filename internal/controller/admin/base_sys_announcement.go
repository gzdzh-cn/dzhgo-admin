package admin

import (
	"context"
	v1 "dzhgo/internal/api/admin_v1"
	logic "dzhgo/internal/logic/sys"
	"dzhgo/internal/service"

	"github.com/gzdzh-cn/dzhcore"
)

type BaseSysAnnouncementController struct {
	*dzhcore.Controller
}

func init() {
	var baseSysAnnouncementController = &BaseSysAnnouncementController{
		&dzhcore.Controller{
			Prefix:  "/admin/base/sys/announcement",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: logic.NewsBaseSysAnnouncementService(),
		},
	}

	// 注册路由
	dzhcore.AddController(baseSysAnnouncementController)
}

// AnnouncementMarkRead 标记公告为已读
func (c *BaseSysAnnouncementController) AnnouncementMarkRead(ctx context.Context, req *v1.AnnouncementMarkReadReq) (res *dzhcore.BaseRes, err error) {
	err = service.BaseSysAnnouncementService().MarkRead(ctx, req.Id)
	if err != nil {
		return
	}
	res = dzhcore.Ok(nil)
	return
}

// AnnouncementMarkAllRead 一键已阅所有公告
func (c *BaseSysAnnouncementController) AnnouncementMarkAllRead(ctx context.Context, req *v1.AnnouncementMarkAllReadReq) (res *dzhcore.BaseRes, err error) {
	err = service.BaseSysAnnouncementService().MarkAllRead(ctx)
	if err != nil {
		return
	}
	res = dzhcore.Ok(nil)
	return
}
