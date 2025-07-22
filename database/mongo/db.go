package mongo

import (
	"context"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewConnection(config *Config) *mongo.Client {
	if len(config.Hosts) == 0 {
		panic("未检测到MongoDB主机配置")
	}

	// 构建连接URI
	uri := buildMongoURI(config)

	// 创建客户端选项
	clientOptions := options.Client().ApplyURI(uri)

	// 设置连接池配置
	if config.MaxPoolSize > 0 {
		clientOptions.SetMaxPoolSize(config.MaxPoolSize)
	} else {
		clientOptions.SetMaxPoolSize(100)
	}

	if config.MinPoolSize > 0 {
		clientOptions.SetMinPoolSize(config.MinPoolSize)
	} else {
		clientOptions.SetMinPoolSize(10)
	}

	if config.MaxConnIdleTime > 0 {
		clientOptions.SetMaxConnIdleTime(config.MaxConnIdleTime)
	} else {
		clientOptions.SetMaxConnIdleTime(30 * time.Minute)
	}

	if config.ConnectTimeout > 0 {
		clientOptions.SetConnectTimeout(config.ConnectTimeout)
	} else {
		clientOptions.SetConnectTimeout(10 * time.Second)
	}

	if config.SocketTimeout > 0 {
		clientOptions.SetSocketTimeout(config.SocketTimeout)
	} else {
		clientOptions.SetSocketTimeout(30 * time.Second)
	}

	if config.ServerSelectionTimeout > 0 {
		clientOptions.SetServerSelectionTimeout(config.ServerSelectionTimeout)
	} else {
		clientOptions.SetServerSelectionTimeout(30 * time.Second)
	}

	// 设置读偏好
	if config.ReadPreference != "" {
		switch config.ReadPreference {
		case "primary":
			clientOptions.SetReadPreference(readpref.Primary())
		case "primaryPreferred":
			clientOptions.SetReadPreference(readpref.PrimaryPreferred())
		case "secondary":
			clientOptions.SetReadPreference(readpref.Secondary())
		case "secondaryPreferred":
			clientOptions.SetReadPreference(readpref.SecondaryPreferred())
		case "nearest":
			clientOptions.SetReadPreference(readpref.Nearest())
		}
	} else {
		// 默认使用 primary
		clientOptions.SetReadPreference(readpref.Primary())
	}

	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(fmt.Sprintf("MongoDB连接失败: %v", err))
	}

	// 测试连接
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(fmt.Sprintf("MongoDB连接测试失败: %v", err))
	}

	return client
}

func buildMongoURI(config *Config) string {
	var uri string

	// 构建认证信息
	var auth string
	if config.Username != "" && config.Password != "" {
		auth = fmt.Sprintf("%s:%s@", config.Username, config.Password)
	}

	// 构建主机列表
	hosts := strings.Join(config.Hosts, ",")

	// 构建基础URI
	uri = fmt.Sprintf("mongodb://%s%s/%s", auth, hosts, config.Database)

	// 添加参数
	var params []string

	if config.AuthDB != "" {
		params = append(params, fmt.Sprintf("authSource=%s", config.AuthDB))
	}

	if config.ReplicaSet != "" {
		params = append(params, fmt.Sprintf("replicaSet=%s", config.ReplicaSet))
	}

	if config.ReadPreference != "" {
		params = append(params, fmt.Sprintf("readPreference=%s", config.ReadPreference))
	}

	if len(params) > 0 {
		uri += "?" + strings.Join(params, "&")
	}

	return uri
}

// GetDatabase 获取数据库实例
func GetDatabase(client *mongo.Client, dbName string) *mongo.Database {
	return client.Database(dbName)
}

// GetCollection 获取集合实例
func GetCollection(client *mongo.Client, dbName, collectionName string) *mongo.Collection {
	return client.Database(dbName).Collection(collectionName)
}
