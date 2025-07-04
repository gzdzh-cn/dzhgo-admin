// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AddonsDictInfoDao is the data access object for the table addons_dict_info.
type AddonsDictInfoDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  AddonsDictInfoColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// AddonsDictInfoColumns defines and stores column names for the table addons_dict_info.
type AddonsDictInfoColumns struct {
	Id         string //
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	DeletedAt  string //
	TypeId     string //
	Name       string //
	Value      string //
	OrderNum   string //
	Remark     string //
	ParentId   string //
}

// addonsDictInfoColumns holds the columns for the table addons_dict_info.
var addonsDictInfoColumns = AddonsDictInfoColumns{
	Id:         "id",
	CreateTime: "createTime",
	UpdateTime: "updateTime",
	DeletedAt:  "deleted_at",
	TypeId:     "typeId",
	Name:       "name",
	Value:      "value",
	OrderNum:   "orderNum",
	Remark:     "remark",
	ParentId:   "parentId",
}

// NewAddonsDictInfoDao creates and returns a new DAO object for table data access.
func NewAddonsDictInfoDao(handlers ...gdb.ModelHandler) *AddonsDictInfoDao {
	return &AddonsDictInfoDao{
		group:    "default",
		table:    "addons_dict_info",
		columns:  addonsDictInfoColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AddonsDictInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AddonsDictInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AddonsDictInfoDao) Columns() AddonsDictInfoColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AddonsDictInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AddonsDictInfoDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *AddonsDictInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
