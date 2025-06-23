package boot

import (
	"flag"
	"github.com/gogf/gf/contrib/config/apollo/v2"
	"github.com/gogf/gf/util/guid"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
)

func init() {
	var (
		ctx     = gctx.GetInitCtx()
		appId   = "SampleApp"
		cluster = "default"
		ip      = "http://localhost:8080"
	)

	// 加载配置
	cfg := g.Cfg(guid.S())
	configFile := flag.String("f", "/manifest/config/config-dev.yaml", "The config file path")
	flag.Parse()

	cfg.GetAdapter().(*gcfg.AdapterFile).SetFileName(*configFile)

	adapter, err := apollo.New(ctx, apollo.Config{
		AppID:   appId,
		IP:      ip,
		Cluster: cluster,
	})
	if err != nil {
		g.Log().Fatalf(ctx, `%+v`, err)
	}
	g.Cfg().SetAdapter(adapter)

}
