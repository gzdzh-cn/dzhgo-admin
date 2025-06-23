// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AddonsSpaceTypeDao is the data access object for table addons_space_type.
type AddonsSpaceTypeDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns AddonsSpaceTypeColumns // columns contains all the column names of Table for convenient usage.
}

// AddonsSpaceTypeColumns defines and stores column names for table addons_space_type.
type AddonsSpaceTypeColumns struct {
	Id         string //
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	DeletedAt  string //
	Name       string // 类别名称
	ParentId   string // 父分类ID
}

// addonsSpaceTypeColumns holds the columns for table addons_space_type.
var addonsSpaceTypeColumns = AddonsSpaceTypeColumns{
	Id:         "id",
	CreateTime: "createTime",
	UpdateTime: "updateTime",
	DeletedAt:  "deleted_at",
	Name:       "name",
	ParentId:   "parentId",
}

// NewAddonsSpaceTypeDao creates and returns a new DAO object for table data access.
func NewAddonsSpaceTypeDao() *AddonsSpaceTypeDao {
	return &AddonsSpaceTypeDao{
		group:   "default",
		table:   "addons_space_type",
		columns: addonsSpaceTypeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AddonsSpaceTypeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AddonsSpaceTypeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AddonsSpaceTypeDao) Columns() AddonsSpaceTypeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AddonsSpaceTypeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AddonsSpaceTypeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AddonsSpaceTypeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
