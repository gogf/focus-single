package main

import (
	_ "focus-single/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"focus-single/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
