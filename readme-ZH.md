中文|[English](readme.md)

## 📚关于go-dandelion
go-dandelion项目意在通过集成相关组件，方便开发者快速构建项目框架，提升开发效率。不在浪费时间在各组组件的集成上，可快速进行业务开发。

[![Go](https://github.com/gly-hub/go-dandelion/workflows/Go/badge.svg?branch=main)](https://github.com/gly-hub/go-dandelion/actions)
[![codecov](https://codecov.io/gh/gly-hub/go-dandelion/branch/main/graph/badge.svg)](https://codecov.io/gh/gly-hub/go-dandelion)
[![MIT license](https://img.shields.io/badge/License-Apache2.0-brightgreen.svg)](https://opensource.org/licenses/apache-2-0/)
[![Release](https://img.shields.io/badge/release-1.2.0-white.svg)](https://pkg.go.dev/github.com/gly-hub/go-dandelion/go-dandelion-cli?tab=doc)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/gly-hub/go-dandelion/go-dandelion-cli?tab=doc)

[![Security Status](https://www.murphysec.com/platform3/v31/badge/1666706410635550720.svg)](https://www.murphysec.com/console/report/1666706410597801984/1666706410635550720)


**集成**
+ [rpcx](https://github.com/smallnest/rpcx)
+ [fasthttp](https://github.com/valyala/fasthttp)
+ [fasthttp-routing](https://github.com/qiangxue/fasthttp-routing)
+ [gorm](https://github.com/go-gorm/gorm)
+ [redigo](https://github.com/gomodule/redigo)
+ [go-swagger](https://github.com/go-swagger/go-swagger)
+ [cobra](https://github.com/spf13/cobra)
+ [viper](https://github.com/spf13/viper)
+ [opentracing-go](https://github.com/opentracing/opentracing-go)

**功能**
+ 快速创建rpc服务、http服务
+ 通过配置，快速初始化mysql、redis、logger、trace链路等
+ 集成日志打印、链路追踪、限流、熔断、服务注册、服务发现等功能
+ 可自定义中间件、插件

## 🖥go-dandelion-cli使用

### 1.安装
```
go get github.com/gly-hub/go-dandelion/go-dandelion-cli@latest
go install github.com/gly-hub/go-dandelion/go-dandelion-cli@latest
```

### 2.创建项目
创建本地项目目录，根据提示创建对应项目
```shell
# 创建应用
go-dandelion-cli app -n go-admin-example
# 进入应用目录
cd go-admin-example
# 构建服务
go-dandelion-cli build -n go-admin-example
需要创建的服务类型，输入数字（1-rpc 2-http）:1
rpc服务名称:example-server
是否初始化mysql（y/n）:y
是否初始化redis（y/n）:y
是否初始化logger（y/n）:y
是否初始化trace链路（y/n）:y
```

### 3.运行项目
```shell
cd example-server
#进入服务目录
go build -o example-server
#运行
./example-server server
```

## 🔥贡献者

<a href="https://github.com/gly-hub/go-dandelion/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=gly-hub/go-dandelion" />
</a>

##  ⭐点个star吧！

如果你对该项目感兴趣，请点个星哦！

## 🔑开源
[Apache License, Version 2.0](LICENSE.txt)
