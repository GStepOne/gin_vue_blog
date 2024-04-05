package routers

import (
	"blog/gin/api"
	"blog/gin/middleware"
)

func (router RouterGroup) ArticleRouter() {
	app := api.ApiGroup{}.ArticleApi
	router.POST("article", middleware.JwtAuth(), app.ArticleCreateView)
	router.GET("article", middleware.JwtAuth(), app.ArticleListView)
	router.GET("article/tags", middleware.JwtAuth(), app.ArticleTagListView)
	router.GET("article/:id", middleware.JwtAuth(), app.ArticleDetailView)
	router.GET("article/detail", middleware.JwtAuth(), app.ArticleDetailByTitle)
	router.GET("article/calendar", middleware.JwtAuth(), app.ArticleCalendarView)
	router.PUT("article", middleware.JwtAuth(), app.ArticleUpdateView)
	router.DELETE("article", middleware.JwtAuth(), app.ArticleRemoveView)
	router.POST("article/collects", middleware.JwtAuth(), app.ArticleCollectCreateView)
	router.GET("article/collects", middleware.JwtAuth(), app.ArticleCollectList)
	router.GET("article/fulltext", app.FullTextSearch)
	router.DELETE("article/collects", middleware.JwtAuth(), app.ArticleCollBatchRemoveView)
}
