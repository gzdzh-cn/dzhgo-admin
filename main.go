package main

import (
	_ "dzhgo/internal/logic"

	"dzhgo/internal/cmd"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {

	//gres.Dump()
	cmd.Main.Run(gctx.New())

}
