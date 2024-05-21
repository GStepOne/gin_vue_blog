package routers

import (
	"blog/gin/api"
)

func (router RouterGroup) CarouselsRouter() {
	app := api.ApiGroupApp.CarouselApi
	router.POST("carousel", app.CarouselsSingleUploadView) //单文件上传
	router.GET("carousels", app.CarouselListView)
	router.DELETE("carousels", app.CarouselRemoveView)
	router.PUT("carousels", app.CarouselListUpdate) //批量更新
	router.GET("carousel", app.CarouselView)
}
