package model

import (
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gzdzh-cn/dzhcore"
)

const TableNameBaseSysInit = "base_sys_init"

// BaseSysInit mapped from table <base_sys_init>
type BaseSysInit struct {
	Id     string `gorm:"primaryKey;autoIncrement:false;varchar(255);index" json:"id"`
	Module string `gorm:"index;not null" json:"module"`
	Tables string `gorm:"index;not null" json:"tables"`
	Group  string `gorm:"index;not null" json:"group"`
}

// TableName BaseSysInit's table namer
func (*BaseSysInit) TableName() string {
	return TableNameBaseSysInit
}

// TableGroup BaseSysInit's table group
func (*BaseSysInit) GroupName() string {
	return "default"
}

// GetStruct BaseSysInit's struct
func (m *BaseSysInit) GetStruct() interface{} {
	return m
}

var (
	ctx = gctx.GetInitCtx()
)

// init 创建表
func init() {
	// dzhcore.CreateTable(&BaseSysInit{})
	dzhcore.AddModel(&BaseSysInit{})
}
