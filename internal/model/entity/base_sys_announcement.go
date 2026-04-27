// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BaseSysAnnouncement is the golang structure for table base_sys_announcement.
type BaseSysAnnouncement struct {
	Id         string      `json:"id"         orm:"id"         ` //
	CreateTime *gtime.Time `json:"createTime" orm:"createTime" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" orm:"updateTime" ` // 更新时间
	DeletedAt  *gtime.Time `json:"deletedAt"  orm:"deleted_at" ` //
	Title      string      `json:"title"      orm:"title"      ` //
	Content    string      `json:"content"    orm:"content"    ` //
	Type       int         `json:"type"       orm:"type"       ` //
	Status     int         `json:"status"     orm:"status"     ` //
	Top        int         `json:"top"        orm:"top"        ` //
	Remark     string      `json:"remark"     orm:"remark"     ` //
}
