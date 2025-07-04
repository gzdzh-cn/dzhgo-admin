// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BaseSysUserRole is the golang structure of table base_sys_user_role for DAO operations like Where/Data.
type BaseSysUserRole struct {
	g.Meta     `orm:"table:base_sys_user_role, do:true"`
	Id         interface{} //
	CreateTime *gtime.Time //
	UpdateTime *gtime.Time //
	DeletedAt  *gtime.Time //
	UserId     interface{} //
	RoleId     interface{} //
}
