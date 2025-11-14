// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BaseSysUser is the golang structure of table base_sys_user for DAO operations like Where/Data.
type BaseSysUser struct {
	g.Meta       `orm:"table:base_sys_user, do:true"`
	Id           interface{} //
	CreateTime   *gtime.Time // 创建时间
	UpdateTime   *gtime.Time // 更新时间
	DeletedAt    *gtime.Time //
	DepartmentId interface{} //
	Name         interface{} //
	Username     interface{} //
	Password     interface{} //
	PasswordV    interface{} //
	NickName     interface{} //
	HeadImg      interface{} //
	Phone        interface{} //
	Email        interface{} //
	Status       interface{} //
	Remark       interface{} //
	SocketId     interface{} //
	LoginIp      interface{} //
}
