package model

import (
	"time"

	"github.com/gzdzh-cn/dzhcore"
)

const TableNameBaseSysNoticeUserRead = "base_sys_notice_user_read"

// BaseSysNoticeUserRead 模型，映射表 <base_sys_notice_user_read>
type BaseSysNoticeUserRead struct {
	*dzhcore.Model
	UserId   string    `gorm:"column:user_id;type:varchar(100);not null" json:"user_id"`      // 用户ID
	NoticeId string    `gorm:"column:notice_id;type:varchar(100);not null" json:"notice_id"`  // 消息ID
	Status   int       `gorm:"column:status;comment:状态;type:int(11);default:0" json:"status"` // 状态 0 未读，1 已读
	ReadTime time.Time `gorm:"column:readTime;type:datetime" json:"readTime"`                 // 阅读时间
}

// TableName BaseSysNoticeUserRead 的表名
func (*BaseSysNoticeUserRead) TableName() string {
	return TableNameBaseSysNoticeUserRead
}

// GroupName BaseSysNoticeUserRead 的表分组
func (*BaseSysNoticeUserRead) GroupName() string {
	return "default"
}

// NewBaseSysNoticeUserRead 创建一个新的 BaseSysNoticeUserRead 实例
func NewBaseSysNoticeUserRead() *BaseSysNoticeUserRead {
	return &BaseSysNoticeUserRead{
		Model: dzhcore.NewModel(),
	}
}

// init 注册模型
func init() {
	dzhcore.AddModel(&BaseSysNoticeUserRead{})
}
