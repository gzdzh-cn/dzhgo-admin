package common

import (
	"context"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
)

type Member struct {
	IsRefresh       bool   `json:"isRefresh"`
	MemberName      string `json:"memberName"`
	NickName        string `json:"nickName"`
	LevelName       string `json:"levelName"`
	Level           int    `json:"level"`
	MemberId        string `json:"memberId"`
	PasswordVersion *int32 `json:"passwordVersion"`
}

// GetMember 获取传入ctx 中的 member 对象
func GetMember(ctx context.Context) *Member {
	r := g.RequestFromCtx(ctx)
	var member *Member
	err := gjson.New(r.GetCtxVar("member").String()).Scan(&member)
	if err != nil {
		g.Log().Error(ctx, err)
	}
	return member
}
