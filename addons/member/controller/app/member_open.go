package app

import (
	"context"
	v1 "dzhgo/addons/member/api/app_v1"
	"dzhgo/addons/member/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gzdzh-cn/dzhcore"
)

type MemberOpenController struct {
	*dzhcore.ControllerSimple
}

func init() {
	var memberOpenController = &MemberOpenController{
		&dzhcore.ControllerSimple{
			Prefix: "/app/member/open",
		},
	}
	// 注册路由
	dzhcore.AddControllerSimple(memberOpenController)
}

// 账号登录
func (c *MemberOpenController) AccountLogin(ctx context.Context, req *v1.AccountLoginReq) (res *dzhcore.BaseRes, err error) {
	data, err := service.MemberManageService().AccountLogin(ctx, req)
	if err != nil {
		g.Log().Error(ctx, err.Error())
		return
	}
	res = dzhcore.Ok(data)
	return
}

// 公众号登录
func (c *MemberOpenController) MpLogin(ctx context.Context, req *v1.MpLoginReq) (res *dzhcore.BaseRes, err error) {
	data, err := service.MemberManageService().MpLogin(ctx, req)
	if err != nil {
		return
	}
	res = dzhcore.Ok(data)
	return
}
