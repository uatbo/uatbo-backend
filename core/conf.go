package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"uatbo-backend/config"
	"uatbo-backend/global"
)

// InitConf 读取yaml文件的配置
func InitConf() {
	const ConfigFile = "settings.yaml"
	c := &config.Config{}
	yamlConf, err := os.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error: %s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		// 配置的log还没有加载出来，只能用本身的log
		log.Fatalf("config Init Unmarshal: %v", err)
	}
	log.Println("config yamlFile load Init success")
	// 存储到全局中
	global.Config = c
}
