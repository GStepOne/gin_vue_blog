package routers

import (
	"blog/gin/api"
)

func (router RouterGroup) LogRouter() {
	app := api.ApiGroupApp.LogApi
	router.GET("log", app.LogListView)
	router.PUT("log", app.LogReadView)
	router.DELETE("log", app.LogRemoveListView)
}
