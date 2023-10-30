package gorm

import (
	timex "github.com/gly-hub/toolbox/time"
	jsoniter "github.com/json-iterator/go"
	"github.com/team-dandelion/go-dandelion/config"
	"github.com/team-dandelion/go-dandelion/database/redigo"
	"github.com/team-dandelion/go-dandelion/database/redigo/tools"
	"gorm.io/gorm"
	"sync"
)

// AppConfigFunc 用于获取需要初始化的应用库
// 初始化时，不会传递参数。应用数据库变更时，
// 会传递 appKey
type AppConfigFunc func(appKey ...string) (apps []AppConfig, err error)

// AppChangeFunc 数据连接变更监听触发方法
type AppChangeFunc func(appKey string, changeType ChangeType) (err error)

type AppDB struct {
	db            sync.Map
	appConfigFunc AppConfigFunc
	appChangeFunc AppChangeFunc
	redis         *redigo.Client
}

func InitAppDB(configFunc AppConfigFunc, changeFunc AppChangeFunc, redis *redigo.Client) *AppDB {
	appDB := &AppDB{
		db:            sync.Map{},
		appConfigFunc: configFunc,
		appChangeFunc: changeFunc,
		redis:         redis,
	}
	apps, err := configFunc()
	if err != nil {
		panic(err)
	}

	for _, app := range apps {
		appDB.db.Store(app.AppKey, NewConnection(&Config{
			DBType:        config.Conf.DB.DBType,
			MaxOpenConn:   config.Conf.DB.MaxOpenConn,
			MaxIdleConn:   config.Conf.DB.MaxIdleConn,
			MaxLifeTime:   config.Conf.DB.MaxLifeTime,
			MaxIdleTime:   config.Conf.DB.MaxIdleTime,
			Level:         config.Conf.DB.Level,
			SlowThreshold: timex.ParseDuration(config.Conf.DB.SlowThreshold),
			Master:        app.DBMaster,
			Slaves:        app.DBSlave,
		}))
	}
	return appDB
}

func (client *AppDB) GetDB(appKey string) *gorm.DB {
	if appKey == "" {
		return nil
	}
	if db, ok := client.db.Load(appKey); ok {
		return db.(*gorm.DB)
	}
	return nil
}

// Listener 数据库连接变更监听
func (client *AppDB) Listener() {
	// 注册订阅方法
	tools.RegisterFunc(GoDandelionConnChange, func(pattern, channel, device string) {
		// 解析数据
		msg := ParseChangeData(device)
		switch msg.ChangeType {
		case Created:
			_ = client.Connection(msg.AppKey)
			_ = client.appChangeFunc(msg.AppKey, msg.ChangeType)
		case Updated:
			_ = client.Connection(msg.AppKey)
			_ = client.appChangeFunc(msg.AppKey, msg.ChangeType)
		case Deleted:
			client.db.Delete(msg.AppKey)
		}
	})
	//go tools.PSubscribeListen(client.redis)
}

func (client *AppDB) Connection(appKey string) error {
	if appKey == "" {
		return nil
	}
	// 重新获取数据库连接
	apps, err := client.appConfigFunc(appKey)
	if err != nil {
		panic(err)
	}
	for _, app := range apps {
		client.db.Store(app.AppKey, NewConnection(&Config{
			DBType:        config.Conf.DB.DBType,
			MaxOpenConn:   config.Conf.DB.MaxOpenConn,
			MaxIdleConn:   config.Conf.DB.MaxIdleConn,
			MaxLifeTime:   config.Conf.DB.MaxLifeTime,
			MaxIdleTime:   config.Conf.DB.MaxIdleTime,
			Level:         config.Conf.DB.Level,
			SlowThreshold: timex.ParseDuration(config.Conf.DB.SlowThreshold),
			Master:        app.DBMaster,
			Slaves:        app.DBSlave,
		}))
	}
	return nil
}

// ParseChangeData
func ParseChangeData(msg string) AppDBMessage {
	var data AppDBMessage
	_ = jsoniter.UnmarshalFromString(msg, &data)
	return data
}
