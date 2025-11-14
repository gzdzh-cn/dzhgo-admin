// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BaseSysRole is the golang structure for table base_sys_role.
type BaseSysRole struct {
	Id         string      `json:"id"         orm:"id"         ` //
	CreateTime *gtime.Time `json:"createTime" orm:"createTime" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" orm:"updateTime" ` // 更新时间
	DeletedAt  *gtime.Time `json:"deletedAt"  orm:"deleted_at" ` //
	UserId     string      `json:"userId"     orm:"userId"     ` //
	Name       string      `json:"name"       orm:"name"       ` //
	Label      string      `json:"label"      orm:"label"      ` //
	Remark     string      `json:"remark"     orm:"remark"     ` //
	Relevance  int         `json:"relevance"  orm:"relevance"  ` //
}
