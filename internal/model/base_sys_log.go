package model

import "github.com/gzdzh-cn/dzhcore"

const TableNameBaseSysLog = "base_sys_log"

// BaseSysLog mapped from table <base_sys_log>
type BaseSysLog struct {
	*dzhcore.Model
	UserID string `gorm:"column:userId;index,priority:1" json:"userId"`                                         // 用户ID
	Action string `gorm:"column:action;not null;index:IDX_938f886fb40e163db174b7f6c3,priority:1" json:"action"` // 行为
	IP     string `gorm:"column:ip;index:IDX_24e18767659f8c7142580893f2,priority:1" json:"ip"`                  // ip
	IPAddr string `gorm:"column:ipAddr;index:IDX_a03a27f75cf8d502b3060823e1,priority:1" json:"ipAddr"`          // ip地址
	Params string `gorm:"column:params" json:"params"`                                                          // 参数
}

// TableName BaseSysLog's table name
func (*BaseSysLog) TableName() string {
	return TableNameBaseSysLog
}

// NewBaseSysLog 创建实例
func NewBaseSysLog() *BaseSysLog {
	return &BaseSysLog{
		Model: dzhcore.NewModel(),
	}
}

// init 创建表
func init() {
	// dzhcore.CreateTable(&BaseSysLog{})
	dzhcore.AddModel(&BaseSysLog{})
}
