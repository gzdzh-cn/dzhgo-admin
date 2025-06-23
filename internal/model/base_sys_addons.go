package model

import (
	"github.com/gzdzh-cn/dzhcore"
)

const TableNameBaseSysAddons = "base_sys_addons"

// BaseSysAddons mapped from table <base_sys_addons>
type BaseSysAddons struct {
	*dzhcore.Model
	Name     string  `gorm:"column:name;not null;comment:标题" json:"name"`
	Image    *string `gorm:"column:image;comment:图片" json:"image"`
	Link     *string `gorm:"column:link;comment:跳转" json:"link"`
	MenuId   string  `gorm:"column:menuId;comment:菜单;index" json:"menuId"`
	TypeId   string  `gorm:"column:typeId;comment:类别;index" json:"typeId"`
	Remark   *string `gorm:"column:remark;comment:备注" json:"remark"`
	Status   string  `gorm:"column:status;comment:状态;type:int;default:1" json:"status"`
	OrderNum int32   `gorm:"column:orderNum;comment:排序;type:int;not null;default:99" json:"orderNum"`
}

// TableName BaseSysAddons's table name
func (*BaseSysAddons) TableName() string {
	return TableNameBaseSysAddons
}

// GroupName BaseSysAddons's table group
func (*BaseSysAddons) GroupName() string {
	return "default"
}

// NewBaseSysAddons create a new BaseSysAddons
func NewBaseSysAddons() *BaseSysAddons {
	return &BaseSysAddons{
		Model: dzhcore.NewModel(),
	}
}

// init 创建表
func init() {
	// dzhcore.CreateTable(&BaseSysAddons{})
	dzhcore.AddModel(&BaseSysAddons{})

}
