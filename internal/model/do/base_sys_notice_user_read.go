// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BaseSysNoticeUserRead is the golang structure of table base_sys_notice_user_read for DAO operations like Where/Data.
type BaseSysNoticeUserRead struct {
	g.Meta     `orm:"table:base_sys_notice_user_read, do:true"`
	Id         interface{} //
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
	DeletedAt  *gtime.Time //
	UserId     interface{} //
	NoticeId   interface{} //
	Status     interface{} // 状态
	ReadTime   *gtime.Time //
}
