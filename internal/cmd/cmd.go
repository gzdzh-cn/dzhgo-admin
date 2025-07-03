package cmd

import (
	"context"
	"dzhgo/addons"
	"dzhgo/internal"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gzdzh-cn/dzhcore"

	// _ "github.com/gzdzh-cn/dzhcore/contrib/drivers/mysql"
	_ "dzhgo/internal/logic"
	_ "dzhgo/internal/packed"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	_ "github.com/gzdzh-cn/dzhcore/contrib/drivers/sqlite"
	_ "github.com/gzdzh-cn/dzhcore/contrib/files/local"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {

			//dzhcore 核心加载
			dzhcore.NewInit()
			// 初始化internale数据
			internal.NewInit()
			// 初始化addons
			addons.NewInit()

			if dzhcore.IsRedisMode {
				go dzhcore.ListenFunc(ctx)
			}

			s := g.Server()
			s.AddStaticPath("/dzhimg/public", "/public")
			s.Run()

			return nil
		},
	}
)
