package api

import (
	"blog/gin/api/advert_api"
	"blog/gin/api/images_api"
	"blog/gin/api/menu_api"
	"blog/gin/api/settings_api"
	"blog/gin/api/user_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi   images_api.ImagesApi
	AdvertApi   advert_api.AdvertApi
	MenuApi     menu_api.MenuApi
	LoginApi    user_api.UserApi
}

var ApiGroupApp = new(ApiGroup)
