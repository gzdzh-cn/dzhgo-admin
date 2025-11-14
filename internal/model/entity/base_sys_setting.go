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
	CreateTime         *gtime.Time `json:"createTime"         orm:"createTime"         ` // 创建时间
	UpdateTime         *gtime.Time `json:"updateTime"         orm:"updateTime"         ` // 更新时间
	DeletedAt          *gtime.Time `json:"deletedAt"          orm:"deleted_at"         ` //
	SiteName           string      `json:"siteName"           orm:"siteName"           ` // 站点名称
	SiteDescribe       string      `json:"siteDescribe"       orm:"siteDescribe"       ` // 站点介绍
	DomainName         string      `json:"domainName"         orm:"domainName"         ` // 网站域名
	Copyright          string      `json:"copyright"          orm:"copyright"          ` // 版权所有
	Logo               string      `json:"logo"               orm:"logo"               ` // logo
	WxCode             string      `json:"wxCode"             orm:"wxCode"             ` // 二维码
	Company            string      `json:"company"            orm:"company"            ` // 公司名称
	Contact            string      `json:"contact"            orm:"contact"            ` // 联系人
	ContactWay         string      `json:"contactWay"         orm:"contactWay"         ` // 座机
	Mobile             string      `json:"mobile"             orm:"mobile"             ` // 手机
	Address            string      `json:"address"            orm:"Address"            ` // 地址
	Keyword            string      `json:"keyword"            orm:"keyword"            ` // 关键词
	Description        string      `json:"description"        orm:"description"        ` // 描述
	Smtp               string      `json:"smtp"               orm:"smtp"               ` // smtp
	SmtpEmail          string      `json:"smtpEmail"          orm:"smtpEmail"          ` // 发送邮箱
	SmtpPass           string      `json:"smtpPass"           orm:"smtpPass"           ` // 邮箱授权码
	RemindEmail        string      `json:"remindEmail"        orm:"remindEmail"        ` // 接收邮箱
	IsRemindEmail      int         `json:"isRemindEmail"      orm:"isRemindEmail"      ` // 到期邮件开启 0关闭 1开启
	IsRemindSms        int         `json:"isRemindSms"        orm:"isRemindSms"        ` // 到期短信开启 0关闭 1开启
	AccessKeyId        string      `json:"accessKeyId"        orm:"accessKeyId"        ` // accessKeyId
	AccessKeySecret    string      `json:"accessKeySecret"    orm:"accessKeySecret"    ` // accessKeySecret
	SignName           string      `json:"signName"           orm:"signName"           ` // 签名
	TemplateCode       string      `json:"templateCode"       orm:"templateCode"       ` // 模板
	Endpoint           string      `json:"endpoint"           orm:"endpoint"           ` // endpoint
	RemindMobile       string      `json:"remindMobile"       orm:"remindMobile"       ` // 通知手机号码
	RemindDay          int         `json:"remindDay"          orm:"remindDay"          ` // 到期提醒提前天数
	FieldJson          string      `json:"fieldJson"          orm:"fieldJson"          ` // 自定义字段
	Notice             string      `json:"notice"             orm:"notice"             ` // 公告
	Policy             string      `json:"policy"             orm:"policy"             ` // 隐私政策
	Image              string      `json:"image"              orm:"image"              ` // 图片
	ContactList        string      `json:"contactList"        orm:"contactList"        ` // 客服列表
	BaiduTranApiKey    string      `json:"baiduTranApiKey"    orm:"baiduTranApiKey"    ` // 百度翻译apikey
	BaiduTranSecretKey string      `json:"baiduTranSecretKey" orm:"baiduTranSecretKey" ` // 百度翻译Secretkey
	CdnProxyUrl        string      `json:"cdnProxyUrl"        orm:"cdn_proxy_url"      ` // 图片代理地址
	Phrase             string      `json:"phrase"             orm:"phrase"             ` // 过滤词
	WxPayAppid         string      `json:"wxPayAppid"         orm:"wxPayAppid"         ` // 普通商户appid
	WxPayMchId         string      `json:"wxPayMchId"         orm:"wxPayMchId"         ` // 普通商户号
	CAPIv3Key          string      `json:"cAPIv3Key"          orm:"cAPIv3Key"          ` // 收款商户v3密钥
	CSerialNo          string      `json:"cSerialNo"          orm:"cSerialNo"          ` // 序列号
	CNotifyUrl         string      `json:"cNotifyUrl"         orm:"cNotifyUrl"         ` // 支付回调地址
	SpMchid            string      `json:"spMchid"            orm:"spMchid"            ` // 服务商商户号
	SpAppid            string      `json:"spAppid"            orm:"spAppid"            ` // 服务商appid
	SubMchId           string      `json:"subMchId"           orm:"subMchId"           ` // 特约商户
	APIv3Key           string      `json:"aPIv3Key"           orm:"aPIv3Key"           ` // 收款商户v3密钥
	SerialNo           string      `json:"serialNo"           orm:"serialNo"           ` // 序列号
	NotifyUrl          string      `json:"notifyUrl"          orm:"notifyUrl"          ` // 支付回调地址
	PayType            int64       `json:"payType"            orm:"payType"            ` // 支付模式
	MpName             string      `json:"mpName"             orm:"mpName"             ` // 公众号名称
	WxAppId            string      `json:"wxAppId"            orm:"wxAppId"            ` // 公众号appId
	WxSecret           string      `json:"wxSecret"           orm:"wxSecret"           ` // 微信secret
	IsWpNotice         int64       `json:"isWpNotice"         orm:"isWpNotice"         ` // 公众号通知
}
