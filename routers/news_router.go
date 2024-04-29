package routers

import (
	"blog/gin/api"
)

func (router RouterGroup) NewsRouter() {
	app := api.ApiGroupApp.NewsApi
	router.GET("news", app.NewListView) // middleware.JwtAuth(),
}
