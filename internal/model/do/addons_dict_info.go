// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AddonsDictInfo is the golang structure of table addons_dict_info for DAO operations like Where/Data.
type AddonsDictInfo struct {
	g.Meta     `orm:"table:addons_dict_info, do:true"`
	Id         interface{} //
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
	DeletedAt  *gtime.Time //
	TypeId     interface{} //
	Name       interface{} //
	Value      interface{} //
	OrderNum   interface{} //
	Remark     interface{} //
	ParentId   interface{} //
}
