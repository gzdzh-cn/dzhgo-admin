package config

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gzdzh-cn/dzhcore"
)

var Version = "v1.1.5"

func init() {
	dzhcore.SetVersions("base", Version)
}

// sConfig 配置
type sConfig struct {
	Jwt        *Jwt
	Middleware *Middleware
	Settiing   *g.Map
}

type Middleware struct {
	Authority *Authority
	Log       *Log
}

type Authority struct {
	Enable bool
}

type Log struct {
	Enable bool
}

type Token struct {
	Expire        uint `json:"expire"`
	RefreshExpire uint `json:"refreshExprire"`
}

type Jwt struct {
	Sso    bool   `json:"sso"`
	Secret string `json:"secret"`
	Token  *Token `json:"token"`
}

// NewConfig new config
func NewConfig() *sConfig {
	var (
		ctx g.Ctx
	)
	config := &sConfig{
		Jwt: &Jwt{
			Sso:    dzhcore.GetCfgWithDefault(ctx, "modules.base.jwt.sso", g.NewVar(false)).Bool(),
			Secret: dzhcore.GetCfgWithDefault(ctx, "modules.base.jwt.secret", g.NewVar(dzhcore.ProcessFlag)).String(),
			Token: &Token{
				Expire:        dzhcore.GetCfgWithDefault(ctx, "modules.base.jwt.token.expire", g.NewVar(2*3600)).Uint(),
				RefreshExpire: dzhcore.GetCfgWithDefault(ctx, "modules.base.jwt.token.refreshExpire", g.NewVar(15*24*3600)).Uint(),
			},
		},
		Middleware: &Middleware{
			Authority: &Authority{
				Enable: dzhcore.GetCfgWithDefault(ctx, "modules.base.middleware.authority.enable", g.NewVar(true)).Bool(),
			},
			Log: &Log{
				Enable: dzhcore.GetCfgWithDefault(ctx, "modules.base.middleware.log.enable", g.NewVar(true)).Bool(),
			},
		},
		Settiing: &g.Map{
			"cdnUrl": g.Cfg().MustGet(ctx, "modules.base.img.cdn_url"),
		},
	}

	return config
}

// Config config
var Config = NewConfig()

// 小程序
type MinConfig struct {
	RequestUrl string `json:"requestUrl"`
	Appid      string `json:"appid"`
	Secret     string `json:"secret"`
	GrantType  string `json:"grant_type"`
}

// 公众号信息
func GetWxCon() (data interface{}) {

	return &MinConfig{
		RequestUrl: "https://api.weixin.qq.com/sns/jscode2session",
		Appid:      "wx7a3c7f891ab07e34",
		Secret:     "51cdfc9e7570c5ac19581f7795fb27c0",
		GrantType:  "authorization_code",
	}
}

type AutoPhone struct {
	RequestUrl string `json:"requestUrl"`
	Code       string `json:"code"`
}

func GetAutoConfig() (data *AutoPhone) {
	return &AutoPhone{
		RequestUrl: "https://api.weixin.qq.com/wxa/business/getuserphonenumber",
	}
}

type AccessToken struct {
	RequestUrl string `json:"requestUrl"`
	GrantType  string `json:"grant_type"`
	Appid      string `json:"appid"`
	Secret     string `json:"secret"`
}

func GetAccessToken() (data *AccessToken) {
	return &AccessToken{
		RequestUrl: "https://api.weixin.qq.com/cgi-bin/token",
		Appid:      "wx7a3c7f891ab07e34",
		Secret:     "51cdfc9e7570c5ac19581f7795fb27c0",
		GrantType:  "client_credential",
	}
}
