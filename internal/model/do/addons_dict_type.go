// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AddonsDictType is the golang structure of table addons_dict_type for DAO operations like Where/Data.
type AddonsDictType struct {
	g.Meta     `orm:"table:addons_dict_type, do:true"`
	Id         interface{} //
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
	DeletedAt  *gtime.Time //
	Name       interface{} //
	Key        interface{} //
}
