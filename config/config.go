package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var CONF_INSTANCE *Conf

// 配置信息
type Conf struct {
	DbConf       `yaml:"db"`
	GenerateConf `yaml:"generate"`
}

type DbConf struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
	SslMode  string `yaml:"sslMode"`
	Timezone string `yaml:"timezone"`
}

type GenerateConf struct {
	Tables      []string `yaml:"tables"`
	ProjectDir  string   `yaml:"projectDir"`
	ModuleDir   string   `yaml:"moduleDir"`
	ModuleName  string   `yaml:"moduleName"`
	ProjectName string   `yaml:"projectName"`
}

func Setup(configFile string) {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(data, &CONF_INSTANCE)
	if err != nil {
		log.Fatal(err)
	}
}
