package admin

import (
	"context"
	logic "dzhgo/internal/logic/sys"
	"dzhgo/internal/service"

	v1 "dzhgo/internal/api/admin_v1"
	"github.com/gzdzh-cn/dzhcore"

	"github.com/gogf/gf/v2/frame/g"
)

type BaseOpen struct {
	*dzhcore.ControllerSimple
	baseSysLoginService service.IBaseSysLoginService
	baseOpenService     service.IBaseOpenService
}

func init() {
	var open = &BaseOpen{
		ControllerSimple:    &dzhcore.ControllerSimple{Prefix: "/admin/base/open"},
		baseSysLoginService: logic.NewsBaseSysLoginService(),
		baseOpenService:     logic.NewsBaseOpenService(),
	}
	// 注册路由
	dzhcore.RegisterControllerSimple(open)
}

// 验证码接口
func (c *BaseOpen) BaseOpenCaptcha(ctx context.Context, req *v1.BaseOpenCaptchaReq) (res *dzhcore.BaseRes, err error) {
	data, err := c.baseSysLoginService.Captcha(req)
	res = dzhcore.Ok(data)
	return
}

// eps 接口
func (c *BaseOpen) Eps(ctx context.Context, req *v1.BaseOpenEpsReq) (res *dzhcore.BaseRes, err error) {
	if !dzhcore.Config.Eps {
		g.Log().Error(ctx, "eps is not open")
		res = dzhcore.Ok(nil)
		return
	}
	data, err := c.baseOpenService.AdminEPS(ctx)
	if err != nil {
		g.Log().Error(ctx, "eps error", err)
		return dzhcore.Fail(err.Error()), err
	}
	res = dzhcore.Ok(data)
	return
}

// login 接口
func (c *BaseOpen) Login(ctx context.Context, req *v1.BaseOpenLoginReq) (res *dzhcore.BaseRes, err error) {
	data, err := c.baseSysLoginService.Login(ctx, req)
	if err != nil {
		return
	}
	res = dzhcore.Ok(data)
	return
}

// 站点配置
func (c *BaseOpen) GetSetting(ctx context.Context, req *v1.GetSettingReq) (res *dzhcore.BaseRes, err error) {
	data, err := service.BaseOpenService().GetSetting(ctx, req)
	if err != nil {
		return
	}
	res = dzhcore.Ok(data)
	return
}

// RefreshToken 刷新token
func (c *BaseOpen) RefreshToken(ctx context.Context, req *v1.RefreshTokenReq) (res *dzhcore.BaseRes, err error) {
	data, err := c.baseSysLoginService.RefreshToken(ctx, req.RefreshToken)
	if err != nil {
		return
	}
	res = dzhcore.Ok(data)
	return
}
