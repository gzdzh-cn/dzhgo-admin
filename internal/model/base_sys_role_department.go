package model

import "github.com/gzdzh-cn/dzhcore"

const TableNameBaseSysRoleDepartment = "base_sys_role_department"

// BaseSysRoleDepartment mapped from table <base_sys_role_department>
type BaseSysRoleDepartment struct {
	*dzhcore.Model
	RoleID       string `gorm:"column:roleId;not null" json:"roleId"`             // 角色ID
	DepartmentID string `gorm:"column:departmentId;not null" json:"departmentId"` // 部门ID
}

// TableName BaseSysRoleDepartment's table name
func (*BaseSysRoleDepartment) TableName() string {
	return TableNameBaseSysRoleDepartment
}

// NewBaseSysRoleDepartment create a new BaseSysRoleDepartment
func NewBaseSysRoleDepartment() *BaseSysRoleDepartment {
	return &BaseSysRoleDepartment{
		Model: dzhcore.NewModel(),
	}
}

// init 创建表
func init() {
	// dzhcore.CreateTable(&BaseSysRoleDepartment{})
	dzhcore.AddModel(&BaseSysRoleDepartment{})
}
