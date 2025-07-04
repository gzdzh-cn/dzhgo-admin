// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BaseSysSetting is the golang structure of table base_sys_setting for DAO operations like Where/Data.
type BaseSysSetting struct {
	g.Meta             `orm:"table:base_sys_setting, do:true"`
	Id                 interface{} //
	CreateTime         *gtime.Time //
	UpdateTime         *gtime.Time //
	DeletedAt          *gtime.Time //
	SiteName           interface{} //
	SiteDescribe       interface{} //
	DomainName         interface{} //
	Copyright          interface{} //
	Logo               interface{} //
	WxCode             interface{} //
	Company            interface{} //
	Contact            interface{} //
	ContactWay         interface{} //
	Mobile             interface{} //
	Address            interface{} //
	Keyword            interface{} //
	Description        interface{} //
	Smtp               interface{} //
	SmtpEmail          interface{} //
	SmtpPass           interface{} //
	RemindEmail        interface{} //
	IsRemindEmail      interface{} //
	IsRemindSms        interface{} //
	AccessKeyId        interface{} //
	AccessKeySecret    interface{} //
	SignName           interface{} //
	TemplateCode       interface{} //
	Endpoint           interface{} //
	RemindMobile       interface{} //
	RemindDay          interface{} //
	FieldJson          interface{} //
	Notice             interface{} //
	Policy             interface{} //
	Image              interface{} //
	ContactList        interface{} //
	BaiduTranApiKey    interface{} //
	BaiduTranSecretKey interface{} //
	CdnProxyUrl        interface{} //
	Phrase             interface{} //
	WxPayAppid         interface{} //
	WxPayMchId         interface{} //
	CAPIv3Key          interface{} //
	CSerialNo          interface{} //
	CNotifyUrl         interface{} //
	SpMchid            interface{} //
	SpAppid            interface{} //
	SubMchId           interface{} //
	APIv3Key           interface{} //
	SerialNo           interface{} //
	NotifyUrl          interface{} //
	PayType            interface{} //
	MpName             interface{} //
	WxAppId            interface{} //
	WxSecret           interface{} //
	IsWpNotice         interface{} //
}
