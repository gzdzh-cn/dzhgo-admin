package common

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
)

type Member struct {
	IsRefresh       bool   `json:"isRefresh"`
	RoleIds         []uint `json:"roleIds"`
	Username        string `json:"username"`
	NickName        string `json:"nickName"`
	LevelName       string `json:"levelName"`
	UserId          string `json:"userId"`
	PasswordVersion *int32 `json:"passwordVersion"`
}

// GetMember 获取传入ctx 中的 member 对象
func GetMember(ctx context.Context) *Member {
	r := g.RequestFromCtx(ctx)
	member := &Member{}
	err := gjson.New(r.GetCtxVar("admin").String()).Scan(member)
	if err != nil {
		g.Log().Error(ctx, err)
	}
	return member
}
