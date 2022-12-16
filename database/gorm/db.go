package gorm

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type DBType string

const (
	Mysql    DBType = "mysql"
	Postgres DBType = "postgres"
)

type Config struct {
	Type     DBType `json:"type"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Name     string `json:"name"`
}

func NewConnection(config *Config) *gorm.DB {
	var dbUri string
	switch config.Type {
	case Mysql:
		dbUri = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
			config.User,
			config.Password,
			config.Host,
			config.Port,
			config.Name)
	case Postgres:
		dbUri = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
			config.Host,
			config.Port,
			config.User,
			config.Name,
			config.Password)
	default:
		panic("未检测到支持的数据库类型")
	}
	conn, err := gorm.Open(string(config.Type), dbUri)
	if err != nil {
		log.Print(err.Error())
	}
	conn.DB().SetMaxIdleConns(10)                   //最大空闲连接数
	conn.DB().SetMaxOpenConns(30)                   //最大连接数
	conn.DB().SetConnMaxLifetime(time.Second * 300) //设置连接空闲超时
	return conn
}
