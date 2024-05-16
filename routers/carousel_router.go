package routers

import (
	"blog/gin/api"
	"blog/gin/middleware"
)

func (router RouterGroup) CarouselsRouter() {
	app := api.ApiGroupApp.CarouselApi
	router.POST("carousel", middleware.JwtAuth(), app.CarouselsSingleUploadView) //单文件上传
	router.GET("carousels", middleware.JwtAuth(), app.CarouselListView)
	router.DELETE("carousels", middleware.JwtAuth(), app.CarouselRemoveView)
	router.PUT("carousels", middleware.JwtAuth(), app.CarouselListUpdate) //批量更新
	router.GET("carousel", middleware.JwtAuth(), app.CarouselView)
}
