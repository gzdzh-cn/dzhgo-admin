// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AddonsMemberAttrDao is the data access object for the table addons_member_attr.
type AddonsMemberAttrDao struct {
	table    string                  // table is the underlying table name of the DAO.
	group    string                  // group is the database configuration group name of the current DAO.
	columns  AddonsMemberAttrColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler      // handlers for customized model modification.
}

// AddonsMemberAttrColumns defines and stores column names for the table addons_member_attr.
type AddonsMemberAttrColumns struct {
	Id             string //
	CreateTime     string // 创建时间
	UpdateTime     string // 更新时间
	DeletedAt      string //
	Type           string // 类型:1=公众号,2=小程序
	Unionid        string // UnionID
	Notify         string // 微信通知:1=是,0=否
	Openid         string // openid
	Nickname       string // 昵称
	Sex            string // 性别:0=默认,1=男,2=女
	Language       string // 语言
	Country        string // 国家
	Province       string // 省份
	City           string // 城市
	Headimgurl     string // 头像
	Subscribe      string // 关注状态:0=未关注,1=已关注
	SubscribeTime  string // 关注时间
	Groupid        string // 分组ID
	TagidList      string // 标签列表
	Privilege      string // 权限
	SubscribeScene string // 关注来源
	QrScene        string // 扫码场景
	QrSceneStr     string // 扫码场景描述
	MemberId       string // 会员ID
}

// addonsMemberAttrColumns holds the columns for the table addons_member_attr.
var addonsMemberAttrColumns = AddonsMemberAttrColumns{
	Id:             "id",
	CreateTime:     "createTime",
	UpdateTime:     "updateTime",
	DeletedAt:      "deleted_at",
	Type:           "type",
	Unionid:        "unionid",
	Notify:         "notify",
	Openid:         "openid",
	Nickname:       "nickname",
	Sex:            "sex",
	Language:       "language",
	Country:        "country",
	Province:       "province",
	City:           "city",
	Headimgurl:     "headimgurl",
	Subscribe:      "subscribe",
	SubscribeTime:  "subscribe_time",
	Groupid:        "groupid",
	TagidList:      "tagid_list",
	Privilege:      "privilege",
	SubscribeScene: "subscribe_scene",
	QrScene:        "qr_scene",
	QrSceneStr:     "qr_scene_str",
	MemberId:       "member_id",
}

// NewAddonsMemberAttrDao creates and returns a new DAO object for table data access.
func NewAddonsMemberAttrDao(handlers ...gdb.ModelHandler) *AddonsMemberAttrDao {
	return &AddonsMemberAttrDao{
		group:    "default",
		table:    "addons_member_attr",
		columns:  addonsMemberAttrColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AddonsMemberAttrDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AddonsMemberAttrDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AddonsMemberAttrDao) Columns() AddonsMemberAttrColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AddonsMemberAttrDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AddonsMemberAttrDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AddonsMemberAttrDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
