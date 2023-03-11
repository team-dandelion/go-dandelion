package gorm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
	"time"
)

const (
	Mysql    = "mysql"
	Postgres = "postgres"
)

func NewConnection(config *Config) *gorm.DB {
	var db *gorm.DB
	switch config.DBType {
	case Mysql:
		db = initMysql(config)
	case Postgres:
		db = initPostgres(config)
	default:
		panic("未检测到支持的数据库类型")
	}
	return db
}

func initMysql(config *Config) *gorm.DB {
	if config.Master == nil {
		panic("未检测到主库配置")
	}
	// 初始化master
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		config.Master.User,
		config.Master.Password,
		config.Master.Host,
		config.Master.Port,
		config.Master.DataBase)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: &Logger{
		Level:         logger.LogLevel(config.Level),
		SlowThreshold: config.SlowThreshold,
	}})
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.LogLevel(config.Level))})
	if err != nil {
		panic(err)
	}

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, dErr := db.DB()
	if dErr != nil {
		panic(dErr)
	}

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	if config.Slaves != nil {
		// 初始化slave
		var slaveDsns []gorm.Dialector
		for _, slave := range config.Slaves {
			slaveDsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
				slave.User,
				slave.Password,
				slave.Host,
				slave.Port,
				slave.DataBase)
			slaveDsns = append(slaveDsns, mysql.Open(slaveDsn))
		}
		db.Use(dbresolver.Register(dbresolver.Config{
			Replicas:          slaveDsns,
			Policy:            dbresolver.RandomPolicy{},
			TraceResolverMode: true,
		}).
			SetConnMaxIdleTime(time.Duration(config.MaxIdleTime) * time.Second).
			SetConnMaxLifetime(time.Duration(config.MaxLifeTime) * time.Second).
			SetMaxIdleConns(config.MaxIdleConn).
			SetMaxOpenConns(config.MaxOpenConn))
	}
	return db
}

func initPostgres(config *Config) *gorm.DB {
	if config.Master == nil {
		panic("未检测到主库配置")
	}
	// 初始化master
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		config.Master.Host,
		config.Master.Port,
		config.Master.User,
		config.Master.DataBase,
		config.Master.Password)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, dErr := db.DB()
	if dErr != nil {
		panic(dErr)
	}

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	if config.Slaves != nil {
		// 初始化slave
		var slaveDsns []gorm.Dialector
		for _, slave := range config.Slaves {
			slaveDsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
				slave.Host,
				slave.Port,
				slave.User,
				slave.DataBase,
				slave.Password)
			slaveDsns = append(slaveDsns, postgres.Open(slaveDsn))
		}
		db.Use(dbresolver.Register(dbresolver.Config{
			Replicas:          slaveDsns,
			Policy:            dbresolver.RandomPolicy{},
			TraceResolverMode: true,
		}).
			SetConnMaxIdleTime(time.Duration(config.MaxIdleTime) * time.Second).
			SetConnMaxLifetime(time.Duration(config.MaxLifeTime) * time.Second).
			SetMaxIdleConns(config.MaxIdleConn).
			SetMaxOpenConns(config.MaxOpenConn))
	}
	return db
}
