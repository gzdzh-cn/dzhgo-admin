// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BaseSysAddons is the golang structure for table base_sys_addons.
type BaseSysAddons struct {
	Id         string      `json:"id"         orm:"id"         ` //
	CreateTime *gtime.Time `json:"createTime" orm:"createTime" ` //
	UpdateTime *gtime.Time `json:"updateTime" orm:"updateTime" ` //
	DeletedAt  *gtime.Time `json:"deletedAt"  orm:"deleted_at" ` //
	Name       string      `json:"name"       orm:"name"       ` //
	Image      string      `json:"image"      orm:"image"      ` //
	Link       string      `json:"link"       orm:"link"       ` //
	MenuId     string      `json:"menuId"     orm:"menuId"     ` //
	TypeId     string      `json:"typeId"     orm:"typeId"     ` //
	Remark     string      `json:"remark"     orm:"remark"     ` //
	Status     int         `json:"status"     orm:"status"     ` //
	OrderNum   int         `json:"orderNum"   orm:"orderNum"   ` //
}
