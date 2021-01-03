/**
 *    ____          __
 *   / __/__  ___ _/ /_____
 *  _\ \/ _ \/ _ `/  '_/ -_)
 * /___/_//_/\_,_/_/\_\\__/
 *
 * generate by http://patorjk.com/software/taag/#p=display&f=Small%20Slant&t=Snake
 */
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"

	"github.com/1024casts/snake/app/api"
	rpc "github.com/1024casts/snake/internal/server"
	"github.com/1024casts/snake/internal/service"
	"github.com/1024casts/snake/pkg/conf"
	"github.com/1024casts/snake/pkg/snake"
	routers "github.com/1024casts/snake/router"
)

var (
	cfg = pflag.StringP("config", "c", "", "snake config file path.")
)

// @title snake docs api
// @version 1.0
// @description snake demo

// @contact.name 1024casts/snake
// @contact.url http://www.swagger.io/support
// @contact.email

// @host localhost:8080
// @BasePath /v1
func main() {
	pflag.Parse()

	// init config
	if err := conf.Init(*cfg); err != nil {
		panic(err)
	}

	// Set gin mode.
	gin.SetMode(conf.Conf.App.RunMode)

	// init app
	snake.App = snake.New(conf.Conf)

	// Create the Gin engine.
	router := snake.App.Router

	// HealthCheck 健康检查路由
	router.GET("/health", api.HealthCheck)
	// metrics router 可以在 prometheus 中进行监控
	// 通过 grafana 可视化查看 prometheus 的监控数据，使用插件6671查看
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// API Routes.
	routers.Load(router)
	// WEB Routes
	routers.LoadWebRouter(router)

	// init service
	svc := service.New(conf.Conf)

	// set global service
	service.Svc = svc
	snake.App.BizService = svc

	// start grpc server
	var rpcSrv *grpc.Server
	go func() {
		rpcSrv = rpc.New(conf.Conf, svc)
		snake.App.RPCServer = rpcSrv
	}()

	// here register to service discovery

	// start server
	snake.App.Run()
}
