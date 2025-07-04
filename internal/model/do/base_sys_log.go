// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BaseSysLog is the golang structure of table base_sys_log for DAO operations like Where/Data.
type BaseSysLog struct {
	g.Meta     `orm:"table:base_sys_log, do:true"`
	Id         interface{} //
	CreateTime *gtime.Time //
	UpdateTime *gtime.Time //
	DeletedAt  *gtime.Time //
	UserId     interface{} //
	Action     interface{} //
	Ip         interface{} //
	IpAddr     interface{} //
	Params     interface{} //
}
