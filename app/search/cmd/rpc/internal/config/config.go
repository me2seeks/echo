package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	EsConf struct {
		Address            []string
		Username           string
		Password           string
		CertFile           string
		KeyFile            string
		InsecureSkipVerify bool
	}
}
