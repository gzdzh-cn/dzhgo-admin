// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BaseSysSetting is the golang structure for table base_sys_setting.
type BaseSysSetting struct {
	Id                 string      `json:"id"                 orm:"id"                 ` //
	CreateTime         *gtime.Time `json:"createTime"         orm:"createTime"         ` //
	UpdateTime         *gtime.Time `json:"updateTime"         orm:"updateTime"         ` //
	DeletedAt          *gtime.Time `json:"deletedAt"          orm:"deleted_at"         ` //
	SiteName           string      `json:"siteName"           orm:"siteName"           ` //
	SiteDescribe       string      `json:"siteDescribe"       orm:"siteDescribe"       ` //
	DomainName         string      `json:"domainName"         orm:"domainName"         ` //
	Copyright          string      `json:"copyright"          orm:"copyright"          ` //
	Logo               string      `json:"logo"               orm:"logo"               ` //
	WxCode             string      `json:"wxCode"             orm:"wxCode"             ` //
	Company            string      `json:"company"            orm:"company"            ` //
	Contact            string      `json:"contact"            orm:"contact"            ` //
	ContactWay         string      `json:"contactWay"         orm:"contactWay"         ` //
	Mobile             string      `json:"mobile"             orm:"mobile"             ` //
	Address            string      `json:"address"            orm:"Address"            ` //
	Keyword            string      `json:"keyword"            orm:"keyword"            ` //
	Description        string      `json:"description"        orm:"description"        ` //
	Smtp               string      `json:"smtp"               orm:"smtp"               ` //
	SmtpEmail          string      `json:"smtpEmail"          orm:"smtpEmail"          ` //
	SmtpPass           string      `json:"smtpPass"           orm:"smtpPass"           ` //
	RemindEmail        string      `json:"remindEmail"        orm:"remindEmail"        ` //
	IsRemindEmail      int         `json:"isRemindEmail"      orm:"isRemindEmail"      ` //
	IsRemindSms        int         `json:"isRemindSms"        orm:"isRemindSms"        ` //
	AccessKeyId        string      `json:"accessKeyId"        orm:"accessKeyId"        ` //
	AccessKeySecret    string      `json:"accessKeySecret"    orm:"accessKeySecret"    ` //
	SignName           string      `json:"signName"           orm:"signName"           ` //
	TemplateCode       string      `json:"templateCode"       orm:"templateCode"       ` //
	Endpoint           string      `json:"endpoint"           orm:"endpoint"           ` //
	RemindMobile       string      `json:"remindMobile"       orm:"remindMobile"       ` //
	RemindDay          int         `json:"remindDay"          orm:"remindDay"          ` //
	FieldJson          string      `json:"fieldJson"          orm:"fieldJson"          ` //
	Notice             string      `json:"notice"             orm:"notice"             ` //
	Policy             string      `json:"policy"             orm:"policy"             ` //
	Image              string      `json:"image"              orm:"image"              ` //
	ContactList        string      `json:"contactList"        orm:"contactList"        ` //
	BaiduTranApiKey    string      `json:"baiduTranApiKey"    orm:"baiduTranApiKey"    ` //
	BaiduTranSecretKey string      `json:"baiduTranSecretKey" orm:"baiduTranSecretKey" ` //
	CdnProxyUrl        string      `json:"cdnProxyUrl"        orm:"cdn_proxy_url"      ` //
	Phrase             string      `json:"phrase"             orm:"phrase"             ` //
	WxPayAppid         string      `json:"wxPayAppid"         orm:"wxPayAppid"         ` //
	WxPayMchId         string      `json:"wxPayMchId"         orm:"wxPayMchId"         ` //
	CAPIv3Key          string      `json:"cAPIv3Key"          orm:"cAPIv3Key"          ` //
	CSerialNo          string      `json:"cSerialNo"          orm:"cSerialNo"          ` //
	CNotifyUrl         string      `json:"cNotifyUrl"         orm:"cNotifyUrl"         ` //
	SpMchid            string      `json:"spMchid"            orm:"spMchid"            ` //
	SpAppid            string      `json:"spAppid"            orm:"spAppid"            ` //
	SubMchId           string      `json:"subMchId"           orm:"subMchId"           ` //
	APIv3Key           string      `json:"aPIv3Key"           orm:"aPIv3Key"           ` //
	SerialNo           string      `json:"serialNo"           orm:"serialNo"           ` //
	NotifyUrl          string      `json:"notifyUrl"          orm:"notifyUrl"          ` //
	PayType            int         `json:"payType"            orm:"payType"            ` //
	MpName             string      `json:"mpName"             orm:"mpName"             ` //
	WxAppId            string      `json:"wxAppId"            orm:"wxAppId"            ` //
	WxSecret           string      `json:"wxSecret"           orm:"wxSecret"           ` //
	IsWpNotice         int         `json:"isWpNotice"         orm:"isWpNotice"         ` //
}
