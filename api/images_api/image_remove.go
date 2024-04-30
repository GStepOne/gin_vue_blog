package images_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (ImagesApi) ImageRemoveView(c *gin.Context) {
	//批量删除
	var query models.RemoveRequest
	err := c.ShouldBindJSON(&query)
	fmt.Println("id_list1", query.IDList)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		fmt.Println(err)
		return
	}
	var imageList []models.BannerModel
	fmt.Println("id_list", query.IDList)
	count := global.DB.Debug().Where("id IN (?)", query.IDList).Find(&imageList).RowsAffected
	if count == 0 {
		res.FailWithMessage("文件不存在", c)
		return
	}
	global.DB.Debug().Delete(&imageList)

	res.OKWithMessage(fmt.Sprintf("共删除 %d 张图片", count), c)
}
