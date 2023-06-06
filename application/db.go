package application

import (
	"errors"
	"github.com/gly-hub/go-dandelion/config"
	dgorm "github.com/gly-hub/go-dandelion/database/gorm"
	"github.com/gly-hub/go-dandelion/logger"
	timex "github.com/gly-hub/toolbox/time"
	zedis "github.com/gomodule/redigo/redis"
	jsoniter "github.com/json-iterator/go"
	"gorm.io/gorm"
)

var (
	baseDB *gorm.DB
	appDB  *dgorm.AppDB
)

func initDb() {
	if config.Conf.DB != nil && config.Conf.DB.Master != nil {
		var slaves []*dgorm.Slave
		if config.Conf.DB.Slaves != nil {
			for _, slave := range config.Conf.DB.Slaves {
				slaves = append(slaves, &dgorm.Slave{
					Host:     slave.Host,
					Port:     slave.Port,
					User:     slave.User,
					Password: slave.Password,
					Database: slave.Database,
				})
			}
		}

		baseDB = dgorm.NewConnection(&dgorm.Config{
			DBType:        config.Conf.DB.DBType,
			MaxOpenConn:   config.Conf.DB.MaxOpenConn,
			MaxIdleConn:   config.Conf.DB.MaxIdleConn,
			MaxLifeTime:   config.Conf.DB.MaxLifeTime,
			MaxIdleTime:   config.Conf.DB.MaxIdleTime,
			Level:         config.Conf.DB.Level,
			SlowThreshold: timex.ParseDuration(config.Conf.DB.SlowThreshold),
			Master: &dgorm.Master{
				Host:     config.Conf.DB.Master.Host,
				Port:     config.Conf.DB.Master.Port,
				User:     config.Conf.DB.Master.User,
				Password: config.Conf.DB.Master.Password,
				Database: config.Conf.DB.Master.Database,
			},
			Slaves: slaves,
		})
	}
	return
}

type DB struct {
}

// GetDB 获取数据库连接
func (*DB) GetDB(appKeys ...string) *gorm.DB {
	if len(appKeys) > 0 {
		return appDB.GetDB(appKeys[0])
	}
	return baseDB
}

const (
	GoDandelionConnChange = "go_dandelion_conn_change"
)

// InitAppDB 初始化应用数据库
// configFunc 定义获取相关应用数据库连接方法。用于在应
// 用数据库连接发生变更行为时，各服务自动获取最新的数据
// 库连接，不用额外进行重连刷新操作
//
// changeFunc 定义相关数据库变更时，业务需要执行的方法
// 发生变更时，进行回调处理。如，增加新的应用，服务需要
// 创建表单等操作
func InitAppDB(configFunc dgorm.AppConfigFunc, changeFunc dgorm.AppChangeFunc) {
	// 使用这个必须有redis配置才行
	if config.Conf.Redis == nil {
		panic("使用该服务必须有redis配置才行。订阅服务依赖于redis")
	}
	appDB = dgorm.InitAppDB(configFunc, changeFunc, redis.Client)
	appDB.Listener()
}

// AppDBChange 用于上报应用数据库发生变更。
// 如中心服务器修改应用数据库，则需要上报，运
// 营服务订阅到消息后，自动刷新应用数据库连接
// 依赖于redis
func AppDBChange(appKey string, changeType dgorm.ChangeType) error {
	if redis == nil {
		return errors.New("redis未初始化")
	}
	msg, err := jsoniter.MarshalToString(&dgorm.AppDBMessage{
		AppKey:     appKey,
		ChangeType: changeType,
	})
	if err != nil {
		logger.Error(err)
		return nil
	}
	_, err = redis.Client.Execute(func(c zedis.Conn) (res interface{}, err error) {
		return c.Do("PUBLISH", GoDandelionConnChange, msg)
	})
	return err
}
