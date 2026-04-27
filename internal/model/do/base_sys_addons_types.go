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
	Id         any         //
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
	DeletedAt  *gtime.Time //
	Name       any         // 标题
	Image      any         // 图片
	Link       any         // 跳转
	Remark     any         // 备注
	Status     any         // 状态
	OrderNum   any         // 排序
}
