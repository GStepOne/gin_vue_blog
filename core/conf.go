package core

import (
	"blog/gin/config"
	"blog/gin/global"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

/*
*
读取yaml文件的配置
*/
func InitCoreConf() {
	const ConfigFile = "settings.yaml"
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)

	if err != nil {
		panic(fmt.Errorf("get yamlConf error:%s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatal("config unmarshal error:%s")
	}

	log.Println("config yamlFile load Init success.")

	//fmt.Println(c)
	//存放到全局变量下
	global.Config = c
}
