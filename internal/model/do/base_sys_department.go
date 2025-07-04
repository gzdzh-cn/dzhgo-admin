// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BaseSysDepartment is the golang structure of table base_sys_department for DAO operations like Where/Data.
type BaseSysDepartment struct {
	g.Meta     `orm:"table:base_sys_department, do:true"`
	Id         interface{} //
	CreateTime *gtime.Time //
	UpdateTime *gtime.Time //
	DeletedAt  *gtime.Time //
	Name       interface{} //
	ParentId   interface{} //
	OrderNum   interface{} //
}
