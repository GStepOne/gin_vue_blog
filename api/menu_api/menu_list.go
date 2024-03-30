package menu_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"github.com/gin-gonic/gin"
)

type Banner struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}

type MenuResponse struct {
	models.MenuModel
	Banners []Banner
}

func (MenuApi) MenuListView(c *gin.Context) {
	//var page models.PageView
	var menuIdList []uint

	var menuList []models.MenuModel

	//先把menu表的id读取出来，放入menuIdList，menu的数据放入menuList
	err := global.DB.Order("sort desc").Find(&menuList).Select("id").Scan(&menuIdList).Error

	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	//查链接表
	var menuBanners []models.MenuBannerModel
	//根据上面读取的menu_id读取menu_banner表的数据
	global.DB.Debug().Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id in ?", menuIdList)

	var menus []MenuResponse

	for _, model := range menuList {
		//model是一个菜单
		//var banners []Banner //只声明 不赋值，最后就等于nil，
		banners := []Banner{}
		//banners := make([]Banner,0)  //解决null值问题
		for _, banner := range menuBanners {
			if model.ID != banner.MenuId {
				continue
			}

			banners = append(banners, Banner{
				ID:   banner.ImageId,
				Path: banner.ImageModel.Path,
			})
		}

		menus = append(menus, MenuResponse{
			MenuModel: model,
			Banners:   banners,
		})
	}

	//fmt.Println(&menuList)
	//fmt.Println(&menuIdList)

	res.OKWithData(menus, c)
}
