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
	Id           any         //
	CreateTime   *gtime.Time // 创建时间
	UpdateTime   *gtime.Time // 更新时间
	DeletedAt    *gtime.Time //
	DepartmentId any         //
	Name         any         //
	Username     any         //
	Password     any         //
	PasswordV    any         //
	NickName     any         //
	HeadImg      any         //
	Phone        any         //
	Email        any         //
	Status       any         //
	Remark       any         //
	SocketId     any         //
	LoginIp      any         //
}
