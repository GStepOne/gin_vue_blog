package routers

import (
	"blog/gin/api"
	"blog/gin/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

func (router RouterGroup) UserRouter() {
	app := api.ApiGroupApp.LoginApi
	var store = cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("sessionid", store))
	router.POST("/user_login", app.EmailLogin)
	router.GET("/user_list", middleware.JwtAuth(), app.UserListView)
	router.PUT("/user_role", middleware.JwtAuth(), app.UserUpdateRoleView)
	router.PUT("/user_password", middleware.JwtAuth(), app.UserUpdatePassword)
	router.POST("/user_logout", middleware.JwtAuth(), app.LogoutView)
	router.DELETE("/user_delete", middleware.JwtAuth(), app.UserRemoveView)
	router.POST("/user_bind_mail", middleware.JwtAuth(), app.UserBindEmailView)
	router.POST("/login", app.QQLoginView)
}
