// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BaseSysParam is the golang structure of table base_sys_param for DAO operations like Where/Data.
type BaseSysParam struct {
	g.Meta     `orm:"table:base_sys_param, do:true"`
	Id         any         //
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
	DeletedAt  *gtime.Time //
	KeyName    any         //
	Name       any         //
	Data       any         //
	DataType   any         //
	Remark     any         //
}
