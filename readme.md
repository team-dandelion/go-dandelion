[中文](readme-ZH.md)|English
## About go-dandelion

The go-dandelion project aims to provide developers with a project framework that integrates various components, making it easy to build projects and improve development efficiency. It eliminates the need to spend time on integrating different components, allowing developers to focus on business development.

[![Go](https://github.com/team-dandelion/go-dandelion/workflows/Go/badge.svg?branch=main)](https://github.com/team-dandelion/go-dandelion/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/team-dandelion/go-dandelion)](https://goreportcard.com/report/github.com/team-dandelion/go-dandelion)
[![codecov](https://codecov.io/gh/gly-hub/go-dandelion/branch/main/graph/badge.svg)](https://codecov.io/gh/gly-hub/go-dandelion)
[![MIT license](https://img.shields.io/badge/License-Apache2.0-brightgreen.svg)](https://opensource.org/licenses/apache-2-0/)
[![Release](https://img.shields.io/badge/release-1.2.0-white.svg)](https://pkg.go.dev/github.com/team-dandelion/go-dandelion/go-dandelion-cli?tab=doc)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/team-dandelion/go-dandelion/go-dandelion-cli?tab=doc)

[![Security Status](https://www.murphysec.com/platform3/v31/badge/1666706410635550720.svg)](https://www.murphysec.com/console/report/1666706410597801984/1666706410635550720)


**Integration**
+ [rpcx](https://github.com/smallnest/rpcx)
+ [fasthttp](https://github.com/valyala/fasthttp)
+ [fasthttp-routing](https://github.com/qiangxue/fasthttp-routing)
+ [gorm](https://github.com/go-gorm/gorm)
+ [redigo](https://github.com/gomodule/redigo)
+ [go-swagger](https://github.com/go-swagger/go-swagger)
+ [cobra](https://github.com/spf13/cobra)
+ [viper](https://github.com/spf13/viper)
+ [opentracing-go](https://github.com/opentracing/opentracing-go)

**Features**
+ Quickly create RPC services and HTTP services.
+ Initialize MySQL, Redis, logger, and trace links quickly through configuration.
+ Integrated logging, distributed tracing, rate limiting, circuit breaking, service registration, service discovery, and other features.
+ Customizable middleware and plugins.

## go-dandelion-cli Usage

## 1. Installation
```
go get github.com/team-dandelion/go-dandelion/go-dandelion-cli@latest
go install github.com/team-dandelion/go-dandelion/go-dandelion-cli@latest
```

## 2. Create a Project
Create a local project directory and create the corresponding project based on the prompts.
```shell
# Create an application
go-dandelion-cli app -n go-admin-example
# Enter the application directory
cd go-admin-example
# Build the service
go-dandelion-cli build -n go-admin-example
Enter the type of service to create, enter a number (1 for rpc, 2 for http): 1
RPC service name: example-server
Initialize MySQL (y/n): y
Initialize Redis (y/n): y
Initialize logger (y/n): y
Initialize trace links (y/n): y
```

## 3. Run the Project
```shell
cd example-server
# Enter the service directory
go build -o example-server
# Run the service
./example-server server
```

## 🔥Contributors

<a href="https://github.com/team-dandelion/go-dandelion/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=gly-hub/go-dandelion" />
</a>

## ⭐ Star the project
if you find it interesting!

## Open Source
[MIT License](LICENSE.txt)
