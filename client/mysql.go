package client

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/gommon/log"
)

func GetMysqlClient(databaseUrl string) *gorm.DB {

	db, err := gorm.Open("mysql", databaseUrl)
	if err != nil {
		log.Fatalf("connect mysql: %s fail, err: %v", databaseUrl, err)
		panic(err)
	}
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	return db
}
