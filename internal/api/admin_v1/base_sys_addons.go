package admin_v1

import "github.com/gogf/gf/v2/frame/g"

// 安装卸载插件
type InstallUpdateStatusReq struct {
	g.Meta `path:"/installUpdateStatus" method:"POST"`
	Id     string `json:"id"`
	Active bool   `json:"active"`
}

// 上下架插件
type LineUpdateStatusReq struct {
	g.Meta `path:"/lineUpdateStatus" method:"POST"`
	Id     string `json:"id"`
	Active bool   `json:"active"`
}
