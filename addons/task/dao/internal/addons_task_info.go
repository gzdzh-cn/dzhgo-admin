// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AddonsTaskInfoDao is the data access object for table addons_task_info.
type AddonsTaskInfoDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns AddonsTaskInfoColumns // columns contains all the column names of Table for convenient usage.
}

// AddonsTaskInfoColumns defines and stores column names for table addons_task_info.
type AddonsTaskInfoColumns struct {
	Id          string //
	CreateTime  string // 创建时间
	UpdateTime  string // 更新时间
	DeletedAt   string //
	JobId       string // 任务ID
	RepeatConf  string // 重复配置
	Name        string // 任务名称
	Cron        string // cron表达式
	Limit       string // 限制次数 不传为不限制
	Every       string // 间隔时间 单位秒
	Remark      string // 备注
	Status      string // 状态 0:关闭 1:开启
	StartDate   string // 开始时间
	EndDate     string // 结束时间
	Data        string // 数据
	Service     string // 执行的服务
	Type        string // 类型 0:系统 1:用户
	NextRunTime string // 下次执行时间
	TaskType    string // 任务类型 0:cron 1:时间间隔
}

// addonsTaskInfoColumns holds the columns for table addons_task_info.
var addonsTaskInfoColumns = AddonsTaskInfoColumns{
	Id:          "id",
	CreateTime:  "createTime",
	UpdateTime:  "updateTime",
	DeletedAt:   "deleted_at",
	JobId:       "jobId",
	RepeatConf:  "repeatConf",
	Name:        "name",
	Cron:        "cron",
	Limit:       "limit",
	Every:       "every",
	Remark:      "remark",
	Status:      "status",
	StartDate:   "startDate",
	EndDate:     "endDate",
	Data:        "data",
	Service:     "service",
	Type:        "type",
	NextRunTime: "nextRunTime",
	TaskType:    "taskType",
}

// NewAddonsTaskInfoDao creates and returns a new DAO object for table data access.
func NewAddonsTaskInfoDao() *AddonsTaskInfoDao {
	return &AddonsTaskInfoDao{
		group:   "default",
		table:   "addons_task_info",
		columns: addonsTaskInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AddonsTaskInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AddonsTaskInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AddonsTaskInfoDao) Columns() AddonsTaskInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AddonsTaskInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AddonsTaskInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AddonsTaskInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
