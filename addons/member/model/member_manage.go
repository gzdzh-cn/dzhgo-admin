package model

import (
	"time"

	"github.com/gzdzh-cn/dzhcore"
)

const TableNameMemberManage = "addons_member_manage"

// MemberManage mapped from table <addons_member_manage>
type MemberManage struct {
	*dzhcore.Model
	AvatarUrl     string    `gorm:"column:avatarUrl;comment:头像;type:varchar(200)" json:"avatarUrl"`
	Password      string    `gorm:"column:password;not null;comment:会员密码;type:varchar(50)" json:"password"`
	PasswordV     *int32    `gorm:"column:passwordV;not null;type:int(11);default:1" json:"passwordV"` // 密码版本, 作用是改完密码，让原来的token失效
	UserName      string    `gorm:"column:username;comment:会员账号;type:varchar(50);index" json:"username"`
	Nickname      string    `gorm:"column:nickname;comment:会员昵称;type:varchar(50);index" json:"nickname"`
	LevelName     string    `gorm:"column:levelName;comment:等级名称;type:varchar(50);default:普通会员" json:"levelName"`
	Level         int       `gorm:"column:level;comment:等级;type:int(11);default:1" json:"level"`
	Sex           int       `gorm:"column:sex;comment:性别;type:int(11);default:1" json:"sex"`
	QQ            *string   `gorm:"column:qq;comment:qq;type:varchar(255);index" json:"qq"`
	Mobile        *string   `gorm:"column:mobile;comment:手机号;type:varchar(50);index" json:"mobile"`
	Wx            *string   `gorm:"column:wx;comment:微信号;type:varchar(50);index" json:"wx"`
	WxImg         *string   `gorm:"column:wxImg;comment:微信二维码;type:varchar(255)" json:"wxImg"`
	Email         *string   `gorm:"column:email;comment:email;type:varchar(50);index" json:"email"`
	Role          string    `gorm:"column:role;comment:家庭角色;" json:"role"`
	LastLoginTime time.Time `gorm:"column:lastLoginTime;comment:最后登录时间;" json:"lastLoginTime"`

	Openid     string `gorm:"column:openid;comment:openid;" json:"openid"`
	UnionId    string `gorm:"column:unionId;comment:unionId;" json:"unionId"`
	SessionKey string `gorm:"column:session_key;comment:session_key;" json:"session_key"`

	Remark      *string `gorm:"column:remark;comment:备注;type:varchar(255)" json:"remark"`
	Status      int     `gorm:"column:status;not null;type:int;default:1" json:"status"` // 状态 0:禁用 1：启用
	Description *string `gorm:"column:description;comment:描述;type:varchar(100);default:写一个霸气侧漏的签名" json:"description"`
}

// TableName MemberManage's table name
func (*MemberManage) TableName() string {
	return TableNameMemberManage
}

// GroupName MemberManage's table group
func (*MemberManage) GroupName() string {
	return "default"
}

// NewMemberManage create a new MemberManage
func NewMemberManage() *MemberManage {
	return &MemberManage{
		Model: dzhcore.NewModel(),
	}
}

// init 创建表
func init() {
	dzhcore.AddModel(&MemberManage{})
}
