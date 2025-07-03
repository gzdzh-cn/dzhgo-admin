package app

import (
	"context"
	v1 "dzhgo/internal/api/admin_v1"
	"dzhgo/internal/config"
	"dzhgo/internal/service"

	"github.com/gzdzh-cn/dzhcore"
	"github.com/gzdzh-cn/dzhcore/coreconfig"

	"regexp"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type BaseOpenController struct {
	*dzhcore.ControllerSimple
}

func init() {
	var baseOpenController = &BaseOpenController{
		&dzhcore.ControllerSimple{
			Prefix: "/app/base/open",
		},
	}
	// 注册路由
	dzhcore.RegisterControllerSimple(baseOpenController)
}

// eps 接口
func (c *BaseOpenController) Eps(ctx context.Context, req *v1.BaseCommControllerEpsReq) (res *dzhcore.BaseRes, err error) {
	if !coreconfig.Config.Core.Eps {
		g.Log().Error(ctx, "eps is not open")
		res = dzhcore.Ok(nil)
		return
	}
	data, err := service.BaseOpenService().AppEPS(ctx)
	if err != nil {
		g.Log().Error(ctx, "eps error", err)
		return dzhcore.Fail(err.Error()), err
	}
	res = dzhcore.Ok(data)
	return
}

// 上传模式
func (c *BaseCommController) UploadMode(ctx context.Context, req *v1.BaseCommUploadModeReq) (res *dzhcore.BaseRes, err error) {
	data, err := dzhcore.File().GetMode()
	res = dzhcore.Ok(data)
	return
}

// 上传
func (c *BaseCommController) Upload(ctx context.Context, req *v1.BaseCommUploadReq) (res *dzhcore.BaseRes, err error) {
	data, err := dzhcore.File().Upload(ctx)
	// 如果是oss模式，cdn域名存在就替换
	mode, _ := dzhcore.File().GetMode()
	if gconv.Map(mode)["type"] == "oss" {
		setting := gconv.Map(config.Cfg.Setting)
		if setting["cdnUrl"] != "" {
			re := regexp.MustCompile(`^(https://[^/]+)`)
			resultURL := re.ReplaceAllString(data, gconv.String(setting["cdnUrl"]))
			data = resultURL
		}
	}
	res = dzhcore.Ok(data)
	return
}

// 小程序登录
//func (c *BaseCommController) Login(ctx g.Ctx, req *v1.LoginReq) (res *dzhcore.BaseRes, err error) {
//
//	data, err := service.BaseSysMemberLoginService().Login(ctx, req)
//	if err != nil {
//		return
//	}
//	res = dzhcore.Ok(data)
//	return
//}

// 微信公众号登录
//func (c *BaseCommController) MpLoginReq(ctx g.Ctx, req *v1.MpLoginReq) (res *dzhcore.BaseRes, err error) {
//
//	data, err := service.BaseSysMemberLoginService().MpLoginReq(ctx, req)
//	if err != nil {
//		return
//	}
//	res = dzhcore.Ok(data)
//	return
//}

// 小程序登录
//func (c *BaseCommController) MiniLogin(ctx g.Ctx, req *v1.MiniLoginReq) (res *dzhcore.BaseRes, err error) {
//
//	data, err := service.BaseSysMemberLoginService().MiniLogin(ctx, req)
//	if err != nil {
//		return
//	}
//	res = dzhcore.Ok(data)
//	return
//}

// 手机授权登录
//func (c *BaseCommController) AutoPhone(ctx g.Ctx, req *v1.AutoPhoneReq) (res *dzhcore.BaseRes, err error) {
//
//	data, err := service.BaseSysMemberLoginService().AutoPhone(ctx, req)
//	if err != nil {
//		return
//	}
//	res = dzhcore.Ok(data)
//	return
//}

// 验证游客次数
//func (c *BaseCommController) VerifyCount(ctx g.Ctx, req *v1.VerifyCountReq) (res *dzhcore.BaseRes, err error) {
//
//	data, err := service.BaseSysMemberLoginService().VerifyCount(ctx, req)
//	if err != nil {
//		return
//	}
//	res = dzhcore.Ok(data)
//	return
//}

// 账号登录
//func (c *BaseCommController) AccountLogin(ctx g.Ctx, req *v1.AccountLoginReq) (res *dzhcore.BaseRes, err error) {
//
//	data, err := service.BaseSysMemberLoginService().AccountLogin(ctx, req)
//	if err != nil {
//		return
//	}
//	res = dzhcore.Ok(data)
//	return
//}

// 获取全部版本
func (c *BaseOpenController) Versions(ctx context.Context, req *v1.VersionsReq) (res *dzhcore.BaseRes, err error) {
	data, err := service.BaseOpenService().Versions(ctx, req)
	if err != nil {
		return
	}
	res = dzhcore.Ok(data)
	return
}

// 账号注册
//func (c *BaseCommController) AccountRegister(ctx g.Ctx, req *v1.AccountRegisterReq) (res *dzhcore.BaseRes, err error) {
//
//	data, err := service.BaseSysMemberLoginService().AccountRegister(ctx, req)
//	if err != nil {
//		return
//	}
//	res = dzhcore.Ok(data)
//	return
//}

//func (c *BaseOpenController) NotifyUrl(ctx context.Context, req *v1.methodName) (res *dzhcore.BaseRes, err error) {
//	data, err := service.BaseOpenService().funcName(ctx, req)
//	if err != nil {
//		return
//	}
//	res = dzhcore.Ok(data)
//	return
//}
