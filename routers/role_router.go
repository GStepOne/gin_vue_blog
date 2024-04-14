package routers

import (
	"blog/gin/api"
)

func (router RouterGroup) RoleRouter() {
	roleApi := api.ApiGroupApp.RoleApi
	router.GET("/role", roleApi.RoleIDListView)
}
