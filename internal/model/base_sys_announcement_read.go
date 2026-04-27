package model

import (
	"github.com/gzdzh-cn/dzhcore"
)

const TableNameBaseSysAnnouncementRead = "base_sys_announcement_read"

// BaseSysAnnouncementRead 模型，映射表 <base_sys_announcement_read>
type BaseSysAnnouncementRead struct {
	*dzhcore.Model
	UserId         string `gorm:"column:user_id;type:varchar(64);not null" json:"userId"`         // 用户ID
	AnnouncementId string `gorm:"column:announcement_id;type:varchar(64);not null" json:"announcementId"` // 公告ID
}

// TableName BaseSysAnnouncementRead 的表名
func (*BaseSysAnnouncementRead) TableName() string {
	return TableNameBaseSysAnnouncementRead
}

// GroupName BaseSysAnnouncementRead 的表分组
func (*BaseSysAnnouncementRead) GroupName() string {
	return "default"
}

// NewBaseSysAnnouncementRead 创建一个新的 BaseSysAnnouncementRead 实例
func NewBaseSysAnnouncementRead() *BaseSysAnnouncementRead {
	return &BaseSysAnnouncementRead{
		Model: dzhcore.NewModel(),
	}
}

// init 注册模型
func init() {
	dzhcore.AddModel(&BaseSysAnnouncementRead{})
}
