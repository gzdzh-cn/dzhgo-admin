// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BaseSysRoleDepartment is the golang structure for table base_sys_role_department.
type BaseSysRoleDepartment struct {
	Id           string      `json:"id"           orm:"id"           ` //
	CreateTime   *gtime.Time `json:"createTime"   orm:"createTime"   ` // 创建时间
	UpdateTime   *gtime.Time `json:"updateTime"   orm:"updateTime"   ` // 更新时间
	DeletedAt    *gtime.Time `json:"deletedAt"    orm:"deleted_at"   ` //
	RoleId       string      `json:"roleId"       orm:"roleId"       ` //
	DepartmentId string      `json:"departmentId" orm:"departmentId" ` //
}
