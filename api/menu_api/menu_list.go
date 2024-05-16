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

type keyRequest struct {
	Key string `json:"key" query:"key" form:"key"`
}

type MenuResponse struct {
	models.MenuModel
	Banners []Banner `json:"banners"`
}

func (MenuApi) MenuListView(c *gin.Context) {
	//var page models.PageView
	var menuIdList []uint

	var menuList []models.MenuModel

	var menuListRequest keyRequest

	err := c.ShouldBindQuery(&menuListRequest)

	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		global.Log.Error(err.Error())
		return
	}

	if menuListRequest.Key != "" {
		title := menuListRequest.Key
		// 先把menu表的id读取出来，放入menuIdList，menu的数据放入menuList
		err = global.DB.Order("sort desc").Where("title LIKE ?", "%"+title+"%").Find(&menuList).Select("id").Scan(&menuIdList).Error
	} else {
		// 如果menuListRequest.Key为空，则不使用like查询，保持原有逻辑
		err = global.DB.Order("sort desc").Find(&menuList).Select("id").Scan(&menuIdList).Error
	}
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	//查链接表
	var menuBanners []models.MenuBannerModel
	//根据上面读取的menu_id读取menu_banner表的数据
	global.DB.Debug().Preload("BannerModel").Order("sort desc").
		Find(&menuBanners, "menu_id in ?", menuIdList)

	var menus = make([]MenuResponse, 0)

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
				ID:   banner.BannerId,
				Path: banner.BannerModel.Path,
			})
		}

		menus = append(menus, MenuResponse{
			MenuModel: model,
			Banners:   banners,
		})
	}

	res.OkWithList(menus, int64(len(menus)), c)
}
