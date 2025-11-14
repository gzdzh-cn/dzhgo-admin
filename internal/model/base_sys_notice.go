package model

import (
	"github.com/gzdzh-cn/dzhcore"
)

const TableNameBaseSysNotice = "base_sys_notice"

// BaseSysNotice 模型，映射表 <base_sys_notice>
type BaseSysNotice struct {
	*dzhcore.Model
	Title  string  `gorm:"column:title;type:varchar(255);not null" json:"title"`        // 标题
	NoType string  `gorm:"column:noType;type:varchar(100);default:info" json:"no_type"` // 通知类型
	Remark *string `gorm:"column:remark;type:varchar(255)" json:"remark"`               // 备注
}

// TableName BaseSysNotice 的表名
func (*BaseSysNotice) TableName() string {
	return TableNameBaseSysNotice
}

// GroupName BaseSysNotice 的表分组
func (*BaseSysNotice) GroupName() string {
	return "default"
}

// NewBaseSysNotice 创建一个新的 BaseSysNotice 实例
func NewBaseSysNotice() *BaseSysNotice {
	return &BaseSysNotice{
		Model: dzhcore.NewModel(),
	}
}

// init 注册模型
func init() {
	dzhcore.AddModel(&BaseSysNotice{})
}
