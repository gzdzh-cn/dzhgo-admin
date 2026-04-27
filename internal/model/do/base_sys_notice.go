// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BaseSysNotice is the golang structure of table base_sys_notice for DAO operations like Where/Data.
type BaseSysNotice struct {
	g.Meta     `orm:"table:base_sys_notice, do:true"`
	Id         any         //
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
	DeletedAt  *gtime.Time //
	Title      any         //
	NoType     any         //
	Remark     any         //
}
