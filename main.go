package main

import (
	"github.com/5-say/goctl-ts/command/ts"
	"github.com/zeromicro/go-zero/tools/goctl/internal/cobrax"
)

var (
	tsCmd = cobrax.NewCommand("ts", cobrax.WithRunE(ts.Command))
)

func main() {
	var (
		tsCmdFlags = tsCmd.Flags()
	)
	tsCmdFlags.StringVar(&ts.VarStringDir, "dir")
	tsCmdFlags.StringVar(&ts.VarStringAPI, "api")
	tsCmdFlags.StringVar(&ts.VarStringCaller, "caller")
	tsCmdFlags.BoolVar(&ts.VarBoolUnWrap, "unwrap")
}
