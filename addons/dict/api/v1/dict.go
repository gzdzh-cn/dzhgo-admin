package v1

import "github.com/gogf/gf/v2/frame/g"

// 方法请求
type DictInfoDataReq struct {
	g.Meta `path:"/data" method:"POST"`
	Types  []string `json:"types"`
}
