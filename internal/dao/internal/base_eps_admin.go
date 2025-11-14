// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BaseEpsAdminDao is the data access object for the table base_eps_admin.
type BaseEpsAdminDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  BaseEpsAdminColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// BaseEpsAdminColumns defines and stores column names for the table base_eps_admin.
type BaseEpsAdminColumns struct {
	Id      string //
	Module  string //
	Method  string //
	Path    string //
	Prefix  string //
	Summary string //
	Tag     string //
	Dts     string //
}

// baseEpsAdminColumns holds the columns for the table base_eps_admin.
var baseEpsAdminColumns = BaseEpsAdminColumns{
	Id:      "id",
	Module:  "module",
	Method:  "method",
	Path:    "path",
	Prefix:  "prefix",
	Summary: "summary",
	Tag:     "tag",
	Dts:     "dts",
}

// NewBaseEpsAdminDao creates and returns a new DAO object for table data access.
func NewBaseEpsAdminDao(handlers ...gdb.ModelHandler) *BaseEpsAdminDao {
	return &BaseEpsAdminDao{
		group:    "default",
		table:    "base_eps_admin",
		columns:  baseEpsAdminColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *BaseEpsAdminDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *BaseEpsAdminDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *BaseEpsAdminDao) Columns() BaseEpsAdminColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *BaseEpsAdminDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *BaseEpsAdminDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *BaseEpsAdminDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
