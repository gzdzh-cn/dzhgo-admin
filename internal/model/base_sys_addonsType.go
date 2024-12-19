package model

import (
	"github.com/gzdzh-cn/dzhcore"
)

const TableNameBaseSysAddonsType = "base_sys_addonsType"

// BaseSysAddonsType mapped from table <member_open>
type BaseSysAddonsType struct {
	*dzhcore.Model
	Name     string  `gorm:"column:name;not null;comment:标题" json:"name"`
	Image    *string `gorm:"column:image;comment:图片" json:"image"`
	Link     *string `gorm:"column:link;comment:跳转" json:"link"`
	Remark   *string `gorm:"column:remark;comment:备注" json:"remark"`
	Status   string  `gorm:"column:status;comment:状态;type:int;default:1" json:"status"`
	OrderNum int32   `gorm:"column:orderNum;comment:排序;type:int;not null;default:99" json:"orderNum"`
}

// TableName BaseSysAddonsType's table name
func (*BaseSysAddonsType) TableName() string {
	return TableNameBaseSysAddonsType
}

// GroupName BaseSysAddonsType's table group
func (*BaseSysAddonsType) GroupName() string {
	return "default"
}

// NewBaseSysAddonsType create a new BaseSysAddonsType
func NewBaseSysAddonsType() *BaseSysAddonsType {
	return &BaseSysAddonsType{
		Model: dzhcore.NewModel(),
	}
}

// init 创建表
func init() {
	err := dzhcore.CreateTable(&BaseSysAddonsType{})
	if err != nil {
		return
	}
}
