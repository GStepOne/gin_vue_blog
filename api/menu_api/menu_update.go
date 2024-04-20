package menu_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"fmt"
	structs2 "github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

func (MenuApi) MenuUpdateView(c *gin.Context) {

	var menuRequest MenuRequest
	err := c.ShouldBindJSON(&menuRequest)
	if err != nil {
		res.FailWithError(err, menuRequest, c)
		return
	}
	//先把之前的banner清空

	id := c.Param("id")

	var menuModel models.MenuModel
	err = global.DB.Take(&menuModel, id).Error
	if err != nil {
		res.FailWithMessage("菜单不存在", c)
		return
	}

	global.DB.Debug().Model(&menuModel).Association("Banners").Clear()

	if len(menuRequest.ImageSortList) > 0 {
		//操作第三张表
		var bannerList []models.MenuBannerModel
		for _, sort := range menuRequest.ImageSortList {
			bannerList = append(bannerList, models.MenuBannerModel{
				MenuId:   menuModel.ID,
				BannerId: sort.ImageId,
				Sort:     sort.Sort,
			})
		}
		err = global.DB.Create(&bannerList).Error

		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("创建菜单图片失败", c)
			return
		}
		//普通更新 menu表的其他字段

		maps := structs2.Map(&menuRequest)
		fmt.Println(maps)
		err = global.DB.Model(&menuModel).Updates(&maps).Error

		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("修改菜单失败", c)
			return
		}
	}

	res.OKWithMessage("修改菜单成功", c)

}
