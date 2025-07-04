// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BaseSysConf is the golang structure for table base_sys_conf.
type BaseSysConf struct {
	Id         string      `json:"id"         orm:"id"         ` //
	CreateTime *gtime.Time `json:"createTime" orm:"createTime" ` //
	UpdateTime *gtime.Time `json:"updateTime" orm:"updateTime" ` //
	DeletedAt  *gtime.Time `json:"deletedAt"  orm:"deleted_at" ` //
	CKey       string      `json:"cKey"       orm:"cKey"       ` //
	CValue     string      `json:"cValue"     orm:"cValue"     ` //
}
