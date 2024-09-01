package main

import (
	"flag"
	"fmt"

	"github.com/me2seeks/echo-hub/app/content/cmd/api/internal/config"
	"github.com/me2seeks/echo-hub/app/content/cmd/api/internal/handler"
	"github.com/me2seeks/echo-hub/app/content/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/content.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	//nolint:forbidigo
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
