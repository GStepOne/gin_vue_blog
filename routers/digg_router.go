package routers

import (
	"blog/gin/api"
	"blog/gin/middleware"
)

func (router RouterGroup) DiggRouter() {
	app := api.ApiGroup{}.DiggApi
	router.POST("digg", middleware.JwtAuth(), app.DiggArticleView)
}
