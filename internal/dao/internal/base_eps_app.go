// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BaseEpsAppDao is the data access object for table base_eps_app.
type BaseEpsAppDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns BaseEpsAppColumns // columns contains all the column names of Table for convenient usage.
}

// BaseEpsAppColumns defines and stores column names for table base_eps_app.
type BaseEpsAppColumns struct {
	Id      string //
	Module  string //
	Method  string //
	Path    string //
	Prefix  string //
	Summary string //
	Tag     string //
	Dts     string //
}

// baseEpsAppColumns holds the columns for table base_eps_app.
var baseEpsAppColumns = BaseEpsAppColumns{
	Id:      "id",
	Module:  "module",
	Method:  "method",
	Path:    "path",
	Prefix:  "prefix",
	Summary: "summary",
	Tag:     "tag",
	Dts:     "dts",
}

// NewBaseEpsAppDao creates and returns a new DAO object for table data access.
func NewBaseEpsAppDao() *BaseEpsAppDao {
	return &BaseEpsAppDao{
		group:   "default",
		table:   "base_eps_app",
		columns: baseEpsAppColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BaseEpsAppDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BaseEpsAppDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BaseEpsAppDao) Columns() BaseEpsAppColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BaseEpsAppDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BaseEpsAppDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BaseEpsAppDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
