package routers

import (
	"blog/gin/api"
)

func (router RouterGroup) ImagesRouter() {
	app := api.ApiGroupApp.ImagesApi
	router.POST("images", app.ImagesMultiUploadView) //多文件上传
	router.POST("image", app.ImagesSingleUploadView) //单文件上传
	router.GET("images", app.ImageListView)
	router.DELETE("images", app.ImageRemoveView)
	router.PUT("images", app.ImageListUpdate) //批量更新
	router.GET("image", app.ImageView)
}

