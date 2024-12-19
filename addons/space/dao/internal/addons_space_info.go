// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AddonsSpaceInfoDao is the data access object for table addons_space_info.
type AddonsSpaceInfoDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns AddonsSpaceInfoColumns // columns contains all the column names of Table for convenient usage.
}

// AddonsSpaceInfoColumns defines and stores column names for table addons_space_info.
type AddonsSpaceInfoColumns struct {
	Id         string //
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	DeletedAt  string //
	Url        string // 地址
	Type       string // 类型
	ClassifyId string // 分类ID
}

// addonsSpaceInfoColumns holds the columns for table addons_space_info.
var addonsSpaceInfoColumns = AddonsSpaceInfoColumns{
	Id:         "id",
	CreateTime: "createTime",
	UpdateTime: "updateTime",
	DeletedAt:  "deleted_at",
	Url:        "url",
	Type:       "type",
	ClassifyId: "classifyId",
}

// NewAddonsSpaceInfoDao creates and returns a new DAO object for table data access.
func NewAddonsSpaceInfoDao() *AddonsSpaceInfoDao {
	return &AddonsSpaceInfoDao{
		group:   "default",
		table:   "addons_space_info",
		columns: addonsSpaceInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AddonsSpaceInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AddonsSpaceInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AddonsSpaceInfoDao) Columns() AddonsSpaceInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AddonsSpaceInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AddonsSpaceInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AddonsSpaceInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
