package routers

import (
	"blog/gin/api"
	"blog/gin/middleware"
)

func (router RouterGroup) ArticleRouter() {
	app := api.ApiGroup{}.ArticleApi
	router.POST("article", middleware.JwtAuth(), app.ArticleCreateView)
	router.GET("article", app.ArticleListView)
	router.GET("article/tags", app.ArticleTagListView)
	router.GET("article/:id", app.ArticleDetailView)
	router.GET("article/detail", app.ArticleDetailByTitle)
	router.GET("article/calendar", app.ArticleCalendarView)
	router.PUT("article", middleware.JwtAuth(), app.ArticleUpdateView)
	router.DELETE("article", middleware.JwtAuth(), app.ArticleRemoveView)
	router.POST("article/collects", middleware.JwtAuth(), app.ArticleCollectCreateView)
	router.GET("article/collects", middleware.JwtAuth(), app.ArticleCollectList)
	router.GET("article/comments", app.ArticleCommentList)
	router.GET("article/fulltext", app.FullTextSearch)
	router.DELETE("article/collects", middleware.JwtAuth(), app.ArticleCollBatchRemoveView)
}
