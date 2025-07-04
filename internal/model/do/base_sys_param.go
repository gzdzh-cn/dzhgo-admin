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
	Id         interface{} //
	CreateTime *gtime.Time //
	UpdateTime *gtime.Time //
	DeletedAt  *gtime.Time //
	KeyName    interface{} //
	Name       interface{} //
	Data       interface{} //
	DataType   interface{} //
	Remark     interface{} //
}
