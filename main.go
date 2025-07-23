package main

import (
	_ "zeal_be/internal/packed"

	_ "zeal_be/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"zeal_be/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
