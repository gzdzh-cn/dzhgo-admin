// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AddonsTaskLogDao is the data access object for table addons_task_log.
type AddonsTaskLogDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns AddonsTaskLogColumns // columns contains all the column names of Table for convenient usage.
}

// AddonsTaskLogColumns defines and stores column names for table addons_task_log.
type AddonsTaskLogColumns struct {
	Id         string //
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	DeletedAt  string //
	TaskId     string // 任务ID
	Status     string // 状态 0:失败 1:成功
	Detail     string // 详情
}

// addonsTaskLogColumns holds the columns for table addons_task_log.
var addonsTaskLogColumns = AddonsTaskLogColumns{
	Id:         "id",
	CreateTime: "createTime",
	UpdateTime: "updateTime",
	DeletedAt:  "deleted_at",
	TaskId:     "taskId",
	Status:     "status",
	Detail:     "detail",
}

// NewAddonsTaskLogDao creates and returns a new DAO object for table data access.
func NewAddonsTaskLogDao() *AddonsTaskLogDao {
	return &AddonsTaskLogDao{
		group:   "default",
		table:   "addons_task_log",
		columns: addonsTaskLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AddonsTaskLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AddonsTaskLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AddonsTaskLogDao) Columns() AddonsTaskLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AddonsTaskLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AddonsTaskLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AddonsTaskLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
