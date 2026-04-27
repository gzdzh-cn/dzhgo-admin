// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BaseSysQuickMenu is the golang structure of table base_sys_quick_menu for DAO operations like Where/Data.
type BaseSysQuickMenu struct {
	g.Meta     `orm:"table:base_sys_quick_menu, do:true"`
	Id         any         //
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
	DeletedAt  *gtime.Time //
	Name       any         //
	Router     any         //
	Icon       any         //
	Status     any         // 状态
	OrderNum   any         // 排序
	Remark     any         // 备注
	MenuId     any         //
	UserId     any         //
}
