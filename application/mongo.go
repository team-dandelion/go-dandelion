package application

import (
	"context"
	"errors"
	"time"

	timex "github.com/gly-hub/toolbox/time"
	zedis "github.com/gomodule/redigo/redis"
	jsoniter "github.com/json-iterator/go"
	"github.com/team-dandelion/go-dandelion/config"
	dmongo "github.com/team-dandelion/go-dandelion/database/mongo"
	"github.com/team-dandelion/go-dandelion/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	baseMongoDB *mongo.Client
	appMongoDB  *dmongo.AppMongoDB
)

func initMongoDB() {
	if config.Conf.Mongo != nil {
		// 创建基础MongoDB连接配置
		mongoConfig := &dmongo.Config{
			MaxPoolSize:            uint64(100),            // 默认值
			MinPoolSize:            uint64(10),             // 默认值
			MaxConnIdleTime:        30 * time.Minute,       // 默认值
			ConnectTimeout:         10 * time.Second,       // 默认值
			SocketTimeout:          30 * time.Second,       // 默认值
			ServerSelectionTimeout: 30 * time.Second,       // 默认值
			LogLevel:               1,                      // 默认值
			SlowThreshold:          100 * time.Millisecond, // 默认值
		}

		// 如果配置存在则使用配置值
		if config.Conf.Mongo.MaxPoolSize > 0 {
			mongoConfig.MaxPoolSize = uint64(config.Conf.Mongo.MaxPoolSize)
		}
		if config.Conf.Mongo.MinPoolSize > 0 {
			mongoConfig.MinPoolSize = uint64(config.Conf.Mongo.MinPoolSize)
		}
		if config.Conf.Mongo.MaxConnIdleTime > 0 {
			mongoConfig.MaxConnIdleTime = time.Duration(config.Conf.Mongo.MaxConnIdleTime) * time.Second
		}
		if config.Conf.Mongo.ConnectTimeout != "" {
			mongoConfig.ConnectTimeout = timex.ParseDuration(config.Conf.Mongo.ConnectTimeout)
		}
		if config.Conf.Mongo.SocketTimeout != "" {
			mongoConfig.SocketTimeout = timex.ParseDuration(config.Conf.Mongo.SocketTimeout)
		}
		if config.Conf.Mongo.ServerSelectionTimeout != "" {
			mongoConfig.ServerSelectionTimeout = timex.ParseDuration(config.Conf.Mongo.ServerSelectionTimeout)
		}
		mongoConfig.LogLevel = config.Conf.Mongo.LogLevel
		if config.Conf.Mongo.SlowThreshold != "" {
			mongoConfig.SlowThreshold = timex.ParseDuration(config.Conf.Mongo.SlowThreshold)
		}

		// 设置MongoDB连接配置
		mongoConfig.Hosts = config.Conf.Mongo.Hosts
		mongoConfig.Database = config.Conf.Mongo.Database
		mongoConfig.Username = config.Conf.Mongo.Username
		mongoConfig.Password = config.Conf.Mongo.Password
		mongoConfig.AuthDB = config.Conf.Mongo.AuthDB
		mongoConfig.ReplicaSet = config.Conf.Mongo.ReplicaSet
		mongoConfig.ReadPreference = config.Conf.Mongo.ReadPreference

		// 如果有主机配置才创建连接
		if len(mongoConfig.Hosts) > 0 && mongoConfig.Database != "" {
			baseMongoDB = dmongo.NewConnection(mongoConfig)
		}
	}
	return
}

type MongoDB struct {
}

// GetClient 获取MongoDB客户端连接
func (*MongoDB) GetClient(appKeys ...string) *mongo.Client {
	if len(appKeys) > 0 {
		return appMongoDB.GetClient(appKeys[0])
	}
	return baseMongoDB
}

// GetDatabase 获取MongoDB数据库实例
func (*MongoDB) GetDatabase(dbName string, appKeys ...string) *mongo.Database {
	if len(appKeys) > 0 {
		return appMongoDB.GetDatabase(appKeys[0], dbName)
	}
	if baseMongoDB != nil {
		return dmongo.GetDatabase(baseMongoDB, dbName)
	}
	return nil
}

// GetCollection 获取MongoDB集合实例
func (*MongoDB) GetCollection(dbName, collectionName string, appKeys ...string) *mongo.Collection {
	if len(appKeys) > 0 {
		return appMongoDB.GetCollection(appKeys[0], dbName, collectionName)
	}
	if baseMongoDB != nil {
		return dmongo.GetCollection(baseMongoDB, dbName, collectionName)
	}
	return nil
}

// Close 关闭MongoDB连接
func (*MongoDB) Close() {
	if baseMongoDB != nil {
		_ = baseMongoDB.Disconnect(context.TODO())
	}
}

const (
	GoDandelionMongoConnChange = "go_dandelion_mongo_conn_change"
)

// InitAppMongoDB 初始化应用MongoDB
// configFunc 定义获取相关应用MongoDB连接方法。用于在应
// 用MongoDB连接发生变更行为时，各服务自动获取最新的MongoDB
// 连接，不用额外进行重连刷新操作
//
// changeFunc 定义相关MongoDB变更时，业务需要执行的方法
// 发生变更时，进行回调处理。如，增加新的应用，服务需要
// 创建集合等操作
func InitAppMongoDB(configFunc dmongo.AppConfigFunc, changeFunc dmongo.AppChangeFunc) {
	// 使用这个必须有redis配置才行
	if config.Conf.Redis == nil {
		panic("使用该服务必须有redis配置才行。订阅服务依赖于redis")
	}
	appMongoDB = dmongo.InitAppMongoDB(configFunc, changeFunc, redis.Client)
	appMongoDB.Listener()
}

// AppMongoDBChange 用于上报应用MongoDB发生变更。
// 如中心服务器修改应用MongoDB，则需要上报，运
// 营服务订阅到消息后，自动刷新应用MongoDB连接
// 依赖于redis
func AppMongoDBChange(appKey string, changeType dmongo.ChangeType) error {
	if redis == nil {
		return errors.New("redis未初始化")
	}
	msg, err := jsoniter.MarshalToString(&dmongo.AppMongoMessage{
		AppKey:     appKey,
		ChangeType: changeType,
	})
	if err != nil {
		logger.Error(err)
		return nil
	}
	_, err = redis.Client.Execute(func(c zedis.Conn) (res interface{}, err error) {
		return c.Do("PUBLISH", GoDandelionMongoConnChange, msg)
	})
	return err
}
