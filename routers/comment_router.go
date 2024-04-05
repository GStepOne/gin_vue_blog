package routers

import (
	"blog/gin/api"
	"blog/gin/middleware"
)

func (router RouterGroup) CommentRouter() {
	app := api.ApiGroupApp.CommentApi
	router.POST("comment", middleware.JwtAuth(), app.CommentCreateView)
	router.GET("comment", middleware.JwtAuth(), app.CommentListView)
}
