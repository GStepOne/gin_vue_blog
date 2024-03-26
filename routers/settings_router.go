package routers

import (
	"blog/gin/api"
	"github.com/gin-gonic/gin"
)

func (router RouterGroup) SettingRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("/settings/:name", settingsApi.SettingsInfoView)
	router.PUT("/settings/:name", settingsApi.SettingsInfoUpdateView)

	//router.GET("/settings/email", settingsApi.SettingsEmailInfoView)
	//router.PUT("/settings/email", settingsApi.SettingsEmailUpdateView)
}

func (RouterGroup) SettingRouter1(router *gin.Engine) {
	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("/", settingsApi.SettingsInfoView)
}
