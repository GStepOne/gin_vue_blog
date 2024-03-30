package routers

import "blog/gin/api"

func (router RouterGroup) MenuRouter() {
	app := api.ApiGroupApp.MenuApi
	router.POST("/menus", app.MenuCreateView)
	router.GET("/menus", app.MenuListView)
	router.GET("/menus/:id", app.MenuView)
	router.GET("/menus/names", app.MenuNameList)
	router.PUT("/menus/:id", app.MenuUpdateView)
	router.DELETE("/menus", app.MenuRemoveView)
}
