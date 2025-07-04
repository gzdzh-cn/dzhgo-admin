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
	CreateTime         string //
	UpdateTime         string //
	DeletedAt          string //
	SiteName           string //
	SiteDescribe       string //
	DomainName         string //
	Copyright          string //
	Logo               string //
	WxCode             string //
	Company            string //
	Contact            string //
	ContactWay         string //
	Mobile             string //
	Address            string //
	Keyword            string //
	Description        string //
	Smtp               string //
	SmtpEmail          string //
	SmtpPass           string //
	RemindEmail        string //
	IsRemindEmail      string //
	IsRemindSms        string //
	AccessKeyId        string //
	AccessKeySecret    string //
	SignName           string //
	TemplateCode       string //
	Endpoint           string //
	RemindMobile       string //
	RemindDay          string //
	FieldJson          string //
	Notice             string //
	Policy             string //
	Image              string //
	ContactList        string //
	BaiduTranApiKey    string //
	BaiduTranSecretKey string //
	CdnProxyUrl        string //
	Phrase             string //
	WxPayAppid         string //
	WxPayMchId         string //
	CAPIv3Key          string //
	CSerialNo          string //
	CNotifyUrl         string //
	SpMchid            string //
	SpAppid            string //
	SubMchId           string //
	APIv3Key           string //
	SerialNo           string //
	NotifyUrl          string //
	PayType            string //
	MpName             string //
	WxAppId            string //
	WxSecret           string //
	IsWpNotice         string //
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
