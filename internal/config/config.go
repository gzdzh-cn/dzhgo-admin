package config

import (
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gzdzh-cn/dzhcore/coreconfig"
	"github.com/gzdzh-cn/dzhcore/defineStruct"
)

var (
	ctx     = gctx.GetInitCtx()
	AppName = "dzhgo"
)

// sConfig 配置
type sConfig struct {
	// Jwt        *Jwt
	// Middleware *Middleware
	// Setting    *g.Map
	*defineStruct.Config
}

// type Middleware struct {
// 	Cors      bool
// 	Authority *Authority
// 	Log       *Log
// }

// type Authority struct {
// 	Enable bool
// }

// type Log struct {
// 	Enable bool
// }

// type Token struct {
// 	Expire        uint `json:"expire"`
// 	RefreshExpire uint `json:"refreshExpire"`
// }

// type Jwt struct {
// 	Sso    bool   `json:"sso"`
// 	Secret string `json:"secret"`
// 	Token  *Token `json:"token"`
// }

func NewConfig() *sConfig {

	return &sConfig{
		coreconfig.Config,
	}

}

var Cfg = NewConfig()
