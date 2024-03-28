package api

import (
	"blog/gin/api/advert_api"
	"blog/gin/api/images_api"
	"blog/gin/api/settings_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi   images_api.ImagesApi
	AdvertApi   advert_api.AdvertApi
}

var ApiGroupApp = new(ApiGroup)
