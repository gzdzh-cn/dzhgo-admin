// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BaseSysMenu is the golang structure of table base_sys_menu for DAO operations like Where/Data.
type BaseSysMenu struct {
	g.Meta     `orm:"table:base_sys_menu, do:true"`
	Id         interface{} //
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
	DeletedAt  *gtime.Time //
	ParentId   interface{} //
	Name       interface{} //
	Router     interface{} //
	Perms      interface{} //
	Type       interface{} //
	Icon       interface{} //
	OrderNum   interface{} //
	ViewPath   interface{} //
	KeepAlive  interface{} //
	IsShow     interface{} //
	IsInstall  interface{} //
	MenuType   interface{} //
	AddonsName interface{} //
}
