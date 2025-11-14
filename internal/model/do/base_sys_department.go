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
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
	DeletedAt  *gtime.Time //
	Name       interface{} //
	ParentId   interface{} //
	OrderNum   interface{} //
}
