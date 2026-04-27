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
	Id                 any         //
	CreateTime         *gtime.Time // 创建时间
	UpdateTime         *gtime.Time // 更新时间
	DeletedAt          *gtime.Time //
	SiteName           any         // 站点名称
	SiteDescribe       any         // 站点介绍
	DomainName         any         // 网站域名
	Copyright          any         // 版权所有
	Logo               any         // logo
	WxCode             any         // 二维码
	Company            any         // 公司名称
	Contact            any         // 联系人
	ContactWay         any         // 座机
	Mobile             any         // 手机
	Address            any         // 地址
	Keyword            any         // 关键词
	Description        any         // 描述
	Smtp               any         // smtp
	SmtpEmail          any         // 发送邮箱
	SmtpPass           any         // 邮箱授权码
	RemindEmail        any         // 接收邮箱
	IsRemindEmail      any         // 到期邮件开启 0关闭 1开启
	IsRemindSms        any         // 到期短信开启 0关闭 1开启
	AccessKeyId        any         // accessKeyId
	AccessKeySecret    any         // accessKeySecret
	SignName           any         // 签名
	TemplateCode       any         // 模板
	Endpoint           any         // endpoint
	RemindMobile       any         // 通知手机号码
	RemindDay          any         // 到期提醒提前天数
	FieldJson          any         // 自定义字段
	Notice             any         // 公告
	Policy             any         // 隐私政策
	Image              any         // 图片
	ContactList        any         // 客服列表
	BaiduTranApiKey    any         // 百度翻译apikey
	BaiduTranSecretKey any         // 百度翻译Secretkey
	CdnProxyUrl        any         // 图片代理地址
	Phrase             any         // 过滤词
	WxPayAppid         any         // 普通商户appid
	WxPayMchId         any         // 普通商户号
	CAPIv3Key          any         // 收款商户v3密钥
	CSerialNo          any         // 序列号
	CNotifyUrl         any         // 支付回调地址
	SpMchid            any         // 服务商商户号
	SpAppid            any         // 服务商appid
	SubMchId           any         // 特约商户
	APIv3Key           any         // 收款商户v3密钥
	SerialNo           any         // 序列号
	NotifyUrl          any         // 支付回调地址
	PayType            any         // 支付模式
	MpName             any         // 公众号名称
	WxAppId            any         // 公众号appId
	WxSecret           any         // 微信secret
	IsWpNotice         any         // 公众号通知
}
