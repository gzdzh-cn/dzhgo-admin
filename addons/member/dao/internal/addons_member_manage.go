// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AddonsMemberManageDao is the data access object for the table addons_member_manage.
type AddonsMemberManageDao struct {
	table    string                    // table is the underlying table name of the DAO.
	group    string                    // group is the database configuration group name of the current DAO.
	columns  AddonsMemberManageColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler        // handlers for customized model modification.
}

// AddonsMemberManageColumns defines and stores column names for the table addons_member_manage.
type AddonsMemberManageColumns struct {
	Id            string //
	CreateTime    string // 创建时间
	UpdateTime    string // 更新时间
	DeletedAt     string //
	AvatarUrl     string // 头像
	Password      string // 会员密码
	PasswordV     string //
	Nickname      string // 会员昵称
	LevelName     string // 等级名称
	Level         string // 等级
	Sex           string // 性别
	Qq            string // qq
	Mobile        string // 手机号
	Wx            string // 微信号
	WxImg         string // 微信二维码
	Email         string // email
	Role          string // 家庭角色
	LastLoginTime string // 最后登录时间
	Openid        string // openid
	UnionId       string // unionId
	SessionKey    string // session_key
	Remark        string // 备注
	Status        string //
	Description   string // 描述
	UserId        string // 多租户从属 id
	MemberName    string // 会员
}

// addonsMemberManageColumns holds the columns for the table addons_member_manage.
var addonsMemberManageColumns = AddonsMemberManageColumns{
	Id:            "id",
	CreateTime:    "createTime",
	UpdateTime:    "updateTime",
	DeletedAt:     "deleted_at",
	AvatarUrl:     "avatarUrl",
	Password:      "password",
	PasswordV:     "passwordV",
	Nickname:      "nickname",
	LevelName:     "levelName",
	Level:         "level",
	Sex:           "sex",
	Qq:            "qq",
	Mobile:        "mobile",
	Wx:            "wx",
	WxImg:         "wxImg",
	Email:         "email",
	Role:          "role",
	LastLoginTime: "lastLoginTime",
	Openid:        "openid",
	UnionId:       "unionId",
	SessionKey:    "session_key",
	Remark:        "remark",
	Status:        "status",
	Description:   "description",
	UserId:        "user_id",
	MemberName:    "member_name",
}

// NewAddonsMemberManageDao creates and returns a new DAO object for table data access.
func NewAddonsMemberManageDao(handlers ...gdb.ModelHandler) *AddonsMemberManageDao {
	return &AddonsMemberManageDao{
		group:    "default",
		table:    "addons_member_manage",
		columns:  addonsMemberManageColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AddonsMemberManageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AddonsMemberManageDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AddonsMemberManageDao) Columns() AddonsMemberManageColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AddonsMemberManageDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AddonsMemberManageDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AddonsMemberManageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
