package model

import (
	"github.com/gzdzh-cn/dzhcore"
)

const TableNameBaseSysActionLog = "base_sys_action_log"

// Base_sys_action_log 模型，映射表 <addons_base_sys_action_log>
type BaseSysActionLog struct {
	*dzhcore.Model
	UserID string  `gorm:"column:user_id;index,priority:1" json:"userId"`                 // 用户ID
	Name   string  `gorm:"column:name;type:varchar(255);not null" json:"name"`            // 名称
	Status int     `gorm:"column:status;comment:状态;type:int(11);default:1" json:"status"` // 状态
	Remark *string `gorm:"column:remark;type:varchar(255)" json:"remark"`                 // 备注
}

// TableName Base_sys_action_log 的表名
func (*BaseSysActionLog) TableName() string {
	return TableNameBaseSysActionLog
}

// GroupName Base_sys_action_log 的表分组
func (*BaseSysActionLog) GroupName() string {
	return "default"
}

// NewBase_sys_action_log 创建一个新的 Base_sys_action_log 实例
func NewBaseSysActionLog() *BaseSysActionLog {
	return &BaseSysActionLog{
		Model: dzhcore.NewModel(),
	}
}

// init 注册模型
func init() {
	dzhcore.AddModel(&BaseSysActionLog{})
}
