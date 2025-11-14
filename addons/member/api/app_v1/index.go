package app_v1

import "github.com/gogf/gf/v2/frame/g"

type TokenRes struct {
	Expire        uint   `json:"expire"`
	Token         string `json:"token"`
	RefreshExpire uint   `json:"refreshExpire"`
	RefreshToken  string `json:"refreshToken"`
}

// 会员信息
type PersonReq struct {
	g.Meta        `path:"/person" method:"GET"`
	Authorization string `json:"Authorization" in:"header"`
}

// 账号注册
type AccountRegisterReq struct {
	g.Meta     `path:"/register" method:"POST"`
	MemberName string `json:"memberName" p:"memberName" v:"required|length:4,30#请输入账号|账号长度为:{min}到:{max}位"`
	PassWord   string `json:"passWord" p:"password" v:"required|length:6,30#请输入密码|密码长度不够"`
}

// 账号登录
type AccountLoginReq struct {
	g.Meta     `path:"/account" method:"POST"`
	MemberName string `json:"memberName" p:"memberName" v:"required|length:4,30#请输入账号|账号长度为:{min}到:{max}位"`
	PassWord   string `json:"passWord" p:"password" v:"required|length:6,30#请输入密码|密码长度不够"`
}

// 微信公众号登录
type MpLoginReq struct {
	g.Meta   `path:"/mp" method:"POST"`
	Code     string  `json:"code" p:"code" v:"required"`
	MemberId *string `json:"memberId" p:"memberId"`
	Notify   int     `json:"notify" p:"notify"`
}

// 微信小程序登录
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
