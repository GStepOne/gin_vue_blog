package menu_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"github.com/gin-gonic/gin"
)

func (MenuApi) MenuView(c *gin.Context) {

	//先查菜单
	id := c.Param("id")
	var menuModel models.MenuModel
	//先把menu表的id读取出来，放入menuIdList，menu的数据放入menuList
	err := global.DB.Take(&menuModel, id).Error

	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	//查链接表
	var menuBanners []models.MenuBannerModel
	//根据上面读取的menu_id读取menu_banner表的数据
	global.DB.Debug().Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id = ?", id)

	//var menuResponse MenuResponse
	//model是一个菜单
	//var banners []Banner //只声明 不赋值，最后就等于nil，
	banners := []Banner{}
	for _, banner := range menuBanners {
		if menuModel.ID != banner.MenuId {
			continue
		}
		banners = append(banners, Banner{
			ID:   banner.ImageId,
			Path: banner.ImageModel.Path,
		})
	}
	menuResponse := MenuResponse{
		MenuModel: menuModel,
		Banners:   banners,
	}

	res.OKWithData(menuResponse, c)
}
