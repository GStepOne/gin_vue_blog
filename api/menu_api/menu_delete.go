package menu_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (MenuApi) MenuRemoveView(c *gin.Context) {
	//批量删除
	var query models.RemoveRequest

	err := c.ShouldBindJSON(&query)

	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		fmt.Println(err)
		return
	}
	var menuList []models.MenuModel
	count := global.DB.Find(&menuList, query.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("文件不存在", c)
		return
	}

	err = global.DB.Transaction(func(tx *gorm.DB) error {
		err = global.DB.Model(&menuList).Association("Banners").Clear()
		if err != nil {
			global.Log.Error()
			return err
		}

		err = global.DB.Delete(&menuList).Error
		if err != nil {
			global.Log.Error()
			return err
		}
		return nil
	})

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("删除菜单失败", c)
	}

	res.OKWithMessage(fmt.Sprintf("共删除 %d 个菜单", count), c)
}
