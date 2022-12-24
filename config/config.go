package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	_ "gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var (
	// 文件名
	configFileName string
	// 文件后缀
	configFilePostfix string
	// 文件地址
	configFilePath string
)

// 支持的配置后缀
const (
	JSON = "json"
	YAML = "yaml"
	TOML = "toml"
	INI  = "ini"
)

// 配置区分
const (
	Local      = "local"      // 本地环境
	Develop    = "develop"    // 测试环境
	Release    = "release"    // 预发布环境
	Production = "production" // 生产环境
)

var Conf Config

type Config struct {
	WDB        *WDB        `json:"wdb" yaml:"wdb"`
	RDB        *RDB        `json:"rdb" yaml:"rdb"`
	Redis      *Redis      `json:"redis" yaml:"redis"`
	Logger     *Logger     `json:"logger" yaml:"logger"`
	HttpServer *HttpServer `json:"http_server" yaml:"httpServer"`
	RpcServer  *RpcServer  `json:"rpc_server" yaml:"rpcServer"`
}

// InitConfig 需要设置默认环境变量，当系统中存在环境变量值时，会优先使用
// 系统标量（DANDELION_ENV）。configs 为自定义配置，可系统加载用户自定
// 义配置
func InitConfig(env string, configs ...interface{}) {
	// 获取系统环境变量
	osEnv := os.Getenv("DANDELION_ENV")
	if osEnv != "" && osEnv != env {
		env = osEnv
	}

	// 判断是否支持该环境类型
	if env != Local && env != Develop && env != Release && env != Production {
		panic(errors.New("不支持该环境"))
	}

	// 根据环境获取对应配置文件
	configFileName = fmt.Sprintf("configs_%s", env)

	// 加载配置
	initConfig(configs)
}

func initConfig(configs ...interface{}) {
	// 获取配置目录
	lookingConfig()

	// 获取指定文件夹下的配置列表
	// 区分不同的配置类型，json、yaml、toml、ini
	// 多个文件类型并存时，取并集
	// 获取目录下的所有文件
	rd, err := ioutil.ReadDir(appConfigPath)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return
	}

	for _, fi := range rd {
		if !fi.IsDir() {
			// 文件后缀
			postfix := strings.ReplaceAll(path.Ext(fi.Name()), ".", "")
			if postfix != JSON && postfix != YAML && postfix != TOML && postfix != INI {
				continue
			}
			if fi.Name()[:len(fi.Name())-len(postfix)-1] == configFileName {
				configFilePostfix = postfix
				break
			}
		}
	}

	// 加载配置
	if configFilePostfix != YAML {
		panic(errors.New("未检测到支持的配置文件类型。支持json、yaml、toml、ini等后缀。"))
	}

	vip := viper.New()
	vip.AddConfigPath(appConfigPath)
	vip.SetConfigName(configFileName)
	vip.SetConfigType(configFilePostfix)

	//尝试进行配置读取
	if err := vip.ReadInConfig(); err != nil {
		panic(err)
	}
	err = vip.Unmarshal(&Conf)
	if err != nil {
		panic(err)
	}

	// 加载自定义配置
	for i, _ := range configs {
		err = vip.Unmarshal(&configs[i])
		if err != nil {
			panic(err)
		}
	}
}
