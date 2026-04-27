// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BaseSysAnnouncementRead is the golang structure for table base_sys_announcement_read.
type BaseSysAnnouncementRead struct {
	Id             string      `json:"id"             orm:"id"              ` //
	CreateTime     *gtime.Time `json:"createTime"     orm:"createTime"      ` // 创建时间
	UpdateTime     *gtime.Time `json:"updateTime"     orm:"updateTime"      ` // 更新时间
	DeletedAt      *gtime.Time `json:"deletedAt"      orm:"deleted_at"      ` //
	UserId         string      `json:"userId"         orm:"user_id"         ` //
	AnnouncementId string      `json:"announcementId" orm:"announcement_id" ` //
}
