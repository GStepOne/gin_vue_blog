package routers

import (
	"blog/gin/api"
	"blog/gin/middleware"
)

func (router RouterGroup) NewsRouter() {
	app := api.ApiGroupApp.NewsApi
	router.GET("news", middleware.JwtAuth(), app.NewListView)
}
