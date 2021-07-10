package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type DBConf struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Cluster  string `yaml:"cluster"`
	Host     int64  `yaml:"host"`
}

var DBConfig = &DBConf{}

func init() {
	yamlFile, err := ioutil.ReadFile("/Users/suchenghao/mygit/scp/service/config/config.yaml")
	if err != nil {
		log.Fatalf("get yaml config err: %v ", err)
		return
	}

	err = yaml.Unmarshal(yamlFile, DBConfig)
	if err != nil {
		log.Fatalf("unmarshal dbconf err: %v", err)
	}
	return
}
