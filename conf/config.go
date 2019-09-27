package conf

import (
	"bytes"
	"echo/utils"
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/labstack/gommon/log"
)

type MysqlConf struct {
	Ip       string `toml:"ip"`
	Port     string `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	DB       string `toml:"db"`
}

type RedisConf struct {
	Ip       string `toml:"ip"`
	Port     string `toml:"port"`
	Password string `toml:"password"`
	DB       int    `toml:"db"`
}

type Conf struct {
	Title     string                  `toml:"title"`
	Env       string                  `toml:"env"`
	Version   string                  `toml:"version"`
	Server    map[string]ServerConfig `toml:"server"`
	Port      string                  `toml:"port"`
	JwtKey    string                  `toml:"jwt_key"`
	WhiteList []string                `toml:"white_list"`
}

type ServerConfig struct {
	ExampleSQL *MysqlConf `toml:"example_sql"`
	CloudSQL   *MysqlConf `toml:"cloud_sql"`
}

var (
	conf *Conf
)

func LoadConfig(configPath string) {
	_, err := toml.DecodeFile(configPath, &conf)
	if err != nil {
		panic(err)
	}
	fmt.Println("read tomlConfig: ", configPath)
	fmt.Println(conf.String())
}

func (c *Conf) String() string {
	b, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		panic(err)
	}
	return out.String()
}

func GetConfig() *Conf {
	if conf == nil {
		reloadConfig()
	}
	return conf
}

func reloadConfig() {
	// 单元测试的时候需要导入config
	configPath := []string{"../conf/config.toml", "../../conf/config.toml", "../../../conf/config.toml"}
	for _, path := range configPath {
		if utils.PathExists(path) {
			if conf == nil {
				LoadConfig(path)
			}
		}
	}
	if conf == nil {
		log.Error("can't find config path")
		panic("can't find config path")
	}
}

func (m *MysqlConf) GetUrl() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local", m.Username,
		m.Password, m.Ip, m.Port, m.DB)
}
