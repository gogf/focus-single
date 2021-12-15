package main

import (
	_ "focus-single/internal/packed"

	"focus-single/internal/cmd"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	var (
		ctx = gctx.New()
	)
	if err := cmd.Main.Run(ctx); err != nil {
		g.Log().Fatal(ctx, err)
	}
}
