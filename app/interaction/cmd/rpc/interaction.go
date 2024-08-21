package main

import (
	"flag"
	"fmt"

	"github.com/me2seeks/echo-hub/app/interaction/cmd/rpc/internal/config"
	"github.com/me2seeks/echo-hub/app/interaction/cmd/rpc/internal/server"
	"github.com/me2seeks/echo-hub/app/interaction/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/interaction/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/interaction.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterInteractionServer(grpcServer, server.NewInteractionServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	//nolint:forbidigo
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
