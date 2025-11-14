package app_v1

import "github.com/gogf/gf/v2/frame/g"

type BaseCommControllerEpsReq struct {
	g.Meta `path:"/eps" method:"GET"`
}

type BaseCommUploadModeReq struct {
	g.Meta        `path:"/uploadMode" method:"GET"`
	Authorization string `json:"Authorization" in:"header"`
}

type BaseCommUploadReq struct {
	g.Meta        `path:"/upload" method:"POST"`
	Authorization string `json:"Authorization" in:"header"`
}
