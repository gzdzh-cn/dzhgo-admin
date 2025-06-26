// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AddonsMemberManage is the golang structure of table addons_member_manage for DAO operations like Where/Data.
type AddonsMemberManage struct {
	g.Meta        `orm:"table:addons_member_manage, do:true"`
	Id            interface{} //
	CreateTime    *gtime.Time // 创建时间
	UpdateTime    *gtime.Time // 更新时间
	DeletedAt     *gtime.Time //
	AvatarUrl     interface{} // 头像
	Password      interface{} // 会员密码
	PasswordV     interface{} //
	Username      interface{} // 会员账号
	Nickname      interface{} // 会员昵称
	LevelName     interface{} // 等级名称
	Level         interface{} // 等级
	Sex           interface{} // 性别
	Qq            interface{} // qq
	Mobile        interface{} // 手机号
	Wx            interface{} // 微信号
	WxImg         interface{} // 微信二维码
	Email         interface{} // email
	Role          interface{} // 家庭角色
	LastLoginTime *gtime.Time // 最后登录时间
	Openid        interface{} // openid
	UnionId       interface{} // unionId
	SessionKey    interface{} // session_key
	Remark        interface{} // 备注
	Status        interface{} //
	Description   interface{} // 描述
}
