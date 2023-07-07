package application

func Init() {
	// 初始化数据库
	initDb()

	// 初始化redis
	initRedis()

	// 初始化日志
	initLogger()

	// 初始化http
	initHttpServer()

	// 初始化错误文本
	//error_support.Init(".")

	// 初始化Rpc Client
	initRpcClient()

	// 初始化分析器
	initAnalysisServer()

	// 初始化链路追踪
	initTracer()
}
