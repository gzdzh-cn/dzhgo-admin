// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BaseSysAddonsDao is the data access object for the table base_sys_addons.
type BaseSysAddonsDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  BaseSysAddonsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// BaseSysAddonsColumns defines and stores column names for the table base_sys_addons.
type BaseSysAddonsColumns struct {
	Id         string //
	CreateTime string //
	UpdateTime string //
	DeletedAt  string //
	Name       string //
	Image      string //
	Link       string //
	MenuId     string //
	TypeId     string //
	Remark     string //
	Status     string //
	OrderNum   string //
}

// baseSysAddonsColumns holds the columns for the table base_sys_addons.
var baseSysAddonsColumns = BaseSysAddonsColumns{
	Id:         "id",
	CreateTime: "createTime",
	UpdateTime: "updateTime",
	DeletedAt:  "deleted_at",
	Name:       "name",
	Image:      "image",
	Link:       "link",
	MenuId:     "menuId",
	TypeId:     "typeId",
	Remark:     "remark",
	Status:     "status",
	OrderNum:   "orderNum",
}

// NewBaseSysAddonsDao creates and returns a new DAO object for table data access.
func NewBaseSysAddonsDao(handlers ...gdb.ModelHandler) *BaseSysAddonsDao {
	return &BaseSysAddonsDao{
		group:    "default",
		table:    "base_sys_addons",
		columns:  baseSysAddonsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *BaseSysAddonsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *BaseSysAddonsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *BaseSysAddonsDao) Columns() BaseSysAddonsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *BaseSysAddonsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *BaseSysAddonsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *BaseSysAddonsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
