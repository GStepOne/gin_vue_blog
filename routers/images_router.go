package routers

import (
	"blog/gin/api"
)

func (router RouterGroup) ImagesRouter() {
	app := api.ApiGroupApp.ImagesApi
	router.POST("image", app.ImageUploadView)
	router.POST("images", app.ImagesMultiUploadView)
	router.GET("images", app.ImageListView)
	router.DELETE("images", app.ImageRemoveView)
	router.PUT("images", app.ImageListUpdate)
}
