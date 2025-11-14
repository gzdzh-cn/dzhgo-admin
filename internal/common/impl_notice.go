package common

import (
	"context"
)

// NoticeTarget 通知目标接口
type NoticeTarget interface {
	GetUserIDs(ctx context.Context) ([]string, error)
}

// SingleUser 单个用户
type SingleUser struct {
	ID string
}

func (s *SingleUser) GetUserIDs(ctx context.Context) ([]string, error) {
	if s.ID == "" {
		return []string{}, nil
	}
	return []string{s.ID}, nil
}

// MultiUser 多个用户
type MultiUser struct {
	IDs []string
}

func (m *MultiUser) GetUserIDs(ctx context.Context) ([]string, error) {
	return m.IDs, nil
}

// AllAdmins 全部管理员
type AllAdmins struct{}

func (a *AllAdmins) GetUserIDs(ctx context.Context) ([]string, error) {
	return getAdminUserIds(ctx)
}

// AllUsers 全部用户
type AllUsers struct{}

func (a *AllUsers) GetUserIDs(ctx context.Context) ([]string, error) {
	return getAllUserIds(ctx)
}

// 便捷构造函数
func ToUser(userId string) NoticeTarget {
	return &SingleUser{ID: userId}
}

func ToUsers(userIds []string) NoticeTarget {
	return &MultiUser{IDs: userIds}
}

func ToAllAdmins() NoticeTarget {
	return &AllAdmins{}
}

func ToAllUsers() NoticeTarget {
	return &AllUsers{}
}
