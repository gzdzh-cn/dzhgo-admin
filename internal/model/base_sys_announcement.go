package model

import (
	"github.com/gzdzh-cn/dzhcore"
)

const TableNameAnnouncement = "base_sys_announcement"

// BaseSysAnnouncement 公告模型，映射表 <base_sys_announcement>
type BaseSysAnnouncement struct {
	*dzhcore.Model
	Title   string  `gorm:"column:title;type:varchar(255);not null" json:"title"`     // 公告标题
	Content string  `gorm:"column:content;type:text;not null" json:"content"`          // 公告内容
	Type    int     `gorm:"column:type;type:int(11);default:0" json:"type"`            // 类型 0:通知 1:公告
	Status  int     `gorm:"column:status;type:int(11);default:1" json:"status"`        // 状态 0:禁用 1:启用
	Top     int     `gorm:"column:top;type:int(11);default:0" json:"top"`              // 是否置顶 0:否 1:是
	Remark  *string `gorm:"column:remark;type:varchar(255)" json:"remark"`             // 备注
}

// TableName BaseSysAnnouncement 的表名
func (*BaseSysAnnouncement) TableName() string {
	return TableNameAnnouncement
}

// GroupName BaseSysAnnouncement 的表分组
func (*BaseSysAnnouncement) GroupName() string {
	return "default"
}

// NewBaseSysAnnouncement 创建一个新的 BaseSysAnnouncement 实例
func NewBaseSysAnnouncement() *BaseSysAnnouncement {
	return &BaseSysAnnouncement{
		Model: dzhcore.NewModel(),
	}
}

// init 注册模型
func init() {
	dzhcore.AddModel(&BaseSysAnnouncement{})
}
