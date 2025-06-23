// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AddonsDictInfo is the golang structure for table addons_dict_info.
type AddonsDictInfo struct {
	Id         string      `json:"id"         orm:"id"         ` //
	CreateTime *gtime.Time `json:"createTime" orm:"createTime" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" orm:"updateTime" ` // 更新时间
	DeletedAt  *gtime.Time `json:"deletedAt"  orm:"deleted_at" ` //
	TypeId     string      `json:"typeId"     orm:"typeId"     ` //
	Name       string      `json:"name"       orm:"name"       ` //
	Value      string      `json:"value"      orm:"value"      ` //
	OrderNum   int         `json:"orderNum"   orm:"orderNum"   ` //
	Remark     string      `json:"remark"     orm:"remark"     ` //
	ParentId   int         `json:"parentId"   orm:"parentId"   ` //
}
