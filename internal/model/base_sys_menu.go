package model

import "github.com/gzdzh-cn/dzhcore"

const TableNameBaseSysMenu = "base_sys_menu"

// BaseSysMenu mapped from table <base_sys_menu>
type BaseSysMenu struct {
	*dzhcore.Model
	ParentID   string  `gorm:"column:parentId" json:"parentId"`                                        // 父菜单ID
	Name       string  `gorm:"column:name;type:varchar(255);not null" json:"name"`                     // 菜单名称
	Router     *string `gorm:"column:router;type:varchar(255)" json:"router"`                          // 菜单地址
	Perms      *string `gorm:"column:perms;type:varchar(255)" json:"perms"`                            // 权限标识
	Type       int32   `gorm:"column:type;not null" json:"type"`                                       // 类型 0：目录 1：菜单 2：按钮
	Icon       *string `gorm:"column:icon;type:varchar(255)" json:"icon"`                              // 图标
	OrderNum   int32   `gorm:"column:orderNum;type:int;not null;default:0" json:"orderNum"`            // 排序
	ViewPath   *string `gorm:"column:viewPath;type:varchar(255)" json:"viewPath"`                      // 视图地址
	KeepAlive  *int32  `gorm:"column:keepAlive;not null;default:1" json:"keepAlive"`                   // 路由缓存
	IsShow     *int32  `gorm:"column:isShow;not null;default:1" json:"isShow"`                         // 是否显示
	IsInstall  *int32  `gorm:"column:isInstall;not null;default:0" json:"isInstall"`                   // 是否安装
	MenuType   *string `gorm:"column:menuType;not null;varchar(255);default:base" json:"menuType"`     // 菜单类型 base:基础菜单 addon:插件菜单
	AddonsName *string `gorm:"column:addonsName;not null;varchar(255);default:base" json:"addonsName"` // 插件名称
}

// TableName BaseSysMenu's table name
func (*BaseSysMenu) TableName() string {
	return TableNameBaseSysMenu
}

// NewBaseSysMenu create a new BaseSysMenu
func NewBaseSysMenu() *BaseSysMenu {
	return &BaseSysMenu{
		Model: dzhcore.NewModel(),
	}
}

// init 创建表
func init() {
	// dzhcore.CreateTable(&BaseSysMenu{})
	dzhcore.AddModel(&BaseSysMenu{})
}
