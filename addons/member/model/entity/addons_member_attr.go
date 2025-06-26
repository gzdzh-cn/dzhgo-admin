// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AddonsMemberAttr is the golang structure for table addons_member_attr.
type AddonsMemberAttr struct {
	Id             string      `json:"id"             orm:"id"              ` //
	CreateTime     *gtime.Time `json:"createTime"     orm:"createTime"      ` // 创建时间
	UpdateTime     *gtime.Time `json:"updateTime"     orm:"updateTime"      ` // 更新时间
	DeletedAt      *gtime.Time `json:"deletedAt"      orm:"deleted_at"      ` //
	Type           int         `json:"type"           orm:"type"            ` // 类型:1=公众号,2=小程序
	UserId         string      `json:"userId"         orm:"user_id"         ` // 用户ID
	Unionid        string      `json:"unionid"        orm:"unionid"         ` // UnionID
	Notify         int         `json:"notify"         orm:"notify"          ` // 微信通知:1=是,0=否
	Openid         string      `json:"openid"         orm:"openid"          ` // openid
	Nickname       string      `json:"nickname"       orm:"nickname"        ` // 昵称
	Sex            int         `json:"sex"            orm:"sex"             ` // 性别:0=默认,1=男,2=女
	Language       string      `json:"language"       orm:"language"        ` // 语言
	Country        string      `json:"country"        orm:"country"         ` // 国家
	Province       string      `json:"province"       orm:"province"        ` // 省份
	City           string      `json:"city"           orm:"city"            ` // 城市
	Headimgurl     string      `json:"headimgurl"     orm:"headimgurl"      ` // 头像
	Subscribe      int         `json:"subscribe"      orm:"subscribe"       ` // 关注状态:0=未关注,1=已关注
	SubscribeTime  int64       `json:"subscribeTime"  orm:"subscribe_time"  ` // 关注时间
	Groupid        int         `json:"groupid"        orm:"groupid"         ` // 分组ID
	TagidList      string      `json:"tagidList"      orm:"tagid_list"      ` // 标签列表
	Privilege      string      `json:"privilege"      orm:"privilege"       ` // 权限
	SubscribeScene string      `json:"subscribeScene" orm:"subscribe_scene" ` // 关注来源
	QrScene        string      `json:"qrScene"        orm:"qr_scene"        ` // 扫码场景
	QrSceneStr     string      `json:"qrSceneStr"     orm:"qr_scene_str"    ` // 扫码场景描述
}
