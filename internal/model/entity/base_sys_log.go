// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BaseSysLog is the golang structure for table base_sys_log.
type BaseSysLog struct {
	Id         string      `json:"id"         orm:"id"         ` //
	CreateTime *gtime.Time `json:"createTime" orm:"createTime" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" orm:"updateTime" ` // 更新时间
	DeletedAt  *gtime.Time `json:"deletedAt"  orm:"deleted_at" ` //
	UserId     string      `json:"userId"     orm:"userId"     ` //
	Action     string      `json:"action"     orm:"action"     ` //
	Ip         string      `json:"ip"         orm:"ip"         ` //
	IpAddr     string      `json:"ipAddr"     orm:"ipAddr"     ` //
	Params     string      `json:"params"     orm:"params"     ` //
}
