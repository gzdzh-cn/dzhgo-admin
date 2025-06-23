package model

import "github.com/gzdzh-cn/dzhcore"

const TableNameBaseSysRoleMenu = "base_sys_role_menu"

// BaseSysRoleMenu mapped from table <base_sys_role_menu>
type BaseSysRoleMenu struct {
	*dzhcore.Model
	RoleID string `gorm:"column:roleId;not null" json:"roleId"` // 角色ID
	MenuID string `gorm:"column:menuId;not null" json:"menuId"` // 菜单ID
}

// TableName BaseSysRoleMenu's table name
func (*BaseSysRoleMenu) TableName() string {
	return TableNameBaseSysRoleMenu
}

// NewBaseSysRoleMenu create a new BaseSysRoleMenu
func NewBaseSysRoleMenu() *BaseSysRoleMenu {
	return &BaseSysRoleMenu{
		Model: &dzhcore.Model{},
	}
}

// init 创建表
func init() {
	// dzhcore.CreateTable(&BaseSysRoleMenu{})
	dzhcore.AddModel(&BaseSysRoleMenu{})
}
