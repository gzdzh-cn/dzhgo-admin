// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AddonsMemberAttr is the golang structure of table addons_member_attr for DAO operations like Where/Data.
type AddonsMemberAttr struct {
	g.Meta         `orm:"table:addons_member_attr, do:true"`
	Id             interface{} //
	CreateTime     *gtime.Time // 创建时间
	UpdateTime     *gtime.Time // 更新时间
	DeletedAt      *gtime.Time //
	Type           interface{} // 类型:1=公众号,2=小程序
	UserId         interface{} // 用户ID
	Unionid        interface{} // UnionID
	Notify         interface{} // 微信通知:1=是,0=否
	Openid         interface{} // openid
	Nickname       interface{} // 昵称
	Sex            interface{} // 性别:0=默认,1=男,2=女
	Language       interface{} // 语言
	Country        interface{} // 国家
	Province       interface{} // 省份
	City           interface{} // 城市
	Headimgurl     interface{} // 头像
	Subscribe      interface{} // 关注状态:0=未关注,1=已关注
	SubscribeTime  interface{} // 关注时间
	Groupid        interface{} // 分组ID
	TagidList      interface{} // 标签列表
	Privilege      interface{} // 权限
	SubscribeScene interface{} // 关注来源
	QrScene        interface{} // 扫码场景
	QrSceneStr     interface{} // 扫码场景描述
}
