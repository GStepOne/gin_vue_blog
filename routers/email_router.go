package routers

import (
	"blog/gin/api"
	"blog/gin/middleware"
)

func (router RouterGroup) EmailRouter() {
	app := api.ApiGroupApp.LoginApi
	router.POST("/user_login", app.EmailLogin)
	router.GET("/user_list", middleware.JwtAuth(), app.UserListView)
	router.PUT("/user_role", middleware.JwtAuth(), app.UserUpdateRoleView)
	router.PUT("/user_password", middleware.JwtAuth(), app.UserUpdatePassword)
	router.POST("/user_logout", middleware.JwtAuth(), app.LogoutView)
}
