package application

import (
	"github.com/gly-hub/go-dandelion/config"
	dgorm "github.com/gly-hub/go-dandelion/database/gorm"
	"github.com/jinzhu/gorm"
)

var (
	wdb *gorm.DB
	rdb *gorm.DB
)

func initDb() {
	if config.Conf.WDB != nil {
		wdb = dgorm.NewConnection(&dgorm.Config{
			Type:     dgorm.DBType(config.Conf.WDB.Type),
			User:     config.Conf.WDB.User,
			Password: config.Conf.WDB.Password,
			Host:     config.Conf.WDB.Host,
			Port:     config.Conf.WDB.Port,
			Name:     config.Conf.WDB.Name,
		})
	}

	if config.Conf.RDB != nil {
		rdb = dgorm.NewConnection(&dgorm.Config{
			Type:     dgorm.DBType(config.Conf.RDB.Type),
			User:     config.Conf.RDB.User,
			Password: config.Conf.RDB.Password,
			Host:     config.Conf.RDB.Host,
			Port:     config.Conf.RDB.Port,
			Name:     config.Conf.RDB.Name,
		})
	}
	return
}

type DB struct {
}

// GetWDB 获取写库
func (d *DB) GetWDB() *gorm.DB {
	return wdb.Debug()
}

// GetRDB 获取读库
func (d *DB) GetRDB() *gorm.DB {
	if rdb == nil {
		return wdb
	}
	return rdb.Debug()
}
