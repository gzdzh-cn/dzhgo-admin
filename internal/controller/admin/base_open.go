package admin

import (
	"context"
	logic "dzhgo/internal/logic/sys"
	"dzhgo/internal/service"

	v1 "dzhgo/internal/api/admin_v1"

	"github.com/gzdzh-cn/dzhcore"

	"github.com/gogf/gf/v2/frame/g"
)

type BaseOpenController struct {
	*dzhcore.ControllerSimple
	baseSysLoginService service.IBaseSysLoginService
	baseOpenService     service.IBaseOpenService
}

func init() {
	var open = &BaseOpenController{
		ControllerSimple:    &dzhcore.ControllerSimple{Prefix: "/admin/base/open"},
		baseSysLoginService: logic.NewsBaseSysLoginService(),
		baseOpenService:     logic.NewsBaseOpenService(),
	}
	// 注册路由
	dzhcore.AddControllerSimple(open)
}

// 验证码接口
func (c *BaseOpenController) BaseOpenCaptcha(ctx context.Context, req *v1.BaseOpenCaptchaReq) (res *dzhcore.BaseRes, err error) {
	data, err := c.baseSysLoginService.Captcha(req)
	res = dzhcore.Ok(data)
	return
}

// eps 接口
func (c *BaseOpenController) Eps(ctx context.Context, req *v1.BaseOpenEpsReq) (res *dzhcore.BaseRes, err error) {
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
func (c *BaseOpenController) Login(ctx context.Context, req *v1.BaseOpenLoginReq) (res *dzhcore.BaseRes, err error) {
	data, err := c.baseSysLoginService.Login(ctx, req)
	if err != nil {
		return
	}
	res = dzhcore.Ok(data)
	return
}

// 站点配置
func (c *BaseOpenController) GetSetting(ctx context.Context, req *v1.GetSettingReq) (res *dzhcore.BaseRes, err error) {
	data, err := service.BaseOpenService().GetSetting(ctx, req)
	if err != nil {
		return
	}
	res = dzhcore.Ok(data)
	return
}

// RefreshToken 刷新token
func (c *BaseOpenController) RefreshToken(ctx context.Context, req *v1.RefreshTokenReq) (res *dzhcore.BaseRes, err error) {
	data, err := c.baseSysLoginService.RefreshToken(ctx, req.RefreshToken)
	if err != nil {
		return
	}
	res = dzhcore.Ok(data)
	return
}

// 版本
func (c *BaseOpenController) Versions(ctx context.Context, req *v1.VersionsReq) (res *dzhcore.BaseRes, err error) {
	data, err := service.BaseOpenService().Versions(ctx, req)
	if err != nil {
		return
	}
	res = dzhcore.Ok(data)
	return
}

// 服务器信息
func (c *BaseOpenController) ServerInfo(ctx context.Context, req *v1.ServerInfoReq) (res *dzhcore.BaseRes, err error) {
	data, err := service.BaseOpenService().ServerInfo(ctx)
	if err != nil {
		return
	}
	res = dzhcore.Ok(data)
	return
}
