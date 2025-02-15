// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"dzhgo/addons/task/dao/internal"
)

// internalAddonsTaskInfoDao is internal type for wrapping internal DAO implements.
type internalAddonsTaskInfoDao = *internal.AddonsTaskInfoDao

// addonsTaskInfoDao is the data access object for table addons_task_info.
// You can define custom methods on it to extend its functionality as you wish.
type addonsTaskInfoDao struct {
	internalAddonsTaskInfoDao
}

var (
	// AddonsTaskInfo is globally public accessible object for table addons_task_info operations.
	AddonsTaskInfo = addonsTaskInfoDao{
		internal.NewAddonsTaskInfoDao(),
	}
)

// Fill with you ideas below.
