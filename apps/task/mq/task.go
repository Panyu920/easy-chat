package main

import (
	"easy-chat/apps/task/mq/internal/config"
	"easy-chat/apps/task/mq/internal/handler"
	"easy-chat/apps/task/mq/internal/svc"
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/service"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/dev/mq.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	svcCtx := svc.NewServiceContext(&c)
	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()

	for _, mq := range handler.NewListen(svcCtx).Services() {
		serviceGroup.Add(mq)
	}

	fmt.Printf("kafka mq start : %s ...\n", c.ListenOn)
	serviceGroup.Start()
}
