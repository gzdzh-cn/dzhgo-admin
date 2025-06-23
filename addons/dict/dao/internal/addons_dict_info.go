// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AddonsDictInfoDao is the data access object for table addons_dict_info.
type AddonsDictInfoDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns AddonsDictInfoColumns // columns contains all the column names of Table for convenient usage.
}

// AddonsDictInfoColumns defines and stores column names for table addons_dict_info.
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

// addonsDictInfoColumns holds the columns for table addons_dict_info.
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
func NewAddonsDictInfoDao() *AddonsDictInfoDao {
	return &AddonsDictInfoDao{
		group:   "default",
		table:   "addons_dict_info",
		columns: addonsDictInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AddonsDictInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AddonsDictInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AddonsDictInfoDao) Columns() AddonsDictInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AddonsDictInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AddonsDictInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AddonsDictInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
