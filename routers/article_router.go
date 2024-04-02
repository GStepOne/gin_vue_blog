package routers

import (
	"blog/gin/api"
	"blog/gin/middleware"
)

func (router RouterGroup) ArticleRouter() {
	app := api.ApiGroup{}.ArticleApi
	router.POST("article", middleware.JwtAuth(), app.ArticleCreateView)
	router.GET("article", middleware.JwtAuth(), app.ArticleListView)
	router.GET("article/:id", middleware.JwtAuth(), app.ArticleDetailView)
}
