package build

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"github.com/gly-hub/toolbox/file"
	"os/exec"
	"path"
)

type RpcTools struct {
	Http      bool
	RpcServer bool
	RpcClient bool
	Logger    bool
	DB        bool
	Redis     bool
	Trace     bool
}

// RpcBuilder rpc服务构建器
type RpcBuilder struct {
	App         string
	ServerName  string
	serverDir   string
	PackageName string
	Tools       RpcTools
}

// BuildRpcServer 构建rpc服务
func (r *RpcBuilder) BuildRpcServer() {
	// 生成服务目录
	pwd, err := file.GetPwd()
	if err != nil {
		return
	}
	serverDir := path.Join(pwd, r.ServerName)
	// 创建应用文件夹
	if err = file.CreateDir(serverDir); err != nil {
		return
	}
	r.serverDir = serverDir
	// 生成Boot目录
	_ = r.buildBoot()
	// 生成cmd目录
	_ = r.buildCmd()
	// 生成config目录
	_ = r.buildConfig()
	// 生成global目录
	_ = r.buildGlobal()
	// 生成internal目录
	_ = r.buildInternal()
	// 生成static目录
	_ = r.buildStatic()
	// 生成Tools目录
	_ = r.buildTools()
	// 生成main文件
	_ = r.buildMain()
	// 执行mod
	// 创建mod文件
	cmd2 := exec.Command("go", "mod", "init", r.App)
	_ = cmd2.Run()
	cmd3 := exec.Command("go", "mod", "tidy")
	_ = cmd3.Run()
}

func (r *RpcBuilder) buildBoot() (err error) {
	bootDir := path.Join(r.serverDir, "boot")
	// 创建应用文件夹
	if err = file.CreateDir(bootDir); err != nil {
		return
	}
	bootGoFile := path.Join(bootDir, "boot.go")
	if err = CreateTemplateFile(bootGoFile, "internal/template/boot/boot.tmpl", r); err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}

func (r *RpcBuilder) buildCmd() (err error) {
	cmdDir := path.Join(r.serverDir, "cmd")
	if err = file.CreateDir(cmdDir); err != nil {
		return err
	}

	bootGoFile := path.Join(cmdDir, "cobra.go")
	if err = CreateTemplateFile(bootGoFile, "internal/template/cmd/cobra.tmpl", r); err != nil {
		return
	}

	apiDir := path.Join(cmdDir, "api")
	// 创建应用文件夹
	if err = file.CreateDir(apiDir); err != nil {
		return
	}
	serverFile := path.Join(apiDir, "server.go")

	if err = CreateTemplateFile(serverFile, "internal/template/cmd/rpcserver.tmpl", r); err != nil {
		return
	}

	return nil
}

func (r *RpcBuilder) buildConfig() (err error) {
	configDir := path.Join(r.serverDir, "config")
	// 创建应用文件夹
	if err = file.CreateDir(configDir); err != nil {
		return
	}
	configYamlFile := path.Join(configDir, "configs_local.yaml")
	if err = CreateTemplateFile(configYamlFile, "internal/template/config/config.tmpl", r); err != nil {
		fmt.Println(err.Error())
		return
	}

	return
}

func (r *RpcBuilder) buildGlobal() (err error) {
	globalDir := path.Join(r.serverDir, "global")
	// 创建应用文件夹
	if err = file.CreateDir(globalDir); err != nil {
		return
	}
	globalGoFile := path.Join(globalDir, "global.go")
	if err = CreateTemplateFile(globalGoFile, "internal/template/global/global.tmpl", r); err != nil {
		return
	}
	return
}

func (r *RpcBuilder) buildTools() (err error) {
	ToolsDir := path.Join(r.serverDir, "Tools")
	// 创建应用文件夹
	if err = file.CreateDir(ToolsDir); err != nil {
		return
	}
	return
}

func (r *RpcBuilder) buildMain() (err error) {
	mainGoFile := path.Join(r.serverDir, "main.go")
	if err = CreateTemplateFile(mainGoFile, "internal/template/main.tmpl", r); err != nil {
		return
	}
	return
}

func (r *RpcBuilder) buildStatic() (err error) {
	staticDir := path.Join(r.serverDir, "static")
	// 创建应用文件夹
	if err = file.CreateDir(staticDir); err != nil {
		return
	}
	staticTxtFile := path.Join(staticDir, fmt.Sprintf("%s.txt", r.ServerName))
	if err = file.CreateFile(staticTxtFile); err != nil {
		return
	}
	myFigure := figure.NewFigure(r.ServerName, "", true)
	data := myFigure.String()
	return file.WriteFile(staticTxtFile, data)
}

func (r *RpcBuilder) buildInternal() (err error) {
	internalDir := path.Join(r.serverDir, "internal")
	// 创建应用文件夹
	if err = file.CreateDir(internalDir); err != nil {
		return
	}
	enumDir := path.Join(internalDir, "enum")
	// 创建应用文件夹
	if err = file.CreateDir(enumDir); err != nil {
		return
	}
	daoDir := path.Join(internalDir, "dao")
	// 创建应用文件夹
	if err = file.CreateDir(daoDir); err != nil {
		return
	}
	logicDir := path.Join(internalDir, "logic")
	// 创建应用文件夹
	if err = file.CreateDir(logicDir); err != nil {
		return
	}
	modelDir := path.Join(internalDir, "model")
	// 创建应用文件夹
	if err = file.CreateDir(modelDir); err != nil {
		return
	}

	serviceDir := path.Join(internalDir, "service")
	// 创建应用文件夹
	if err = file.CreateDir(serviceDir); err != nil {
		return
	}
	apiGoFile := path.Join(serviceDir, "api.go")
	if err = CreateTemplateFile(apiGoFile, "internal/template/internal/rpcapi.tmpl", r); err != nil {
		return
	}
	return
}
