// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BaseSysSettingDao is the data access object for the table base_sys_setting.
type BaseSysSettingDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  BaseSysSettingColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// BaseSysSettingColumns defines and stores column names for the table base_sys_setting.
type BaseSysSettingColumns struct {
	Id                 string //
	CreateTime         string // 创建时间
	UpdateTime         string // 更新时间
	DeletedAt          string //
	SiteName           string // 站点名称
	SiteDescribe       string // 站点介绍
	DomainName         string // 网站域名
	Copyright          string // 版权所有
	Logo               string // logo
	WxCode             string // 二维码
	Company            string // 公司名称
	Contact            string // 联系人
	ContactWay         string // 座机
	Mobile             string // 手机
	Address            string // 地址
	Keyword            string // 关键词
	Description        string // 描述
	Smtp               string // smtp
	SmtpEmail          string // 发送邮箱
	SmtpPass           string // 邮箱授权码
	RemindEmail        string // 接收邮箱
	IsRemindEmail      string // 到期邮件开启 0关闭 1开启
	IsRemindSms        string // 到期短信开启 0关闭 1开启
	AccessKeyId        string // accessKeyId
	AccessKeySecret    string // accessKeySecret
	SignName           string // 签名
	TemplateCode       string // 模板
	Endpoint           string // endpoint
	RemindMobile       string // 通知手机号码
	RemindDay          string // 到期提醒提前天数
	FieldJson          string // 自定义字段
	Notice             string // 公告
	Policy             string // 隐私政策
	Image              string // 图片
	ContactList        string // 客服列表
	BaiduTranApiKey    string // 百度翻译apikey
	BaiduTranSecretKey string // 百度翻译Secretkey
	CdnProxyUrl        string // 图片代理地址
	Phrase             string // 过滤词
	WxPayAppid         string // 普通商户appid
	WxPayMchId         string // 普通商户号
	CAPIv3Key          string // 收款商户v3密钥
	CSerialNo          string // 序列号
	CNotifyUrl         string // 支付回调地址
	SpMchid            string // 服务商商户号
	SpAppid            string // 服务商appid
	SubMchId           string // 特约商户
	APIv3Key           string // 收款商户v3密钥
	SerialNo           string // 序列号
	NotifyUrl          string // 支付回调地址
	PayType            string // 支付模式
	MpName             string // 公众号名称
	WxAppId            string // 公众号appId
	WxSecret           string // 微信secret
	IsWpNotice         string // 公众号通知
}

// baseSysSettingColumns holds the columns for the table base_sys_setting.
var baseSysSettingColumns = BaseSysSettingColumns{
	Id:                 "id",
	CreateTime:         "createTime",
	UpdateTime:         "updateTime",
	DeletedAt:          "deleted_at",
	SiteName:           "siteName",
	SiteDescribe:       "siteDescribe",
	DomainName:         "domainName",
	Copyright:          "copyright",
	Logo:               "logo",
	WxCode:             "wxCode",
	Company:            "company",
	Contact:            "contact",
	ContactWay:         "contactWay",
	Mobile:             "mobile",
	Address:            "Address",
	Keyword:            "keyword",
	Description:        "description",
	Smtp:               "smtp",
	SmtpEmail:          "smtpEmail",
	SmtpPass:           "smtpPass",
	RemindEmail:        "remindEmail",
	IsRemindEmail:      "isRemindEmail",
	IsRemindSms:        "isRemindSms",
	AccessKeyId:        "accessKeyId",
	AccessKeySecret:    "accessKeySecret",
	SignName:           "signName",
	TemplateCode:       "templateCode",
	Endpoint:           "endpoint",
	RemindMobile:       "remindMobile",
	RemindDay:          "remindDay",
	FieldJson:          "fieldJson",
	Notice:             "notice",
	Policy:             "policy",
	Image:              "image",
	ContactList:        "contactList",
	BaiduTranApiKey:    "baiduTranApiKey",
	BaiduTranSecretKey: "baiduTranSecretKey",
	CdnProxyUrl:        "cdn_proxy_url",
	Phrase:             "phrase",
	WxPayAppid:         "wxPayAppid",
	WxPayMchId:         "wxPayMchId",
	CAPIv3Key:          "cAPIv3Key",
	CSerialNo:          "cSerialNo",
	CNotifyUrl:         "cNotifyUrl",
	SpMchid:            "spMchid",
	SpAppid:            "spAppid",
	SubMchId:           "subMchId",
	APIv3Key:           "aPIv3Key",
	SerialNo:           "serialNo",
	NotifyUrl:          "notifyUrl",
	PayType:            "payType",
	MpName:             "mpName",
	WxAppId:            "wxAppId",
	WxSecret:           "wxSecret",
	IsWpNotice:         "isWpNotice",
}

// NewBaseSysSettingDao creates and returns a new DAO object for table data access.
func NewBaseSysSettingDao(handlers ...gdb.ModelHandler) *BaseSysSettingDao {
	return &BaseSysSettingDao{
		group:    "default",
		table:    "base_sys_setting",
		columns:  baseSysSettingColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *BaseSysSettingDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *BaseSysSettingDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *BaseSysSettingDao) Columns() BaseSysSettingColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *BaseSysSettingDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *BaseSysSettingDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *BaseSysSettingDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
