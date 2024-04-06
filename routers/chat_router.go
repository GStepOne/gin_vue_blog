package routers

import (
	"blog/gin/api"
)

func (router RouterGroup) ChatRouter() {
	app := api.ApiGroupApp.ChatApi
	router.GET("chat_group", app.ChatGroup)
	router.GET("chat_message_list", app.ChatListView)
}
