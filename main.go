package main

import (
	_ "dzhgo/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"dzhgo/internal/cmd"
)

func main() {

	//gres.Dump()
	cmd.Main.Run(gctx.New())

}
