// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BaseSysAddonsTypes is the golang structure of table base_sys_addons_types for DAO operations like Where/Data.
type BaseSysAddonsTypes struct {
	g.Meta     `orm:"table:base_sys_addons_types, do:true"`
	Id         interface{} //
	CreateTime *gtime.Time //
	UpdateTime *gtime.Time //
	DeletedAt  *gtime.Time //
	Name       interface{} //
	Image      interface{} //
	Link       interface{} //
	Remark     interface{} //
	Status     interface{} //
	OrderNum   interface{} //
}
