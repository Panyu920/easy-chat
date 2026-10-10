// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"syscall"

	"easy-chat/apps/user/api/internal/config"
	"easy-chat/apps/user/api/internal/handler"
	"easy-chat/apps/user/api/internal/svc"
	"easy-chat/pkg/resultx"

	"github.com/google/go-cmp/cmp"

	configserver "easy-chat/pkg/config_server"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")
var restartChan = make(chan struct{})
var server *rest.Server
var wg sync.WaitGroup

func main() {
	flag.Parse()

	var c config.Config
	// conf.MustLoad(*configFile, &c)
	sail := configserver.NewSail(&configserver.Config{
		ETCDEndpoints:  "192.168.1.5:3379", // 逗号分隔的ETCD地址，0.0.0.0:2379,0.0.0.0:12379,0.0.0.0:22379
		ProjectKey:     "98c6f2c2287f4c73cea3d40ae7ec3ff2",
		Namespace:      "user",
		Configs:        "user_api.yaml",
		ConfigFilePath: "apps/user/api/etc/conf", // 本地配置文件存放路径，空代表不存储本地配置文件
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
		// fmt.Println("durantion :", c.)
		// proc.WrapUp()
		// proc.Shutdown()
		// proc.SetTimeToForceQuit(10 * time.Millisecond)

		sendSIGINT()
		c = nc
		<-restartChan
		wait_run(&nc)

		wg.Wait()
		return nil
	})
	if err != nil {
		panic(err)
	}
	wait_run(&c)

	wg.Wait()
	select {}
}
func run(c *config.Config) {
	server = rest.MustNewServer(c.RestConf)
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
	go func() {
		run(c)
		defer wg.Done()
		restartChan <- struct{}{}
	}()
}
func sendSIGINT() error {
	pid := os.Getpid()
	proc, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	return proc.Signal(syscall.SIGUSR2)
}
