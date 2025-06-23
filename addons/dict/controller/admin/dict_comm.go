package admin

import (
	"context"
	v1 "dzhgo/addons/dict/api/v1"
	logic "dzhgo/addons/dict/logic/sys"

	"github.com/gzdzh-cn/dzhcore"
)

type DictCommController struct {
	*dzhcore.ControllerSimple
}

func init() {
	var dictCommController = &DictCommController{
		&dzhcore.ControllerSimple{
			Prefix: "/admin/dict/comm",
		},
	}
	// 注册路由
	dzhcore.AddControllerSimple(dictCommController)
}

// Data 方法 获得字典数据
func (c *DictCommController) Data(ctx context.Context, req *v1.DictInfoDataReq) (res *dzhcore.BaseRes, err error) {

	data, err := logic.NewsDictInfoService().Data(ctx, req.Types)
	res = dzhcore.Ok(data)
	return
}
