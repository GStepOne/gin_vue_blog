package routers

import (
	"blog/gin/global"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	apiRouterGroup := router.Group("api")
	routerGroupApp := RouterGroup{apiRouterGroup}
	routerGroupApp.SettingRouter()
	routerGroupApp.ImagesRouter()
	routerGroupApp.AdvertRouter()
	routerGroupApp.MenuRouter()
	routerGroupApp.UserRouter()

	return router
}
