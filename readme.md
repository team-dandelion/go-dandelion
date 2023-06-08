[![Security Status](https://www.murphysec.com/platform3/v31/badge/1666706410635550720.svg)](https://www.murphysec.com/console/report/1666706410597801984/1666706410635550720)
## 关于go-dandelion
go-dandelion项目意在通过集成相关组件，方便开发者快速构建项目框架，提升开发效率。不在浪费时间在各组组件的集成上，可快速进行业务开发。

**集成**
+ [rpc](https://github.com/smallnest/rpcx)
+ <a herf="https://github.com/valyala/fasthttp">fasthttp</a>
+ <a herf="https://github.com/qiangxue/fasthttp-routing">fasthttp-routing</a>
+ <a herf="https://github.com/go-gorm/gorm">gorm</a>
+ <a herf="https://github.com/gomodule/redigo">redigo</a>
+ <a herf="https://github.com/go-swagger/go-swagger">swagger</a>
+ <a herf="https://github.com/spf13/cobra">cobra</a>
+ <a herf="https://github.com/spf13/viper">viper</a>
+ <a herf="https://github.com/opentracing/opentracing-go">opentracing-go</a>


## go-dandelion-cli使用

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
```

### 3.运行项目
```shell
#进入服务目录
go build -o 服务名称
#运行
./服务名称 server
```

## 🔥贡献者

<a href="https://github.com/gly-hub/go-dandelion/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=gly-hub/go-dandelion" />
</a>

##  ⭐点个star吧！

如果你对该项目感兴趣，请点个星哦！

## 🔑开源
[Apache License, Version 2.0](LICENSE.txt)
