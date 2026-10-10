package main

import (
	"flag"
	"fmt"
	"sync"

	"easy-chat/apps/user/rpc/internal/config"
	"easy-chat/apps/user/rpc/internal/server"
	"easy-chat/apps/user/rpc/internal/svc"
	"easy-chat/apps/user/rpc/user"
	configserver "easy-chat/pkg/config_server"
	rpcserver "easy-chat/pkg/interceptor/rpcServer"

	"github.com/google/go-cmp/cmp"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/dev/rpc.yaml", "the config file")
var wg sync.WaitGroup

func main() {
	flag.Parse()

	var c config.Config
	// conf.MustLoad(*configFile, &c)
	sail := configserver.NewSail(&configserver.Config{
		ETCDEndpoints:  "192.168.1.5:3379", // 逗号分隔的ETCD地址，0.0.0.0:2379,0.0.0.0:12379,0.0.0.0:22379
		ProjectKey:     "98c6f2c2287f4c73cea3d40ae7ec3ff2",
		Namespace:      "user",
		Configs:        "user_rpc.yaml",
		ConfigFilePath: "", // 本地配置文件存放路径，空代表不存储本地配置文件
		LogLevel:       "DEBUG",
	})
	configServer := configserver.NewConfigServer(*configFile, sail)
	err := configServer.MustLoad(&c, func(bytes []byte) error {
		var nc config.Config
		err := configserver.LoadFromJsonBytes(bytes, &nc)
		if err != nil {
			// panic(err)
			return err
		}
		if cmp.Equal(c, nc) {
			return nil
		}
		// proc.WrapUp()
		// grpcSvr.GracefulStop()
		c = nc
		wait_run(&nc)
		return nil
	})
	if err != nil {
		panic(err)
	}
	wait_run(&c)

	wg.Wait()
}

func wait_run(c *config.Config) {
	wg.Add(1)
	go run(c)
}
func run(c *config.Config) {

	defer wg.Done()
	ctx := svc.NewServiceContext(*c)
	// 设置系统根用户
	if err := ctx.SetSystemRootToken(); err != nil {
		panic(err)
	}

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServiceServer(grpcServer, server.NewUserServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	// 注册UnaryInterceptor
	s.AddUnaryInterceptors(rpcserver.LogInterceptor)
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
