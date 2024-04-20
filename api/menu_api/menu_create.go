package menu_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/ctype"
	"blog/gin/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

type BannerSort struct {
	ImageId uint `json:"image_id"`
	Sort    int  `json:"sort"`
}

type MenuRequest struct {
	Title         string       `json:"title"  binding:"required" msg:"请完善菜单名称" structs:"title"`
	TitleEn       string       `json:"title_en"  msg:"请完善菜单英文名称" structs:"title_en"`
	Slogan        string       `json:"slogan" structs:"slogan"`
	Path          string       `json:"path" structs:"path"`
	Abstract      ctype.Array  `json:"abstract" structs:"abstract"`
	AbstractTime  int          `json:"abstract_time" structs:"abstract_time"`
	BannerTime    int          `json:"banner_time" structs:"banner_time"` //菜单图片的切换时间 为0 不切换
	Sort          int          `json:"sort" binding:"required" msg:"请输入菜单序号" structs:"sort"`
	ImageSortList []BannerSort `json:"image_sort_list" structs:"-"`
}

func (MenuApi) MenuCreateView(c *gin.Context) {

	var menuRequest MenuRequest
	err := c.ShouldBindJSON(&menuRequest)
	if err != nil {
		res.FailWithError(err, menuRequest, c)
		return
	}

	var menuModel models.MenuModel

	err = global.DB.Take(&menuModel, "title=? or path=?", menuRequest.Title, menuRequest.Path).Error
	if err == nil {
		res.FailWithMessage("菜单已经存在", c)
		return
	}

	menuModel = models.MenuModel{
		Title:        menuRequest.Title,
		TitleEn:      menuRequest.TitleEn,
		Slogan:       menuRequest.Slogan,
		Path:         menuRequest.Path,
		Abstract:     menuRequest.Abstract,
		AbstractTime: menuRequest.AbstractTime,
		BannerTime:   menuRequest.BannerTime,
		Sort:         menuRequest.Sort,
	}

	err = global.DB.Create(&menuModel).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("菜单创建失败", c)
		return
	}

	if len(menuRequest.ImageSortList) == 0 {
		res.OKWithMessage("菜单创建成功", c)
		return
	}

	var menuBannerList []models.MenuBannerModel

	for _, sort := range menuRequest.ImageSortList {
		menuBannerList = append(menuBannerList, models.MenuBannerModel{
			MenuId:   menuModel.ID,
			BannerId: sort.ImageId,
			Sort:     sort.Sort,
		})
	}

	fmt.Println(menuBannerList)
	global.DB.Create(&menuBannerList)

	res.OKWithMessage("菜单创建成功", c)

}
