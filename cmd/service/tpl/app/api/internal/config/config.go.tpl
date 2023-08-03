package config

import (
	"github.com/5-say/go-tools/tools/db"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	DB          db.Config
	CorsOrigins []string
	JWTXRpc     zrpc.RpcClientConf
}
