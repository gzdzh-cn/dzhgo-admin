package model

import (
	"github.com/gzdzh-cn/dzhcore"
)

const TableNameBaseSysFeedback = "base_sys_feedback"

// BaseSysFeedback 模型，映射表 <base_sys_feedback>
type BaseSysFeedback struct {
	*dzhcore.Model
	UserId   string  `gorm:"column:user_id;type:varchar(100);not null" json:"user_id"`               // 用户ID
	Priority string  `gorm:"column:priority;type:varchar(100);default:1" json:"priority"`            // 优先级
	FeType   string  `gorm:"column:fe_type;type:varchar(100);default:bug" json:"feType"`             // 反馈类型
	Title    string  `gorm:"column:title;type:varchar(255);not null" json:"title"`                   // 标题
	Img      *string `gorm:"column:img;type:varchar(255)" json:"img"`                                // 图片
	Status   int     `gorm:"column:status;comment:状态;type:int(11);default:1" json:"status"`          // 状态
	Remark   *string `gorm:"column:remark;type:varchar(255)" json:"remark"`                          // 反馈内容
	Process  string  `gorm:"column:process;comment:处理状态;type:varchar(100);default:1" json:"process"` // 处理状态 1:未处理 2:处理中 3:已处理
}

// TableName BaseSysFeedback 的表名
func (*BaseSysFeedback) TableName() string {
	return TableNameBaseSysFeedback
}

// GroupName BaseSysFeedback 的表分组
func (*BaseSysFeedback) GroupName() string {
	return "default"
}

// NewBaseSysFeedback 创建一个新的 BaseSysFeedback 实例
func NewBaseSysFeedback() *BaseSysFeedback {
	return &BaseSysFeedback{
		Model: dzhcore.NewModel(),
	}
}

// init 注册模型
func init() {
	dzhcore.AddModel(&BaseSysFeedback{})
}
