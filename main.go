package main

import (
	_ "focus-single/internal/packed"

	_ "focus-single/internal/logic"

	"focus-single/internal/cmd"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.New())
}
