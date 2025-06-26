package model

import (
	"github.com/gzdzh-cn/dzhcore"
)

const TableNameBaseSysAddonsTypes = "base_sys_addons_types"

// BaseSysAddonsTypes mapped from table <base_sys_addons_types>
type BaseSysAddonsTypes struct {
	*dzhcore.Model
	Name     string  `gorm:"column:name;not null;comment:标题" json:"name"`
	Image    *string `gorm:"column:image;comment:图片" json:"image"`
	Link     *string `gorm:"column:link;comment:跳转" json:"link"`
	Remark   *string `gorm:"column:remark;comment:备注" json:"remark"`
	Status   string  `gorm:"column:status;comment:状态;type:int;default:1" json:"status"`
	OrderNum int32   `gorm:"column:orderNum;comment:排序;type:int;not null;default:99" json:"orderNum"`
}

// TableName BaseSysAddonsTypes's table name
func (*BaseSysAddonsTypes) TableName() string {
	return TableNameBaseSysAddonsTypes
}

// GroupName BaseSysAddonsTypes's table group
func (*BaseSysAddonsTypes) GroupName() string {
	return "default"
}

// NewBaseSysAddonsTypes create a new BaseSysAddonsTypes
func NewBaseSysAddonsTypes() *BaseSysAddonsTypes {
	return &BaseSysAddonsTypes{
		Model: dzhcore.NewModel(),
	}
}

// init 创建表
func init() {
	// dzhcore.CreateTable(&BaseSysAddonsTypes{})
	dzhcore.AddModel(&BaseSysAddonsTypes{})
}
