// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BaseSysFeedbackDao is the data access object for the table base_sys_feedback.
type BaseSysFeedbackDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  BaseSysFeedbackColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// BaseSysFeedbackColumns defines and stores column names for the table base_sys_feedback.
type BaseSysFeedbackColumns struct {
	Id         string //
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	DeletedAt  string //
	UserId     string //
	Priority   string //
	FeType     string //
	Title      string //
	Img        string //
	Status     string // 状态
	Remark     string //
	Process    string // 处理状态
}

// baseSysFeedbackColumns holds the columns for the table base_sys_feedback.
var baseSysFeedbackColumns = BaseSysFeedbackColumns{
	Id:         "id",
	CreateTime: "createTime",
	UpdateTime: "updateTime",
	DeletedAt:  "deleted_at",
	UserId:     "user_id",
	Priority:   "priority",
	FeType:     "fe_type",
	Title:      "title",
	Img:        "img",
	Status:     "status",
	Remark:     "remark",
	Process:    "process",
}

// NewBaseSysFeedbackDao creates and returns a new DAO object for table data access.
func NewBaseSysFeedbackDao(handlers ...gdb.ModelHandler) *BaseSysFeedbackDao {
	return &BaseSysFeedbackDao{
		group:    "default",
		table:    "base_sys_feedback",
		columns:  baseSysFeedbackColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *BaseSysFeedbackDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *BaseSysFeedbackDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *BaseSysFeedbackDao) Columns() BaseSysFeedbackColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *BaseSysFeedbackDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *BaseSysFeedbackDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *BaseSysFeedbackDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
