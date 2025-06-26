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

// 账号登录
type AccountLoginReq struct {
	g.Meta   `path:"/account" method:"POST"`
	UserName string `json:"userName" p:"username" v:"required"`
	PassWord string `json:"passWord" p:"password" v:"required"`
}

// 微信公众号登录
type MpLoginReq struct {
	g.Meta `path:"/mp" method:"POST"`
	Code   string  `json:"code" p:"code" v:"required"`
	UserId *string `json:"userId" p:"userId"`
	Notify int     `json:"notify" p:"notify"`
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
