package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"service/config"
)

type DB struct {
	Name string
}

func (d *DB) GetDB() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBConfig.User, config.DBConfig.Password, config.DBConfig.Cluster, config.DBConfig.Host, d.Name)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return
}
