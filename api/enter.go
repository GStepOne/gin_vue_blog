package api

import (
	"blog/gin/api/advert_api"
	"blog/gin/api/article_api"
	"blog/gin/api/category_api"
	"blog/gin/api/chat_api"
	"blog/gin/api/comment_api"
	"blog/gin/api/data_api"
	"blog/gin/api/digg_api"
	"blog/gin/api/images_api"
	"blog/gin/api/log_api"
	"blog/gin/api/menu_api"
	"blog/gin/api/message_api"
	"blog/gin/api/news_api"
	"blog/gin/api/role_api"
	"blog/gin/api/settings_api"
	"blog/gin/api/tag_api"
	"blog/gin/api/user_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi   images_api.ImagesApi
	AdvertApi   advert_api.AdvertApi
	MenuApi     menu_api.MenuApi
	LoginApi    user_api.UserApi
	TagApi      tag_api.TagApi
	MessageApi  message_api.MessageApi
	ArticleApi  article_api.ArticleApi
	DiggApi     digg_api.DiggApi
	CommentApi  comment_api.CommentApi
	NewsApi     news_api.NewsApi
	ChatApi     chat_api.ChatApi
	LogApi      log_api.LogApi
	DataApi     data_api.DataApi
	RoleApi     role_api.RoleApi
	CategoryApi category_api.CategoryApi
}

var ApiGroupApp = new(ApiGroup)
