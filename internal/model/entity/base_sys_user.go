// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BaseSysUser is the golang structure for table base_sys_user.
type BaseSysUser struct {
	Id           string      `json:"id"           orm:"id"           ` //
	CreateTime   *gtime.Time `json:"createTime"   orm:"createTime"   ` // 创建时间
	UpdateTime   *gtime.Time `json:"updateTime"   orm:"updateTime"   ` // 更新时间
	DeletedAt    *gtime.Time `json:"deletedAt"    orm:"deleted_at"   ` //
	DepartmentId string      `json:"departmentId" orm:"departmentId" ` //
	Name         string      `json:"name"         orm:"name"         ` //
	Username     string      `json:"username"     orm:"username"     ` //
	Password     string      `json:"password"     orm:"password"     ` //
	PasswordV    int         `json:"passwordV"    orm:"passwordV"    ` //
	NickName     string      `json:"nickName"     orm:"nickName"     ` //
	HeadImg      string      `json:"headImg"      orm:"headImg"      ` //
	Phone        string      `json:"phone"        orm:"phone"        ` //
	Email        string      `json:"email"        orm:"email"        ` //
	Status       int         `json:"status"       orm:"status"       ` //
	Remark       string      `json:"remark"       orm:"remark"       ` //
	SocketId     string      `json:"socketId"     orm:"socketId"     ` //
	LoginIp      string      `json:"loginIp"      orm:"loginIp"      ` //
}
