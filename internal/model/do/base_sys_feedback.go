// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BaseSysFeedback is the golang structure of table base_sys_feedback for DAO operations like Where/Data.
type BaseSysFeedback struct {
	g.Meta     `orm:"table:base_sys_feedback, do:true"`
	Id         interface{} //
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
	DeletedAt  *gtime.Time //
	UserId     interface{} //
	Priority   interface{} //
	FeType     interface{} //
	Title      interface{} //
	Img        interface{} //
	Status     interface{} // 状态
	Remark     interface{} //
	Process    interface{} // 处理状态
}
