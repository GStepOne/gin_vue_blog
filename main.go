package main

import (
	"blog/gin/core"
	"blog/gin/global"
	"blog/gin/routers"
)

func main() {
	//初始化配置
	core.InitCoreConf()
	//初始化日志
	global.Log = core.InitLogger()
	//初始化数据库
	global.DB = core.InitGorm()
	//初始化路由
	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("程序gvb server 运行在 %s", addr)
	router.Run(addr)
}
