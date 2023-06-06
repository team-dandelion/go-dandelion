package build

import (
	"fmt"
	"github.com/gly-hub/go-dandelion-cli/internal/build"
	"github.com/spf13/cobra"
	"strings"
)

var (
	appName  string
	StartCmd = &cobra.Command{
		Use:          "build",
		Short:        "生成服务结构代码",
		Example:      "go-dandelion-cli build -n example-application",
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
	StartCmd.PersistentFlags().StringVarP(&appName, "name", "n", "example-server", "服务器名称")
}

func setup() {
}

func run() error {
	fmt.Print("需要创建的服务类型，输入数字（1-rpc 2-http）:")
	var serverType int
	if _, err := fmt.Scanln(&serverType); err != nil {
		fmt.Println("读取输入时发生错误:", err)
		return nil
	}
	switch serverType {
	case 1:
		return rpc()
	case 2:
		return http()
	}

	return nil
}

func rpc() error {
	var serverName string
	fmt.Print("rpc服务名称:")
	if _, err := fmt.Scanln(&serverName); err != nil {
		fmt.Println("读取输入时发生错误:", err)
		return nil
	}
	var rpcBuilder build.RpcBuilder
	rpcBuilder.App = appName
	rpcBuilder.ServerName = serverName
	fmt.Print("是否初始化mysql（y/n）:")
	var need string
	if _, err := fmt.Scanln(&need); err != nil {
		fmt.Println("读取输入时发生错误:", err)
		return nil
	}
	need = strings.ToLower(need)
	if need == "y" || need == "yes" {
		rpcBuilder.Tools.DB = true
	}

	fmt.Print("是否初始化redis（y/n）:")
	if _, err := fmt.Scanln(&need); err != nil {
		fmt.Println("读取输入时发生错误:", err)
		return nil
	}
	need = strings.ToLower(need)
	if need == "y" || need == "yes" {
		rpcBuilder.Tools.Redis = true
	}

	fmt.Print("是否初始化logger（y/n）:")
	if _, err := fmt.Scanln(&need); err != nil {
		fmt.Println("读取输入时发生错误:", err)
		return nil
	}
	need = strings.ToLower(need)
	if need == "y" || need == "yes" {
		rpcBuilder.Tools.Logger = true
	}

	fmt.Print("是否初始化trace链路（y/n）:")
	if _, err := fmt.Scanln(&need); err != nil {
		fmt.Println("读取输入时发生错误:", err)
		return nil
	}
	need = strings.ToLower(need)
	if need == "y" || need == "yes" {
		rpcBuilder.Tools.Trace = true
	}
	rpcBuilder.BuildRpcServer()
	return nil
}

func http() error {
	var serverName string
	fmt.Print("http服务名称:")
	if _, err := fmt.Scanln(&serverName); err != nil {
		fmt.Println("读取输入时发生错误:", err)
		return nil
	}
	var httpBuilder build.HttpBuilder
	httpBuilder.App = appName
	httpBuilder.ServerName = serverName
	fmt.Print("是否初始化mysql（y/n）:")
	var need string
	if _, err := fmt.Scanln(&need); err != nil {
		fmt.Println("读取输入时发生错误:", err)
		return nil
	}
	need = strings.ToLower(need)
	if need == "y" || need == "yes" {
		httpBuilder.Tools.DB = true
	}

	fmt.Print("是否初始化redis（y/n）:")
	if _, err := fmt.Scanln(&need); err != nil {
		fmt.Println("读取输入时发生错误:", err)
		return nil
	}
	need = strings.ToLower(need)
	if need == "y" || need == "yes" {
		httpBuilder.Tools.Redis = true
	}

	fmt.Print("是否初始化logger（y/n）:")
	if _, err := fmt.Scanln(&need); err != nil {
		fmt.Println("读取输入时发生错误:", err)
		return nil
	}
	need = strings.ToLower(need)
	if need == "y" || need == "yes" {
		httpBuilder.Tools.Logger = true
	}

	fmt.Print("是否初始化trace链路（y/n）:")
	if _, err := fmt.Scanln(&need); err != nil {
		fmt.Println("读取输入时发生错误:", err)
		return nil
	}
	need = strings.ToLower(need)
	if need == "y" || need == "yes" {
		httpBuilder.Tools.Trace = true
	}
	httpBuilder.BuildHttpServer()
	return nil
}
