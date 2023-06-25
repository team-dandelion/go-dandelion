package build

import (
	"errors"
	"fmt"
	"github.com/gly-hub/toolbox/file"
	"path"
	"strings"
)

// 构建应用基础架构

// BuildApplication 创建应用
func BuildApplication(app string) {
	pwd, err := file.GetPwd()
	if err != nil {
		return
	}
	appDir := path.Join(pwd, app)
	// 创建应用文件夹
	if err = file.CreateDir(appDir); err != nil {
		return
	}
}

func Rpc(appName string) error {
	var serverName string
	fmt.Print("RPC Server Name:")
	if _, err := fmt.Scanln(&serverName); err != nil {
		fmt.Println("An error occurred while reading the input:", err)
		return nil
	}
	var rpcBuilder RpcBuilder
	rpcBuilder.Tools.RpcServer = true
	rpcBuilder.PackageName = fmt.Sprintf("%s/%s", appName, serverName)
	rpcBuilder.App = appName
	rpcBuilder.ServerName = serverName

	var err error
	rpcBuilder.Tools.DB, err = EnterBool("Initialize Mysql")
	if err != nil {
		return err
	}

	rpcBuilder.Tools.Redis, err = EnterBool("Initialize Redis")
	if err != nil {
		return err
	}

	rpcBuilder.Tools.Logger, err = EnterBool("Initialize Logger")
	if err != nil {
		return err
	}

	rpcBuilder.Tools.Trace, err = EnterBool("Initialize Trace Links")
	if err != nil {
		return err
	}

	rpcBuilder.BuildRpcServer()
	return nil
}

func Http(appName string) error {
	var serverName string
	fmt.Print("HTTP Server Name:")
	if _, err := fmt.Scanln(&serverName); err != nil {
		fmt.Println("An error occurred while reading the input:", err)
		return nil
	}
	var httpBuilder HttpBuilder
	httpBuilder.Tools.Http = true
	httpBuilder.Tools.RpcClient = true
	httpBuilder.PackageName = fmt.Sprintf("%s/%s", appName, serverName)
	httpBuilder.App = appName
	httpBuilder.ServerName = serverName
	var err error
	httpBuilder.Tools.DB, err = EnterBool("Initialize Mysql")
	if err != nil {
		return err
	}

	httpBuilder.Tools.Redis, err = EnterBool("Initialize Redis")
	if err != nil {
		return err
	}

	httpBuilder.Tools.Logger, err = EnterBool("Initialize Logger")
	if err != nil {
		return err
	}

	httpBuilder.Tools.Trace, err = EnterBool("Initialize Trace Links")
	if err != nil {
		return err
	}
	httpBuilder.BuildHttpServer()
	return nil
}

func EnterBool(text string) (bool, error) {
	var need string
	fmt.Print(fmt.Sprintf("%s（y/n）:", text))
	if _, err := fmt.Scanln(&need); err != nil {
		fmt.Println()
		return false, errors.New(fmt.Sprintf("An error occurred while reading the input:%s", err))
	}
	need = strings.ToLower(need)
	if need == "y" || need == "yes" {
		return true, nil
	}
	return false, nil
}
