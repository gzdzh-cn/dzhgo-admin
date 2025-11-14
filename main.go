package main

import (
	_ "dzhgo/internal/logic"

	"dzhgo/internal/cmd"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	//gres.Dump()
	ctx := gctx.New()
	cmd.Main.Run(ctx)
}
