package application

import "github.com/gly-hub/go-dandelion/config"

/*
 插件功能。
	— 自动初始化插件
*/

type Plugin interface {
	// Config 返回需要初始化的配置，建议使用第二层级
	Config() interface{}
	InitPlugin() error
}

func Plugs(plugins ...Plugin) error {
	for _, plug := range plugins {
		config.LoadCustomConfig(plug.Config())
		if err := plug.InitPlugin(); err != nil {
			return err
		}
	}
	return nil
}
