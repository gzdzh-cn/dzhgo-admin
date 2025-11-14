package admin_v1

import "github.com/gogf/gf/v2/frame/g"

// 接口请求参数
type BaseCommPersonReq struct {
	g.Meta        `path:"/person" method:"GET"`
	Authorization string `json:"Authorization" in:"header"`
}

// BaseCommPermmenuReq 接口请求参数
type BaseCommPermmenuReq struct {
	g.Meta        `path:"/permmenu" method:"GET"`
	Authorization string `json:"Authorization" in:"header"`
}

type BaseCommLogoutReq struct {
	g.Meta        `path:"/logout" method:"POST"`
	Authorization string `json:"Authorization" in:"header"`
}

type BaseCommUploadModeReq struct {
	g.Meta        `path:"/uploadMode" method:"GET"`
	Authorization string `json:"Authorization" in:"header"`
}

type BaseCommUploadReq struct {
	g.Meta        `path:"/upload" method:"POST"`
	Authorization string `json:"Authorization" in:"header"`
}

type PersonUpdateReq struct {
	g.Meta        `path:"/personUpdate" method:"POST"`
	Authorization string `json:"Authorization" in:"header"`
}

// eps 接口请求
type BaseOpenEpsReq struct {
	g.Meta `path:"/eps" method:"GET"`
}

// 刷新token请求
type RefreshTokenReq struct {
	g.Meta       `path:"/refreshToken" method:"GET"`
	RefreshToken string `json:"refreshToken" v:"required#refreshToken不能为空"`
}

// 接口请求参数
type OrderReq struct {
	g.Meta        `path:"/order" method:"GET"`
	Authorization string `json:"Authorization" in:"header"`
}

// 获取保留天数
type GetKeepReq struct {
	g.Meta `method:"GET" path:"/getKeep" summary:"获取保留天数" tags:"系统日志"`
}

// 清空日志
type ClearReq struct {
	g.Meta `method:"POST" path:"/clear" summary:"清空日志" tags:"系统日志"`
}

// 请求参数
type BaseSysParamHtmlReq struct {
	g.Meta `path:"/html" method:"GET"`
	Key    string `v:"required#请输入key"`
}

type UserMoveReq struct {
	g.Meta        `path:"/move" method:"GET"`
	Authorization string `json:"Authorization" in:"header"`
}

// eps 接口请求
type BaseCommControllerEpsReq struct {
	g.Meta `path:"/eps" method:"GET"`
}

// 设置保留天数
type SetKeepReq struct {
	g.Meta `method:"POST" path:"/setKeep" summary:"设置保留天数" tags:"系统日志"`
	Value  int `json:"value" v:"required#请输入保留天数"`
}

type NoticeSseReq struct {
	g.Meta `path:"/noticeSse" method:"GET"`
}

type NoticeSseCheckedReq struct {
	g.Meta       `path:"/noticeSseChecked" method:"POST"`
	ConnectionID string `json:"connectionID" v:"required#connectionID不能为空"`
}

// 一键已阅
type NoticeReadAllReq struct {
	g.Meta `path:"/noticeReadAll" method:"POST"`
}
