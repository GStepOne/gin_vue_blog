package main

import (
	"blog/gin/core"
	_ "blog/gin/docs"
	"blog/gin/flag"
	"blog/gin/global"
	"blog/gin/routers"
	"fmt"
)

// @title gvb_server API文档
// @version 1.0
// @description gvb_server API文档
// @host 127.0.0.1:8081
// @Basepath /
func main() {
	//初始化配置
	core.InitCoreConf()
	//初始化日志
	global.Log = core.InitLogger()
	//初始化数据库
	global.DB = core.InitGorm()

	global.Redis = core.ConnectRedis()

	//建表
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
	}

	if flag.IsCreateUser(option) == "user" {
		flag.SwitchOption(option)
	}
	//初始化路由
	router := routers.InitRouter()

	//routes := router.Routes()
	//for _, route := range routes {
	//	fmt.Printf("路由方法: %s, 路径: %s, 处理函数: %p\n", route.Method, route.Path, route.HandlerFunc)
	//}
	addr := global.Config.System.Addr()
	global.Log.Infof("程序gvb server 运行在 %s", addr)

	err := router.Run(addr)
	if err != nil {
		fmt.Println("路由启动失败", err) //tmd 端口被用了，错误日志没有打印出来
	}
}
