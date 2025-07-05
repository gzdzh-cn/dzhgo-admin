package cmd

import (
	"context"
	"dzhgo/addons"
	"dzhgo/internal"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gzdzh-cn/dzhcore"

	_ "dzhgo/internal/logic"
	_ "dzhgo/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
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

			s.Run()

			return nil
		},
	}
)
