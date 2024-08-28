package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	ContentRPCConf    zrpc.RpcClientConf
	UsercenterRPCConf zrpc.RpcClientConf
	MiniConf          struct {
		EndPoint   string
		AccessKey  string
		SecretKey  string
		BucketName string
		Expires    int64
	}
}
