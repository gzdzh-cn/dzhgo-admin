package app_v1

import "github.com/gogf/gf/v2/frame/g"

// sse 流式推送
type NoticeSseReq struct {
	g.Meta `path:"/noticeSse" method:"GET"`
}

// 插件版本
type VersionsReq struct {
	g.Meta `path:"/versions" method:"POST"`
	Addons string `json:"addons" p:"addons" default:"all"`
}
