package main

import (
	"flag"
	"fmt"

	"easy-chat/apps/im/ws/internal/config"
	"easy-chat/apps/im/ws/internal/handler"
	"easy-chat/apps/im/ws/internal/svc"
	"easy-chat/apps/im/ws/websocketx"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/im.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	if err := c.SetUp(); err != nil {
		panic(err)
	}
	ctx := svc.NewServiceContext(&c)
	auth := handler.NewJWTAuth(ctx)

	server := websocketx.NewServer(c.ListenOn, websocketx.WithServerAuthOption(auth))
	handler.RegisterRoutes(&server, ctx)

	fmt.Printf("im ws start %s", c.ListenOn)
	server.Start()
	defer server.Stop()
}
