package model

import (
	"github.com/gzdzh-cn/dzhcore"
)

const TableNameBaseSysSetting = "base_sys_setting"

// BaseSysSetting mapped from table <base_sys_setting>
type BaseSysSetting struct {
	*dzhcore.Model
	SiteName     string  `gorm:"column:siteName;comment:站点名称" json:"siteName"`
	SiteDescribe string  `gorm:"column:siteDescribe;comment:站点介绍" json:"siteDescribe"`
	DomainName   *string `gorm:"column:domainName;comment:网站域名;type:varchar(50)" json:"domainName"`
	Copyright    *string `gorm:"column:copyright;comment:版权所有;type:varchar(255)" json:"copyright"`

	Logo        *string `gorm:"column:logo;comment:logo" json:"logo"`
	WxCode      *string `gorm:"column:wxCode;comment:二维码" json:"wxCode"`
	Company     *string `gorm:"column:company;comment:公司名称;type:varchar(50)" json:"company"`
	Contact     *string `gorm:"column:contact;comment:联系人;type:varchar(50)" json:"contact"`
	ContactWay  *string `gorm:"column:contactWay;comment:座机;type:varchar(50)" json:"contactWay"`
	Mobile      *string `gorm:"column:mobile;comment:手机;type:varchar(50)" json:"mobile"`
	Address     *string `gorm:"column:Address;comment:地址;type:varchar(50)" json:"address"`
	Keyword     *string `gorm:"column:keyword;comment:关键词;type:varchar(50)" json:"keyword"`
	Description *string `gorm:"column:description;comment:描述;type:varchar(50)" json:"description"`

	Smtp            *string `gorm:"column:smtp;comment:smtp;type:varchar(50)" json:"smtp"`
	SmtpEmail       *string `gorm:"column:smtpEmail;comment:发送邮箱;type:varchar(50)" json:"smtpEmail"`
	SmtpPass        *string `gorm:"column:smtpPass;comment:邮箱授权码;type:varchar(50)" json:"smtpPass"`
	RemindEmail     *string `gorm:"column:remindEmail;comment:接收邮箱;type:varchar(50)" json:"remindEmail"`
	IsRemindEmail   *int    `gorm:"column:isRemindEmail;comment:到期邮件开启 0关闭 1开启;default:0;type:int(11)" json:"isRemindEmail"`
	IsRemindSms     *int    `gorm:"column:isRemindSms;comment:到期短信开启 0关闭 1开启;default:0;type:int(11)" json:"isRemindSms"`
	AccessKeyId     *string `gorm:"column:accessKeyId;comment:accessKeyId;type:varchar(100)" json:"accessKeyId"`
	AccessKeySecret *string `gorm:"column:accessKeySecret;comment:accessKeySecret;type:varchar(100)" json:"accessKeySecret"`

	SignName     *string `gorm:"column:signName;comment:签名;type:varchar(50)" json:"signName"`
	TemplateCode *string `gorm:"column:templateCode;comment:模板;type:varchar(50)" json:"templateCode"`
	Endpoint     *string `gorm:"column:endpoint;comment:endpoint;type:varchar(50)" json:"endpoint"`
	RemindMobile *string `gorm:"column:remindMobile;comment:通知手机号码;type:varchar(50)" json:"remindMobile"`
	RemindDay    *string `gorm:"column:remindDay;comment:到期提醒提前天数;type:varchar(50)" json:"remindDay"`
	FieldJson    *string `gorm:"column:fieldJson;comment:自定义字段" json:"fieldJson"`

	Notice      *string `gorm:"column:notice;comment:公告" json:"notice"`
	Policy      *string `gorm:"column:policy;comment:隐私政策;" json:"policy"`
	Image       *string `gorm:"column:image;comment:图片" json:"image"`
	ContactList *string `gorm:"column:contactList;comment:客服列表" json:"contactList"`

	BaiduTranApiKey    *string `gorm:"column:baiduTranApiKey;comment:百度翻译apikey" json:"baiduTranApiKey"`
	BaiduTranSecretKey *string `gorm:"column:baiduTranSecretKey;comment:百度翻译Secretkey" json:"baiduTranSecretKey"`
	CdnProxyUrl        *string `gorm:"column:cdn_proxy_url;comment:图片代理地址" json:"cdn_proxy_url"`
	Phrase             *string `gorm:"column:phrase;comment:过滤词" json:"phrase"`
	WxPayAppid         *string `gorm:"column:wxPayAppid;comment:普通商户appid" json:"wxPayAppid"`
	WxPayMchId         *string `gorm:"column:wxPayMchId;comment:普通商户号" json:"wxPayMchId"`
	CAPIv3Key          *string `gorm:"column:cAPIv3Key;comment:收款商户v3密钥" json:"cAPIv3Key"`
	CSerialNo          *string `gorm:"column:cSerialNo;comment:序列号" json:"cSerialNo"`
	CNotifyUrl         *string `gorm:"column:cNotifyUrl;comment:支付回调地址" json:"cNotifyUrl"`
	SpMchid            *string `gorm:"column:spMchid;comment:服务商商户号" json:"spMchid"`
	SpAppid            *string `gorm:"column:spAppid;comment:服务商appid" json:"spAppid"`
	SubMchId           *string `gorm:"column:subMchId;comment:特约商户" json:"subMchId"`
	APIv3Key           *string `gorm:"column:aPIv3Key;comment:收款商户v3密钥" json:"aPIv3Key"`
	SerialNo           *string `gorm:"column:serialNo;comment:序列号" json:"serialNo"`
	NotifyUrl          *string `gorm:"column:notifyUrl;comment:支付回调地址" json:"notifyUrl"`
	PayType            *int    `gorm:"column:payType;comment:支付模式;not null;default:1;index" json:"payType"`

	MpName     string `gorm:"column:mpName;comment:公众号名称" json:"mpName"`
	WxAppID    string `gorm:"column:wxAppId;comment:公众号appId;" json:"wxAppId"`
	WxSecret   string `gorm:"column:wxSecret;comment:微信secret;" json:"wxSecret"`
	IsWpNotice int    `gorm:"column:isWpNotice;comment:公众号通知;" json:"isWpNotice"`
}

// TableName BaseSysSetting's table name
func (*BaseSysSetting) TableName() string {
	return TableNameBaseSysSetting
}

// GroupName BaseSysSetting's table group
func (*BaseSysSetting) GroupName() string {
	return "default"
}

// NewBaseSysSetting create a new BaseSysSetting
func NewBaseSysSetting() *BaseSysSetting {
	return &BaseSysSetting{
		Model: dzhcore.NewModel(),
	}
}

// init 创建表
func init() {
	// dzhcore.CreateTable(&BaseSysSetting{})
	dzhcore.AddModel(&BaseSysSetting{})
}
