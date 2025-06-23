package model

import "github.com/gzdzh-cn/dzhcore"

const TableNameBaseSysUserRole = "base_sys_user_role"

// BaseSysUserRole mapped from table <base_sys_user_role>
type BaseSysUserRole struct {
	*dzhcore.Model
	UserID string `gorm:"column:userId;not null" json:"userId"` // 用户ID
	RoleID string `gorm:"column:roleId;not null" json:"roleId"` // 角色ID
}

// TableName BaseSysUserRole's table name
func (*BaseSysUserRole) TableName() string {
	return TableNameBaseSysUserRole
}

// NewBaseSysUserRole create a new BaseSysUserRole
func NewBaseSysUserRole() *BaseSysUserRole {
	return &BaseSysUserRole{
		Model: dzhcore.NewModel(),
	}
}

// init 创建表
func init() {
	// dzhcore.CreateTable(&BaseSysUserRole{})
	dzhcore.AddModel(&BaseSysUserRole{})
}
