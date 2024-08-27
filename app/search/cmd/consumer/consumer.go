package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/me2seeks/echo-hub/app/search/cmd/consumer/internal/config"
	"github.com/me2seeks/echo-hub/app/search/cmd/consumer/internal/svc"
	"github.com/me2seeks/echo-hub/app/search/cmd/consumer/mqs"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
)

var configFile = flag.String("f", "etc/consumer.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	svcCtx := svc.NewServiceContext(c)
	ctx := context.Background()
	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()

	for _, mq := range mqs.Consumers(c, ctx, svcCtx) {
		serviceGroup.Add(mq)
	}

	//nolint:forbidigo
	fmt.Printf("Starting consumer service Topic:%s...\n", c.KqConsumerConf.Topic)
	serviceGroup.Start()
}
