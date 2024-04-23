package routers

import (
	"blog/gin/api"
)

func (router RouterGroup) CategoryRouter() {
	app := api.ApiGroupApp.CategoryApi
	router.POST("category", app.CategoryCreate)
	router.GET("category", app.CategoryListView)
	router.GET("categories", app.CategoryListViewNoPage) //不带分页的
	router.PUT("category/:id", app.CategoryUpdateView)
	router.DELETE("category", app.CategoryRemoveView)
}
