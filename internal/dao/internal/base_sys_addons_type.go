// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BaseSysAddonsTypeDao is the data access object for table base_sys_addonsType.
type BaseSysAddonsTypeDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns BaseSysAddonsTypeColumns // columns contains all the column names of Table for convenient usage.
}

// BaseSysAddonsTypeColumns defines and stores column names for table base_sys_addonsType.
type BaseSysAddonsTypeColumns struct {
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

// baseSysAddonsTypeColumns holds the columns for table base_sys_addonsType.
var baseSysAddonsTypeColumns = BaseSysAddonsTypeColumns{
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

// NewBaseSysAddonsTypeDao creates and returns a new DAO object for table data access.
func NewBaseSysAddonsTypeDao() *BaseSysAddonsTypeDao {
	return &BaseSysAddonsTypeDao{
		group:   "default",
		table:   "base_sys_addonsType",
		columns: baseSysAddonsTypeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BaseSysAddonsTypeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BaseSysAddonsTypeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BaseSysAddonsTypeDao) Columns() BaseSysAddonsTypeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BaseSysAddonsTypeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BaseSysAddonsTypeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BaseSysAddonsTypeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
