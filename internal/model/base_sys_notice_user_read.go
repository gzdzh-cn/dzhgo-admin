package model

import (
	"time"

	"github.com/gzdzh-cn/dzhcore"
)

const TableNameBaseSysNoticeUserRead = "base_sys_notice_user_read"

// BaseSysNoticeUserRead mapped from table <base_sys_notice_user_read>
type BaseSysNoticeUserRead struct {
	*dzhcore.Model
	UserID    string    `gorm:"column:user_id;type:varchar(100);not null" json:"userId"`      // 用户ID
	NoticeID  string    `gorm:"column:notice_id;type:varchar(100);not null" json:"noticeId"`  // 消息ID
	Status    int       `gorm:"column:status;comment:状态;type:int(11);default:0" json:"status"` // 状态 0 未读，1 已读
	ReadTime  time.Time `gorm:"column:readTime;type:datetime" json:"readTime"`                 // 阅读时间
}

// TableName BaseSysNoticeUserRead's table name
func (*BaseSysNoticeUserRead) TableName() string {
	return TableNameBaseSysNoticeUserRead
}

// GroupName BaseSysNoticeUserRead's table group
func (*BaseSysNoticeUserRead) GroupName() string {
	return "default"
}

// NewBaseSysNoticeUserRead 创建一个新的 BaseSysNoticeUserRead 实例
func NewBaseSysNoticeUserRead() *BaseSysNoticeUserRead {
	return &BaseSysNoticeUserRead{
		Model: dzhcore.NewModel(),
	}
}

// init 创建表
func init() {
	// dzhcore.CreateTable(&BaseSysNoticeUserRead{})
	dzhcore.AddModel(&BaseSysNoticeUserRead{})
}
