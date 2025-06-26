package config

import (
	"context"
	baseDao "dzhgo/internal/dao"
	"dzhgo/internal/defineType"
	baseEntity "dzhgo/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// sBaseConfig 配置
type sBaseConfig struct {
	WxConfig *defineType.WxConfig
}

// 公众号信息
func GetWxCf(ctx context.Context) (data *defineType.WxConfig) {

	var setting *baseEntity.BaseSysSetting
	err := baseDao.BaseSysSetting.Ctx(ctx).Where(g.Map{"id": 1}).Scan(&setting)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil
	}

	return &defineType.WxConfig{
		Appid:     gconv.String(setting.WxAppId),
		Secret:    gconv.String(setting.WxSecret),
		GrantType: "authorization_code",
	}
}

func GetAutoConfig() (data *defineType.AutoPhone) {
	return &defineType.AutoPhone{
		RequestUrl: "https://api.weixin.qq.com/wxa/business/getuserphonenumber",
	}
}

func GetAccessToken() (data *defineType.AccessToken) {
	return &defineType.AccessToken{
		RequestUrl: "https://api.weixin.qq.com/cgi-bin/token",
		Appid:      "wx7a3c7f891ab07e34",
		Secret:     "51cdfc9e7570c5ac19581f7795fb27c0",
		GrantType:  "client_credential",
	}
}

func NewBaseConfig() *sBaseConfig {

	return &sBaseConfig{
		WxConfig: GetWxCf(ctx),
	}

}

//var BaseConfig = NewBaseConfig()
