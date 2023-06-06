package code

import "fmt"

// 配置相关代码

// ConfigYamlLogger 配置yaml日志
func ConfigYamlLogger() string {
	return `logger:
  #Level 0 紧急的 1警报 2重要的 3错误 4警告 5提示 6信息 7调试
  consoleShow: true
  consoleLevel:  7
  fileWrite:  false
  fileLevel:  7
  multiFileWrite: false
  multiFileLevel: 7
`
}

func ConfigYamlHttpServer() string {
	return `httpServer:
  port: 8080
  pprof: 8088`
}

// ConfigYamlRpcServer 配置yaml rpc
func ConfigYamlRpcServer(app, server string) string {
	return fmt.Sprintf(`rpcServer:
  model: 1
  serverName: "%s"
  etcd: ["127.0.0.1:2379"]
  basePath: "%s"
  addr: ""
  port: 8899
  pprof: 18899
`, server, app)
}

// ConfigYamlRpcServerForHttp 配置yaml rpc
func ConfigYamlRpcServerForHttp(app, server string) string {
	return fmt.Sprintf(`rpcServer:
  model: 1
  serverName: "%s"
  etcd: ["127.0.0.1:2379"]
  basePath: "%s"
`, server, app)
}

// ConfigYamlDB 配置yaml db
func ConfigYamlDB() string {
	return `db:
  dbType: "mysql"
  maxOpenConn: 20
  maxIdleConn: 4
  maxIdleTime: 100
  maxLifeTime: 3600
  level: 4
  slowThreshold: "100ms"
  master:
    user: "root"
    password: "password"
    host: "127.0.0.1"
    port: "3306"
    database: "database"
  slave:
    - user: "root"
      password: "password"
      host: "127.0.0.1"
      port: "3306"
      database: "database"
`
}

// ConfigYamlRedis 配置yaml redis
func ConfigYamlRedis() string {
	return `redis:
  redisType: "alone"
  network: "127.0.0.1:6379"
  startAddr: ["127.0.0.1:6379"]
  active: 100
  idle: 100
  auth: ""
  connTimeout: "100ms"
  readTimeout: "100ms"
  writeTimeout: "100ms"
  idleTimeout: "100ms"
`
}

// ConfigYamlTrace 配置yaml trace
func ConfigYamlTrace() string {
	return `tracer:
  openTrace: true
  traceName: "serverName"
  host: "127.0.0.1:6831"
`
}
