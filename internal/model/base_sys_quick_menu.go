package model

import (
	"github.com/gzdzh-cn/dzhcore"
)

const TableNameQuickMenu = "base_sys_quick_menu"

// QuickMenu 模型，映射表 <base_sys_quick_menu>
type BaseSysQuickMenu struct {
	*dzhcore.Model
	UserID string  `gorm:"column:user_id;type:varchar(64);not null;default:''" json:"userId"` // 用户ID
	MenuID string  `gorm:"column:menu_id" json:"menuId"`
	Name   string  `gorm:"column:name;type:varchar(255);not null" json:"name"` // 菜单名称
	Router *string `gorm:"column:router;type:varchar(255)" json:"router"`
	Icon   *string `gorm:"column:icon;type:varchar(255)" json:"icon"`

	Status   int     `gorm:"column:status;comment:状态;type:int(11);default:1" json:"status"`            // 状态
	OrderNum int32   `gorm:"column:order_num;comment:排序;type:int;not null;default:99" json:"orderNum"` // 排序
	Remark   *string `gorm:"column:remark;comment:备注;type:varchar(255)" json:"remark"`                 // 备注
}

// TableName QuickMenu 的表名
func (*BaseSysQuickMenu) TableName() string {
	return TableNameQuickMenu
}

// GroupName QuickMenu 的表分组
func (*BaseSysQuickMenu) GroupName() string {
	return "default"
}

// NewBaseSysQuickMenu 创建一个新的 QuickMenu 实例
func NewBaseSysQuickMenu() *BaseSysQuickMenu {
	return &BaseSysQuickMenu{
		Model: dzhcore.NewModel(),
	}
}

// init 注册模型
func init() {
	dzhcore.AddModel(&BaseSysQuickMenu{})
}
