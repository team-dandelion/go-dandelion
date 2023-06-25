package build

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"github.com/gly-hub/toolbox/file"
	"os/exec"
	"path"
)

// 构建网关服务
type HttpTools struct {
	Http      bool
	RpcServer bool
	RpcClient bool
	Logger    bool
	DB        bool
	Redis     bool
	Trace     bool
}

// HttpBuilder http服务构建器
type HttpBuilder struct {
	App         string
	ServerName  string
	serverDir   string
	PackageName string
	Tools       RpcTools
}

// BuildHttpServer 构建rpc服务
func (r *HttpBuilder) BuildHttpServer() {
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
	// 生成cmd目录
	_ = r.buildCmd()
	// 生成config目录
	_ = r.buildConfig()
	// 生成internal目录
	_ = r.buildInternal()
	// 生成static目录
	_ = r.buildStatic()
	// 生成main文件
	_ = r.buildMain()
	// 执行mod
	// 创建mod文件
	cmd2 := exec.Command("go", "mod", "init", r.App)
	_ = cmd2.Run()
	cmd3 := exec.Command("go", "mod", "tidy")
	_ = cmd3.Run()
}

func (r *HttpBuilder) buildCmd() (err error) {
	cmdDir := path.Join(r.serverDir, "cmd")
	if err := file.CreateDir(cmdDir); err != nil {
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

	if err = CreateTemplateFile(serverFile, "internal/template/cmd/apiserver.tmpl", r); err != nil {
		return
	}

	return nil
}

func (r *HttpBuilder) buildConfig() (err error) {
	configDir := path.Join(r.serverDir, "config")
	// 创建应用文件夹
	if err = file.CreateDir(configDir); err != nil {
		return
	}
	configYamlFile := path.Join(configDir, "configs_local.yaml")
	if err = CreateTemplateFile(configYamlFile, "internal/template/config/config.tmpl", r); err != nil {
		return
	}

	return
}

func (r *HttpBuilder) buildInternal() (err error) {
	internalDir := path.Join(r.serverDir, "internal")
	// 创建应用文件夹
	if err = file.CreateDir(internalDir); err != nil {
		return
	}
	middlewareDir := path.Join(internalDir, "middleware")
	// 创建应用文件夹
	if err = file.CreateDir(middlewareDir); err != nil {
		return
	}
	serviceDir := path.Join(internalDir, "service")
	// 创建应用文件夹
	if err = file.CreateDir(serviceDir); err != nil {
		return
	}

	routeDir := path.Join(internalDir, "route")
	// 创建应用文件夹
	if err = file.CreateDir(routeDir); err != nil {
		return
	}
	routeGoFile := path.Join(routeDir, "route.go")
	if err = CreateTemplateFile(routeGoFile, "internal/template/internal/route.tmpl", r); err != nil {
		return
	}
	return
}

func (r *HttpBuilder) buildMain() (err error) {
	mainGoFile := path.Join(r.serverDir, "main.go")
	if err = CreateTemplateFile(mainGoFile, "internal/template/main.tmpl", r); err != nil {
		return
	}
	return
}

func (r *HttpBuilder) buildStatic() (err error) {
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
