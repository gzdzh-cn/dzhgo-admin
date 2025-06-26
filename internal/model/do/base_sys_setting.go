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
	CreateTime         *gtime.Time // 创建时间
	UpdateTime         *gtime.Time // 更新时间
	DeletedAt          *gtime.Time //
	SiteName           interface{} // 站点名称
	SiteDescribe       interface{} // 站点介绍
	DomainName         interface{} // 网站域名
	Copyright          interface{} // 版权所有
	Logo               interface{} // logo
	WxCode             interface{} // 二维码
	Company            interface{} // 公司名称
	Contact            interface{} // 联系人
	ContactWay         interface{} // 座机
	Mobile             interface{} // 手机
	Address            interface{} // 地址
	Keyword            interface{} // 关键词
	Description        interface{} // 描述
	Smtp               interface{} // smtp
	SmtpEmail          interface{} // 发送邮箱
	SmtpPass           interface{} // 邮箱授权码
	RemindEmail        interface{} // 接收邮箱
	IsRemindEmail      interface{} // 到期邮件开启 0关闭 1开启
	IsRemindSms        interface{} // 到期短信开启 0关闭 1开启
	AccessKeyId        interface{} // accessKeyId
	AccessKeySecret    interface{} // accessKeySecret
	SignName           interface{} // 签名
	TemplateCode       interface{} // 模板
	Endpoint           interface{} // endpoint
	RemindMobile       interface{} // 通知手机号码
	RemindDay          interface{} // 到期提醒提前天数
	FieldJson          interface{} // 自定义字段
	Notice             interface{} // 公告
	Policy             interface{} // 隐私政策
	Image              interface{} // 图片
	ContactList        interface{} // 客服列表
	BaiduTranApiKey    interface{} // 百度翻译apikey
	BaiduTranSecretKey interface{} // 百度翻译Secretkey
	CdnProxyUrl        interface{} // 图片代理地址
	Phrase             interface{} // 过滤词
	WxPayAppid         interface{} // 普通商户appid
	WxPayMchId         interface{} // 普通商户号
	CAPIv3Key          interface{} // 收款商户v3密钥
	CSerialNo          interface{} // 序列号
	CNotifyUrl         interface{} // 支付回调地址
	SpMchid            interface{} // 服务商商户号
	SpAppid            interface{} // 服务商appid
	SubMchId           interface{} // 特约商户
	APIv3Key           interface{} // 收款商户v3密钥
	SerialNo           interface{} // 序列号
	NotifyUrl          interface{} // 支付回调地址
	PayType            interface{} // 支付模式
	MpName             interface{} // 公众号名称
	WxAppId            interface{} // 公众号appId
	WxSecret           interface{} // 微信secret
	IsWpNotice         interface{} // 公众号通知
}
