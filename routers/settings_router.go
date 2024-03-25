package routers

import (
	"blog/gin/api"
	"github.com/gin-gonic/gin"
)

func (router RouterGroup) SettingRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("/", settingsApi.SettingsInfoView)
}

func (RouterGroup) SettingRouter1(router *gin.Engine) {
	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("/", settingsApi.SettingsInfoView)
}
