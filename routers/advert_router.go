package routers

import (
	"blog/gin/api"
)

func (router RouterGroup) AdvertRouter() {
	app := api.ApiGroupApp.AdvertApi
	router.POST("advert", app.AdvertCreate)
	router.GET("advert", app.AdvertListView)
	router.PUT("advert/:id", app.AdvertUpdateView)
	router.DELETE("advert", app.AdvertRemoveView)
}
