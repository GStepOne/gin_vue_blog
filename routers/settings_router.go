package routers

import (
	"blog/gin/api"
	"github.com/gin-gonic/gin"
)

func (router RouterGroup) SettingRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("/settings", settingsApi.SettingsInfoView)
	router.PUT("/settings", settingsApi.SettingsInfoUpdateView)
}

func (RouterGroup) SettingRouter1(router *gin.Engine) {
	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("/", settingsApi.SettingsInfoView)
}
