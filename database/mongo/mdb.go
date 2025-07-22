package mongo

import (
	"sync"
	"time"

	timex "github.com/gly-hub/toolbox/time"
	jsoniter "github.com/json-iterator/go"
	"github.com/team-dandelion/go-dandelion/config"
	"github.com/team-dandelion/go-dandelion/database/redigo"
	"github.com/team-dandelion/go-dandelion/database/redigo/tools"
	"go.mongodb.org/mongo-driver/mongo"
)

// AppConfigFunc 用于获取需要初始化的应用MongoDB配置
// 初始化时，不会传递参数。应用数据库变更时，
// 会传递 appKey
type AppConfigFunc func(appKey ...string) (apps []AppConfig, err error)

// AppChangeFunc 数据连接变更监听触发方法
type AppChangeFunc func(appKey string, changeType ChangeType) (err error)

type AppMongoDB struct {
	clients       sync.Map
	appConfigFunc AppConfigFunc
	appChangeFunc AppChangeFunc
	redis         *redigo.Client
}

func InitAppMongoDB(configFunc AppConfigFunc, changeFunc AppChangeFunc, redis *redigo.Client) *AppMongoDB {
	appMongoDB := &AppMongoDB{
		clients:       sync.Map{},
		appConfigFunc: configFunc,
		appChangeFunc: changeFunc,
		redis:         redis,
	}

	apps, err := configFunc()
	if err != nil {
		panic(err)
	}

	for _, app := range apps {
		mongoConfig := &Config{
			MaxPoolSize:            uint64(100),            // 默认值
			MinPoolSize:            uint64(10),             // 默认值
			MaxConnIdleTime:        30 * time.Minute,       // 默认值
			ConnectTimeout:         10 * time.Second,       // 默认值
			SocketTimeout:          30 * time.Second,       // 默认值
			ServerSelectionTimeout: 30 * time.Second,       // 默认值
			LogLevel:               1,                      // 默认值
			SlowThreshold:          100 * time.Millisecond, // 默认值
			Hosts:                  app.MongoHosts,
			Database:               app.Database,
			Username:               app.Username,
			Password:               app.Password,
			AuthDB:                 app.AuthDB,
			ReplicaSet:             app.ReplicaSet,
		}

		// 如果配置存在则使用配置值
		if config.Conf.Mongo != nil {
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
			if config.Conf.Mongo.ReadPreference != "" {
				mongoConfig.ReadPreference = config.Conf.Mongo.ReadPreference
			}
		}

		appMongoDB.clients.Store(app.AppKey, NewConnection(mongoConfig))
	}
	return appMongoDB
}

func (client *AppMongoDB) GetClient(appKey string) *mongo.Client {
	if appKey == "" {
		return nil
	}
	if mongoClient, ok := client.clients.Load(appKey); ok {
		return mongoClient.(*mongo.Client)
	}
	return nil
}

func (client *AppMongoDB) GetDatabase(appKey, dbName string) *mongo.Database {
	mongoClient := client.GetClient(appKey)
	if mongoClient == nil {
		return nil
	}
	return GetDatabase(mongoClient, dbName)
}

func (client *AppMongoDB) GetCollection(appKey, dbName, collectionName string) *mongo.Collection {
	mongoClient := client.GetClient(appKey)
	if mongoClient == nil {
		return nil
	}
	return GetCollection(mongoClient, dbName, collectionName)
}

// Listener MongoDB连接变更监听
func (client *AppMongoDB) Listener() {
	// 注册订阅方法
	tools.RegisterFunc(GoDandelionMongoConnChange, func(pattern, channel, device string) {
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
			// 关闭现有连接
			if mongoClient, ok := client.clients.Load(msg.AppKey); ok {
				if mc, ok := mongoClient.(*mongo.Client); ok {
					_ = mc.Disconnect(nil)
				}
			}
			client.clients.Delete(msg.AppKey)
		}
	})
}

func (client *AppMongoDB) Connection(appKey string) error {
	if appKey == "" {
		return nil
	}

	// 关闭现有连接
	if mongoClient, ok := client.clients.Load(appKey); ok {
		if mc, ok := mongoClient.(*mongo.Client); ok {
			_ = mc.Disconnect(nil)
		}
	}

	// 重新获取数据库连接
	apps, err := client.appConfigFunc(appKey)
	if err != nil {
		panic(err)
	}

	for _, app := range apps {
		mongoConfig := &Config{
			MaxPoolSize:            uint64(100),            // 默认值
			MinPoolSize:            uint64(10),             // 默认值
			MaxConnIdleTime:        30 * time.Minute,       // 默认值
			ConnectTimeout:         10 * time.Second,       // 默认值
			SocketTimeout:          30 * time.Second,       // 默认值
			ServerSelectionTimeout: 30 * time.Second,       // 默认值
			LogLevel:               1,                      // 默认值
			SlowThreshold:          100 * time.Millisecond, // 默认值
			Hosts:                  app.MongoHosts,
			Database:               app.Database,
			Username:               app.Username,
			Password:               app.Password,
			AuthDB:                 app.AuthDB,
			ReplicaSet:             app.ReplicaSet,
		}

		// 如果配置存在则使用配置值
		if config.Conf.Mongo != nil {
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
			if config.Conf.Mongo.ReadPreference != "" {
				mongoConfig.ReadPreference = config.Conf.Mongo.ReadPreference
			}
		}

		client.clients.Store(app.AppKey, NewConnection(mongoConfig))
	}
	return nil
}

// ParseChangeData 解析变更消息
func ParseChangeData(msg string) AppMongoMessage {
	var data AppMongoMessage
	_ = jsoniter.UnmarshalFromString(msg, &data)
	return data
}
