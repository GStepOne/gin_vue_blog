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
	router.GET("article/detail", middleware.JwtAuth(), app.ArticleDetailByTitle)
	router.GET("article/calendar", middleware.JwtAuth(), app.ArticleCalendarView)
}
