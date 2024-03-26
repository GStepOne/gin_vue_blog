package settings_api

import (
	"blog/gin/global"
	"blog/gin/models/res"
	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsEmailInfoView(c *gin.Context) {
	res.OKWithData(global.Config.Email, c)
}
