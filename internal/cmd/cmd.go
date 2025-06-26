package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gzdzh-cn/dzhcore"

	// "github.com/gzdzh-cn/dzhcore/contrib/drivers/mysql"
	// "github.com/gzdzh-cn/dzhcore/contrib/drivers/sqlite"
	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
	"github.com/gzdzh-cn/dzhcore/contrib/files/local"

	"dzhgo/addons"
	"dzhgo/internal"

	_ "dzhgo/internal/logic"

	_ "dzhgo/internal/packed"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {

			//初始化数据库
			// mysql.NewInit()
			// sqlite.NewInit()
			// 初始化本地文件上传驱动
			local.NewInit()
			// 初始化 oss上传驱动
			// oss.NewInit()
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
