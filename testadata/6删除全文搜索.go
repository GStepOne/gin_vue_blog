package main

import (
	"blog/gin/core"
	"blog/gin/global"
	"blog/gin/service/es_ser"
)

func main() {
	core.InitCoreConf()
	core.InitLogger()
	global.EsClient = core.EsConnect()
	es_ser.DeleteFullTextByArticleId("taUNp44BReI8ZXCFVxV3")
}
