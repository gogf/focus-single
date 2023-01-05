package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"

	_ "focus-single/internal/packed"

	_ "focus-single/internal/logic"

	"focus-single/internal/cmd"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.New())
}
