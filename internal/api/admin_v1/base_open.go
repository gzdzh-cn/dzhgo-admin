package admin_v1

import "github.com/gogf/gf/v2/frame/g"

// login 接口请求
type BaseOpenLoginReq struct {
	g.Meta     `path:"/login" method:"POST"`
	Username   string `json:"username" p:"username" v:"required"`
	Password   string `json:"password" p:"password" v:"required"`
	CaptchaId  string `json:"captchaId" p:"captchaId" v:"required"`
	VerifyCode string `json:"verifyCode" p:"verifyCode" v:"required"`
}

// captcha 验证码接口
type BaseOpenCaptchaReq struct {
	g.Meta `path:"/captcha" method:"GET"`
	Height int `json:"height" in:"query" default:"40"`
	Width  int `json:"width"  in:"query" default:"150"`
}

// 插件版本
type VersionsReq struct {
	g.Meta `path:"/versions" method:"POST"`
	Addons string `json:"addons" p:"addons" default:"all"`
}

// 支付回调通知
type NotifyUrlReq struct {
	g.Meta `path:"/notifyUrl" method:"POST"`
}

// 站点配置
type GetSettingReq struct {
	g.Meta `path:"/getSetting" method:"POST"`
	Addons string `json:"addons" p:"addons" default:"all"`
}

// 服务器信息
type ServerInfoReq struct {
	g.Meta `path:"/serverInfo" method:"POST"`
}

// login 接口请求
type LoginReq struct {
	g.Meta    `path:"/login" method:"POST"`
	JsCode    string `json:"code" p:"code" v:"required"`
	NickName  string `json:"nickName" p:"nickName" v:"required"`
	AvatarUrl string `json:"avatarUrl" p:"avatarUrl" v:"required"`
	LevelName string `json:"levelName" p:"levelName" v:"levelName"`
}

type TokenRes struct {
	Expire        uint   `json:"expire"`
	Token         string `json:"token"`
	RefreshExpire uint   `json:"refreshExpire"`
	RefreshToken  string `json:"refreshToken"`
}

type MpLoginReq struct {
	g.Meta `path:"/mp" method:"POST"`
}

// 接口请求 微信登录
type MiniLoginReq struct {
	g.Meta `path:"/mini" method:"POST"`
	JsCode string `json:"code" p:"code" v:"required"`
}

// 微信手机授权登录
type AutoPhoneReq struct {
	g.Meta `path:"/autoPhone" method:"POST"`
	Code   string `json:"code" p:"code" v:"required"`
}

// 游客登录
type TouristLoginReq struct {
	g.Meta `path:"/tourist" method:"POST"`
}

// 验证游客次数
type VerifyCountReq struct {
	g.Meta `path:"/verifyCount" method:"POST"`
	Token  string `json:"token" p:"token" v:"required"`
}

// 账号登录
//type AccountLoginReq struct {
//	g.Meta   `path:"/account" method:"POST"`
//	UserName string `json:"userName" p:"username" v:"required"`
//	PassWord string `json:"passWord" p:"password" v:"required"`
//}

// 账号注册
type AccountRegisterReq struct {
	g.Meta     `path:"/accountReg" method:"POST"`
	UserName   string `json:"userName" p:"userName" v:"required"`
	PassWord   string `json:"passWord" p:"passWord" v:"required"`
	RePassWord string `json:"rePassWord" p:"rePassWord" v:"required"`
}

// 会员
//type GetMemberReq struct {
//	g.Meta `path:"/getMember" method:"POST"`
//}
