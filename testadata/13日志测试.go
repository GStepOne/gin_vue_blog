package main

import (
	"blog/gin/core"
	"blog/gin/plugins/logstash"
	"fmt"
)

func main() {
	core.InitCoreConf()
	core.InitLogger()
	core.InitGorm()
	log := logstash.New("127.0.01", "ddd")
	log.Debug(fmt.Sprintf("%s你好啊", "feng"))
}
