package routers

import (
	"blog/gin/api"
)

func (router RouterGroup) AdvertRouter() {
	app := api.ApiGroupApp.AdvertApi
	router.POST("advert", app.AdvertCreate)
	router.GET("advert", app.AdvertListView)
	//router.GET("adverts", app.AdvertCreateView)
	//router.DELETE("adverts", app.AdvertCreateView)
	//router.PUT("adverts", app.AdvertCreateView)
}
