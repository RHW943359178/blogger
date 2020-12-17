package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Conf struct {
	Host      string `yaml:"host"`
	Port      string `yaml:"port"`
	MysqlAddr string `yaml:"mysql_addr"`
	RedisAddr string `yaml:"redis_addr"`
}

func GetConf() (conf *Conf) {
	//	获取项目当前的绝对路径
	dir, err := os.Getwd()
	//	替换转义符
	path := strings.Replace(dir, "\\", "/", -1)
	if err != nil {
		log.Println(err.Error())
	}
	file, err := ioutil.ReadFile(path + "/config/conf.yaml")
	if err != nil {
		log.Println(err.Error())
	}
	err = yaml.Unmarshal(file, conf)
	if err != nil {
		log.Println(err.Error())
	}
	return conf
}
