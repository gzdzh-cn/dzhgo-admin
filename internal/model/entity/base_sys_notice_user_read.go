// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BaseSysNoticeUserRead is the golang structure for table base_sys_notice_user_read.
type BaseSysNoticeUserRead struct {
	Id         string      `json:"id"         orm:"id"         ` //
	CreateTime *gtime.Time `json:"createTime" orm:"createTime" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" orm:"updateTime" ` // 更新时间
	DeletedAt  *gtime.Time `json:"deletedAt"  orm:"deleted_at" ` //
	UserId     string      `json:"userId"     orm:"user_id"    ` //
	NoticeId   string      `json:"noticeId"   orm:"notice_id"  ` //
	Status     int         `json:"status"     orm:"status"     ` // 状态
	ReadTime   *gtime.Time `json:"readTime"   orm:"readTime"   ` //
}
