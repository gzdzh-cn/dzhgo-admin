// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BaseSysMenu is the golang structure for table base_sys_menu.
type BaseSysMenu struct {
	Id         string      `json:"id"         orm:"id"         ` //
	CreateTime *gtime.Time `json:"createTime" orm:"createTime" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" orm:"updateTime" ` // 更新时间
	DeletedAt  *gtime.Time `json:"deletedAt"  orm:"deleted_at" ` //
	ParentId   string      `json:"parentId"   orm:"parentId"   ` //
	Name       string      `json:"name"       orm:"name"       ` //
	Router     string      `json:"router"     orm:"router"     ` //
	Perms      string      `json:"perms"      orm:"perms"      ` //
	Type       int         `json:"type"       orm:"type"       ` //
	Icon       string      `json:"icon"       orm:"icon"       ` //
	OrderNum   int         `json:"orderNum"   orm:"orderNum"   ` //
	ViewPath   string      `json:"viewPath"   orm:"viewPath"   ` //
	KeepAlive  int         `json:"keepAlive"  orm:"keepAlive"  ` //
	IsShow     int         `json:"isShow"     orm:"isShow"     ` //
	IsInstall  int         `json:"isInstall"  orm:"isInstall"  ` //
	MenuType   string      `json:"menuType"   orm:"menuType"   ` //
	AddonsName string      `json:"addonsName" orm:"addonsName" ` //
}
