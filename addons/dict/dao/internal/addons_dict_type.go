// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AddonsDictTypeDao is the data access object for table addons_dict_type.
type AddonsDictTypeDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns AddonsDictTypeColumns // columns contains all the column names of Table for convenient usage.
}

// AddonsDictTypeColumns defines and stores column names for table addons_dict_type.
type AddonsDictTypeColumns struct {
	Id         string //
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	DeletedAt  string //
	Name       string //
	Key        string //
}

// addonsDictTypeColumns holds the columns for table addons_dict_type.
var addonsDictTypeColumns = AddonsDictTypeColumns{
	Id:         "id",
	CreateTime: "createTime",
	UpdateTime: "updateTime",
	DeletedAt:  "deleted_at",
	Name:       "name",
	Key:        "key",
}

// NewAddonsDictTypeDao creates and returns a new DAO object for table data access.
func NewAddonsDictTypeDao() *AddonsDictTypeDao {
	return &AddonsDictTypeDao{
		group:   "default",
		table:   "addons_dict_type",
		columns: addonsDictTypeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AddonsDictTypeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AddonsDictTypeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AddonsDictTypeDao) Columns() AddonsDictTypeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AddonsDictTypeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AddonsDictTypeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AddonsDictTypeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
