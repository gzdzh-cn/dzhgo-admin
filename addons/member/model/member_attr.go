package model

import (
	"github.com/gzdzh-cn/dzhcore"
)

const TableNameMemberAttr = "addons_member_attr"

// MemberAttr mapped from table <addons_member_attr>
type MemberAttr struct {
	*dzhcore.Model
	Type           int    `gorm:"column:type;comment:类型:1=公众号,2=小程序;type:int(11);;not null;default:1" json:"type"`
	UserID         string `gorm:"column:user_id;comment:用户ID;not null" json:"user_id"`
	UnionID        string `gorm:"column:unionid;comment:UnionID" json:"unionid"`
	Notify         int    `gorm:"column:notify;comment:微信通知:1=是,0=否;not null;type:int(11);default:0" json:"notify"`
	OpenID         string `gorm:"column:openid;comment:openid;default:''" json:"openid"`
	Nickname       string `gorm:"column:nickname;comment:昵称" json:"nickname"`
	Sex            int    `gorm:"column:sex;comment:性别:0=默认,1=男,2=女;not null;type:int(11);default:0" json:"sex"`
	Language       string `gorm:"column:language;comment:语言" json:"language"`
	Country        string `gorm:"column:country;comment:国家" json:"country"`
	Province       string `gorm:"column:province;comment:省份" json:"province"`
	City           string `gorm:"column:city;comment:城市" json:"city"`
	HeadImgURL     string `gorm:"column:headimgurl;comment:头像" json:"headimgurl"`
	Subscribe      int    `gorm:"column:subscribe;comment:关注状态:0=未关注,1=已关注;type:int(11);not null;default:0" json:"subscribe"`
	SubscribeTime  int    `gorm:"column:subscribe_time;comment:关注时间" json:"subscribe_time"`
	GroupID        int    `gorm:"column:groupid;comment:分组ID;type:int(11);" json:"groupid"`
	TagIDList      string `gorm:"column:tagid_list;comment:标签列表" json:"tagid_list"`
	Privilege      string `gorm:"column:privilege;comment:权限" json:"privilege"`
	SubscribeScene string `gorm:"column:subscribe_scene;comment:关注来源" json:"subscribe_scene"`
	QRScene        string `gorm:"column:qr_scene;comment:扫码场景" json:"qr_scene"`
	QRSceneStr     string `gorm:"column:qr_scene_str;comment:扫码场景描述" json:"qr_scene_str"`
}

// TableName MemberAttr's table name
func (*MemberAttr) TableName() string {
	return TableNameMemberAttr
}

// GroupName MemberAttr's table group
func (*MemberAttr) GroupName() string {
	return "default"
}

// NewMemberAttr create a new MemberAttr
func NewMemberAttr() *MemberAttr {
	return &MemberAttr{
		Model: dzhcore.NewModel(),
	}
}

// init 创建表
func init() {
	dzhcore.AddModel(&MemberAttr{})
}
