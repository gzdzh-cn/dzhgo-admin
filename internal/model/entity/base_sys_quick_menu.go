// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BaseSysQuickMenu is the golang structure for table base_sys_quick_menu.
type BaseSysQuickMenu struct {
	Id         string      `json:"id"         orm:"id"         ` //
	CreateTime *gtime.Time `json:"createTime" orm:"createTime" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" orm:"updateTime" ` // 更新时间
	DeletedAt  *gtime.Time `json:"deletedAt"  orm:"deleted_at" ` //
	Name       string      `json:"name"       orm:"name"       ` //
	Router     string      `json:"router"     orm:"router"     ` //
	Icon       string      `json:"icon"       orm:"icon"       ` //
	Status     int         `json:"status"     orm:"status"     ` // 状态
	OrderNum   int         `json:"orderNum"   orm:"order_num"  ` // 排序
	Remark     string      `json:"remark"     orm:"remark"     ` // 备注
	MenuId     string      `json:"menuId"     orm:"menu_id"    ` //
	UserId     string      `json:"userId"     orm:"user_id"    ` //
}
