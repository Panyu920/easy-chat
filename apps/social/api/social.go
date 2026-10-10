// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package main

import (
	"flag"
	"fmt"
	"sync"

	"easy-chat/apps/social/api/internal/config"
	"easy-chat/apps/social/api/internal/handler"
	"easy-chat/apps/social/api/internal/svc"
	"easy-chat/pkg/resultx"

	configserver "easy-chat/pkg/config_server"

	"github.com/zeromicro/go-zero/core/proc"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/social.yaml", "the config file")
var wg sync.WaitGroup

func main() {
	flag.Parse()

	var c config.Config
	sail := configserver.NewSail(&configserver.Config{
		ETCDEndpoints:  "192.168.1.5:3379", // 逗号分隔的ETCD地址，0.0.0.0:2379,0.0.0.0:12379,0.0.0.0:22379
		ProjectKey:     "98c6f2c2287f4c73cea3d40ae7ec3ff2",
		Namespace:      "social",
		Configs:        "social_api.yaml",
		ConfigFilePath: "", // 本地配置文件存放路径，空代表不存储本地配置文件
		LogLevel:       "DEBUG",
	})
	configServer := configserver.NewConfigServer(*configFile, sail)
	err := configServer.MustLoad(&c, func(bytes []byte) error {
		err := configserver.LoadFromJsonBytes(bytes, &c)
		if err != nil {
			// panic(err)
			return err
		}
		proc.WrapUp()
		wait_run(&c)
		return nil
	})
	if err != nil {
		panic(err)
	}
	wait_run(&c)
	wg.Wait()

}
func run(c *config.Config) {
	server := rest.MustNewServer(c.RestConf)
	defer wg.Done()
	defer server.Stop()

	ctx := svc.NewServiceContext(*c)
	handler.RegisterHandlers(server, ctx)

	httpx.SetOkHandler(resultx.OkHandler)
	httpx.SetErrorHandlerCtx(resultx.ErrorHandler(c.Name))

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func wait_run(c *config.Config) {
	wg.Add(1)
	go run(c)
}
