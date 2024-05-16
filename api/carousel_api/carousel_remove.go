package carousel_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (CarouselApi) CarouselRemoveView(c *gin.Context) {
	//批量删除
	var query models.RemoveRequest
	err := c.ShouldBindJSON(&query)

	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		fmt.Println(err)
		return
	}
	var imageList []models.CarouselModel

	count := global.DB.Debug().Where("id IN (?)", query.IDList).Find(&imageList).RowsAffected
	if count == 0 {
		res.FailWithMessage("文件不存在", c)
		return
	}
	global.DB.Debug().Delete(&imageList)

	res.OKWithMessage(fmt.Sprintf("共删除 %d 张轮播图片", count), c)
}
