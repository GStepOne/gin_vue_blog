package routers

import (
	"blog/gin/api"
)

func (router RouterGroup) TagRouter() {
	app := api.ApiGroupApp.TagApi
	router.POST("tag", app.TagCreate)
	router.GET("tag", app.TagListView)
	router.GET("tags", app.TagListViewNoPage)
	router.PUT("tag/:id", app.TagUpdateView)
	router.DELETE("tag", app.TagRemoveView)
}
