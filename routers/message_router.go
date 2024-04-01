package routers

import (
	"blog/gin/api"
	"blog/gin/middleware"
)

func (router RouterGroup) MessageRouter() {
	app := api.ApiGroupApp.MessageApi
	router.POST("/message", middleware.JwtAuth(), app.MessageCreateView) //添加消息
	router.GET("/message_all", app.MessageListAllView)                   //添加消息
	router.GET("/message", middleware.JwtAuth(), app.MessageListView)    //消息列表
	router.GET("/messages", middleware.JwtAuth(), app.MessageRecordView) //获取消息记录
}
