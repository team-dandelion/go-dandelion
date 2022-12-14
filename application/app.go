package application

func Init(){
	// 初始化数据库
	initDb()

	// 初始化redis
	initRedis()

	// 初始化日志
	initLogger()
}
