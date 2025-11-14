// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BaseSysAddonsTypesDao is the data access object for the table base_sys_addons_types.
type BaseSysAddonsTypesDao struct {
	table    string                    // table is the underlying table name of the DAO.
	group    string                    // group is the database configuration group name of the current DAO.
	columns  BaseSysAddonsTypesColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler        // handlers for customized model modification.
}

// BaseSysAddonsTypesColumns defines and stores column names for the table base_sys_addons_types.
type BaseSysAddonsTypesColumns struct {
	Id         string //
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	DeletedAt  string //
	Name       string // 标题
	Image      string // 图片
	Link       string // 跳转
	Remark     string // 备注
	Status     string // 状态
	OrderNum   string // 排序
}

// baseSysAddonsTypesColumns holds the columns for the table base_sys_addons_types.
var baseSysAddonsTypesColumns = BaseSysAddonsTypesColumns{
	Id:         "id",
	CreateTime: "createTime",
	UpdateTime: "updateTime",
	DeletedAt:  "deleted_at",
	Name:       "name",
	Image:      "image",
	Link:       "link",
	Remark:     "remark",
	Status:     "status",
	OrderNum:   "orderNum",
}

// NewBaseSysAddonsTypesDao creates and returns a new DAO object for table data access.
func NewBaseSysAddonsTypesDao(handlers ...gdb.ModelHandler) *BaseSysAddonsTypesDao {
	return &BaseSysAddonsTypesDao{
		group:    "default",
		table:    "base_sys_addons_types",
		columns:  baseSysAddonsTypesColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *BaseSysAddonsTypesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *BaseSysAddonsTypesDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *BaseSysAddonsTypesDao) Columns() BaseSysAddonsTypesColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *BaseSysAddonsTypesDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *BaseSysAddonsTypesDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *BaseSysAddonsTypesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
