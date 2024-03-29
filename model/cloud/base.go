package cloud

import (
	"echo/client"
	"echo/conf"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
)

var db *gorm.DB

func LoadCloudSQL() {
	config := conf.GetConfig()
	server := config.Server[config.Env]
	//todo: 通过反射获取mysql数据库类型，dbs 可存放多个mysql 客户端
	url := server.CloudSQL.GetUrl()
	db = client.GetMysqlClient(url)
	autoMigrate()
	log.Info("init cloud mysql success")
}

func autoMigrate() {
	db.AutoMigrate(&Host{})
}
