package routers

import (
	"blog/gin/api"
)

func (router RouterGroup) DataRouter() {
	app := api.ApiGroupApp.DataApi
	router.GET("data_seven", app.SevenLogin)
	router.GET("data_sum", app.DataSumView)
	router.GET("weather", app.DataWeather)
}
