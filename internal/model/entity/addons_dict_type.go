// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AddonsDictType is the golang structure for table addons_dict_type.
type AddonsDictType struct {
	Id         string      `json:"id"         orm:"id"         ` //
	CreateTime *gtime.Time `json:"createTime" orm:"createTime" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" orm:"updateTime" ` // 更新时间
	DeletedAt  *gtime.Time `json:"deletedAt"  orm:"deleted_at" ` //
	Name       string      `json:"name"       orm:"name"       ` //
	Key        string      `json:"key"        orm:"key"        ` //
}
