package config

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gzdzh-cn/dzhcore"
	"github.com/gzdzh-cn/dzhcore/utility/env"
)

var (
	ctx     = gctx.GetInitCtx()
	AppName = "dzhgo"
)

// sConfig 配置
type sConfig struct {
	Jwt        *Jwt
	Middleware *Middleware
	Setting    *g.Map
}

type Middleware struct {
	Cors      bool
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

func NewConfig() *sConfig {

	return &sConfig{
		Jwt: &Jwt{
			Sso:    env.GetCfgWithDefault(ctx, "modules.base.jwt.sso", g.NewVar(false)).Bool(),
			Secret: env.GetCfgWithDefault(ctx, "modules.base.jwt.secret", g.NewVar(dzhcore.ProcessFlag)).String(),
			Token: &Token{
				Expire:        env.GetCfgWithDefault(ctx, "modules.base.jwt.token.expire", g.NewVar(2*3600)).Uint(),
				RefreshExpire: env.GetCfgWithDefault(ctx, "modules.base.jwt.token.refreshExpire", g.NewVar(15*24*3600)).Uint(),
			},
		},
		Middleware: &Middleware{
			Cors: env.GetCfgWithDefault(ctx, "modules.base.middleware.cors", g.NewVar(false)).Bool(),
			Authority: &Authority{
				Enable: env.GetCfgWithDefault(ctx, "modules.base.middleware.authority.enable", g.NewVar(true)).Bool(),
			},
			Log: &Log{
				Enable: env.GetCfgWithDefault(ctx, "modules.base.middleware.log.enable", g.NewVar(true)).Bool(),
			},
		},
		Setting: &g.Map{
			"cdnUrl": g.Cfg().MustGet(ctx, "modules.base.img.cdn_url"),
		},
	}

}

var Cfg = NewConfig()
