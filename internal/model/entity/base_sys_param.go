// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BaseSysParam is the golang structure for table base_sys_param.
type BaseSysParam struct {
	Id         string      `json:"id"         orm:"id"         ` //
	CreateTime *gtime.Time `json:"createTime" orm:"createTime" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" orm:"updateTime" ` // 更新时间
	DeletedAt  *gtime.Time `json:"deletedAt"  orm:"deleted_at" ` //
	KeyName    string      `json:"keyName"    orm:"keyName"    ` //
	Name       string      `json:"name"       orm:"name"       ` //
	Data       string      `json:"data"       orm:"data"       ` //
	DataType   int         `json:"dataType"   orm:"dataType"   ` //
	Remark     string      `json:"remark"     orm:"remark"     ` //
}
