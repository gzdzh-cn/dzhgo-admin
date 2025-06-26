// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AddonsMemberAttrDao is the data access object for table addons_member_attr.
type AddonsMemberAttrDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns AddonsMemberAttrColumns // columns contains all the column names of Table for convenient usage.
}

// AddonsMemberAttrColumns defines and stores column names for table addons_member_attr.
type AddonsMemberAttrColumns struct {
	Id             string //
	CreateTime     string // 创建时间
	UpdateTime     string // 更新时间
	DeletedAt      string //
	Type           string // 类型:1=公众号,2=小程序
	UserId         string // 用户ID
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
}

// addonsMemberAttrColumns holds the columns for table addons_member_attr.
var addonsMemberAttrColumns = AddonsMemberAttrColumns{
	Id:             "id",
	CreateTime:     "createTime",
	UpdateTime:     "updateTime",
	DeletedAt:      "deleted_at",
	Type:           "type",
	UserId:         "user_id",
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
}

// NewAddonsMemberAttrDao creates and returns a new DAO object for table data access.
func NewAddonsMemberAttrDao() *AddonsMemberAttrDao {
	return &AddonsMemberAttrDao{
		group:   "default",
		table:   "addons_member_attr",
		columns: addonsMemberAttrColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AddonsMemberAttrDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AddonsMemberAttrDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AddonsMemberAttrDao) Columns() AddonsMemberAttrColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AddonsMemberAttrDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AddonsMemberAttrDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AddonsMemberAttrDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
