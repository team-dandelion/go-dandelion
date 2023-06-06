package code

import "fmt"

// CmdCobra 构建cmd
func CmdCobra(app, server string) string {
	return fmt.Sprintf(`package cmd

import (
	"errors"
	"%s/%s/cmd/api"
	"github.com/gly-hub/go-dandelion/logger"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "%s/%s",
	Short: "%s",
	SilenceUsage:true,
	Long: "authorize",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New(logger.Red("requires at least one arg"))
		}
		return nil
	},
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init(){
	rootCmd.AddCommand(api.StartCmd)
}

func Execute(){
	if err := rootCmd.Execute(); err != nil{
		os.Exit(-1)
	}
}`, app, server, app, server, server)
}

// CmdRpc 构建rpc cmd
func CmdRpc(app, server string) string {
	pack := app + "/" + server

	c := `fmt.Printf("%s Shutdown Server ... \r\n", stringx.GetCurrentTimeStr())
	logger.Info("Server exiting")
	return nil`

	return fmt.Sprintf(`package api

import (
	"fmt"
	"%s/boot"
	"%s/internal/service"
	"github.com/gly-hub/go-dandelion/application"
	"github.com/gly-hub/go-dandelion/config"
	"github.com/gly-hub/go-dandelion/logger"
	"github.com/gly-hub/go-dandelion/tools/stringx"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/signal"
)

var (
	env      string
	StartCmd = &cobra.Command{
		Use:          "server",
		Short:        "Start RPC server",
		Example:      "%s server -e local",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&env, "env", "e", "local", "Env")
}

func setup() {
	// 配置初始化
	config.InitConfig(env)
	// 应用初始化
	application.Init()
	// 初始化服务方法
	boot.Init()
}

func run() error {
	// 初始化rpc model
	go func() {
		application.RpcServer(new(service.RpcApi))
	}()
	content, _ := ioutil.ReadFile("./static/%s.txt")
	fmt.Println(logger.Green(string(content)))
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	%s
}
`, pack, pack, server, server, c)
}

// CmdHttp 构建http cmd
func CmdHttp(app, server string) string {
	pack := app + "/" + server

	c := `fmt.Println(logger.Green(string(content)))
	fmt.Println(logger.Green("Server run at:"))
	fmt.Printf("-  Local:   http://localhost:%d/ \r\n", application.HttpServer().Port())
	fmt.Printf("-  Network: http://%s:%d/ \r\n", ip.GetLocalHost(), application.HttpServer().Port())
	fmt.Println()
	if config.GetEnv() != "production" {
		fmt.Println(logger.Green("Swagger run at:"))
		fmt.Printf("-  Local:   http://localhost:%d/api/swagger/index.html \r\n", application.HttpServer().Port())
		fmt.Printf("-  Network: http://%s:%d/api/swagger/index.html \r\n", ip.GetLocalHost(), application.HttpServer().Port())
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Printf("%s Shutdown Server ... \r\n", stringx.GetCurrentTimeStr())
	logger.Info("Server exiting")
	return nil`

	return fmt.Sprintf(`package api

import (
	"fmt"
	routing "github.com/gly-hub/fasthttp-routing"
	"%s/internal/route"
	"github.com/gly-hub/go-dandelion/application"
	"github.com/gly-hub/go-dandelion/config"
	"github.com/gly-hub/go-dandelion/logger"
	"github.com/gly-hub/go-dandelion/tools/ip"
	"github.com/gly-hub/go-dandelion/tools/stringx"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/signal"
)

var (
	env      string
	StartCmd = &cobra.Command{
		Use:          "server",
		Short:        "Start API server",
		Example:      "%s server -e local",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&env, "env", "e", "local", "Env")
}

func setup() {
	// 配置初始化
	config.InitConfig(env)
	// 应用初始化
	application.Init()
	// 路由初始化
	route.InitRoute()
	// 注册头部context链路
	application.RegisterHeaderFunc(HeaderFunc)
}

func HeaderFunc(ctx *routing.Context, data map[string]string) map[string]string {
	// 自定义头部链路。该方法能将需要的参数通过rpc进行传递 TODO

	return data
}

func run() error {
	// 启动http服务
	go func() {
		application.HttpServer().Server()
	}()
	content, _ := ioutil.ReadFile("./static/internal-gateway.txt")
	%s
}`, pack, server, c)
}
