// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BaseSysNoticeUserReadDao is the data access object for the table base_sys_notice_user_read.
type BaseSysNoticeUserReadDao struct {
	table    string                       // table is the underlying table name of the DAO.
	group    string                       // group is the database configuration group name of the current DAO.
	columns  BaseSysNoticeUserReadColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler           // handlers for customized model modification.
}

// BaseSysNoticeUserReadColumns defines and stores column names for the table base_sys_notice_user_read.
type BaseSysNoticeUserReadColumns struct {
	Id         string //
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	DeletedAt  string //
	UserId     string //
	NoticeId   string //
	Status     string // 状态
	ReadTime   string //
}

// baseSysNoticeUserReadColumns holds the columns for the table base_sys_notice_user_read.
var baseSysNoticeUserReadColumns = BaseSysNoticeUserReadColumns{
	Id:         "id",
	CreateTime: "createTime",
	UpdateTime: "updateTime",
	DeletedAt:  "deleted_at",
	UserId:     "user_id",
	NoticeId:   "notice_id",
	Status:     "status",
	ReadTime:   "readTime",
}

// NewBaseSysNoticeUserReadDao creates and returns a new DAO object for table data access.
func NewBaseSysNoticeUserReadDao(handlers ...gdb.ModelHandler) *BaseSysNoticeUserReadDao {
	return &BaseSysNoticeUserReadDao{
		group:    "default",
		table:    "base_sys_notice_user_read",
		columns:  baseSysNoticeUserReadColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *BaseSysNoticeUserReadDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *BaseSysNoticeUserReadDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *BaseSysNoticeUserReadDao) Columns() BaseSysNoticeUserReadColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *BaseSysNoticeUserReadDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *BaseSysNoticeUserReadDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *BaseSysNoticeUserReadDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
