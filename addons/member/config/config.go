package config

import (
	"dzhgo/addons/member/defineType"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	ctx = gctx.GetInitCtx()
)

// 公众号信息
//func GetWxCf() (data *defineType.WxConfig, err error) {
//
//	var setting *baseEntity.BaseSysSetting
//	err = baseDao.BaseSysSetting.Ctx(ctx).Where(g.Map{"id": 1}).Scan(&setting)
//	if err != nil {
//		g.Log().Error(ctx, err)
//		return nil, err
//	}
//
//	return &defineType.WxConfig{
//		Appid:     gconv.String(setting.WxAppId),
//		Secret:    gconv.String(setting.WxSecret),
//		GrantType: "authorization_code",
//	}, nil
//}

// 微信手机号授权
func GetAutoConfig() *defineType.AutoPhone {
	return &defineType.AutoPhone{
		RequestUrl: "https://api.weixin.qq.com/wxa/business/getuserphonenumber",
	}
}

//func GetAccessToken() *AccessToken {
//
//	conf, err := baseDao.BaseSysSetting.Ctx(ctx).Where("id", 1).One()
//	if err != nil {
//		g.Log().Error(ctx, err)
//	}
//	return &AccessToken{
//		RequestUrl: "https://api.weixin.qq.com/cgi-bin/token",
//		Appid:      gconv.String(conf["wxAppId"]),
//		Secret:     gconv.String(conf["wxSecret"]),
//		GrantType:  "client_credential",
//	}
//}
