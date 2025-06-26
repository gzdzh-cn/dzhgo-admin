package app

import (
	"context"
	v1 "dzhgo/addons/member/api/app_v1"
	"dzhgo/addons/member/common"
	"dzhgo/addons/member/service"

	"github.com/gzdzh-cn/dzhcore"
)

type MemberCommController struct {
	*dzhcore.ControllerSimple
}

func init() {
	var memberCommController = &MemberCommController{
		&dzhcore.ControllerSimple{
			Prefix: "/app/member/comm",
		},
	}
	// 注册路由
	dzhcore.AddControllerSimple(memberCommController)
}

// Person 方法 返回不带密码的用户信息
func (c *MemberCommController) Person(ctx context.Context, req *v1.PersonReq) (res *dzhcore.BaseRes, err error) {

	member := common.GetMember(ctx)
	data, err := service.MemberManageService().Person(ctx, member.UserId)
	if err != nil {
		return
	}
	res = dzhcore.Ok(data)
	return
}
