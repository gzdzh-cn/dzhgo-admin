// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BaseSysFeedback is the golang structure for table base_sys_feedback.
type BaseSysFeedback struct {
	Id         string      `json:"id"         orm:"id"         ` //
	CreateTime *gtime.Time `json:"createTime" orm:"createTime" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" orm:"updateTime" ` // 更新时间
	DeletedAt  *gtime.Time `json:"deletedAt"  orm:"deleted_at" ` //
	UserId     string      `json:"userId"     orm:"user_id"    ` //
	Priority   string      `json:"priority"   orm:"priority"   ` //
	FeType     string      `json:"feType"     orm:"fe_type"    ` //
	Title      string      `json:"title"      orm:"title"      ` //
	Img        string      `json:"img"        orm:"img"        ` //
	Status     int         `json:"status"     orm:"status"     ` // 状态
	Remark     string      `json:"remark"     orm:"remark"     ` //
	Process    string      `json:"process"    orm:"process"    ` // 处理状态
}
