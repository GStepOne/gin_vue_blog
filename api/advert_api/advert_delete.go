package advert_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

// AdvertRemoveView 广告删除
// @Tags 广告删除
// @Summary 广告删除
// @Description 广告删除
// @Param data body models.RemoveRequest false "广告的id列表"
// @Router /api/advert [delete]
// @Produce json
// @Success 200 {object} res.Response{data=string}
func (AdvertApi) AdvertRemoveView(c *gin.Context) {
	//批量删除
	var query models.RemoveRequest

	err := c.ShouldBindJSON(&query)

	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		fmt.Println(err)
		return
	}
	var imageList []models.AdvertModel
	count := global.DB.Find(&imageList, query.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("文件不存在", c)
		return
	}
	global.DB.Delete(&imageList)
	res.OKWithMessage(fmt.Sprintf("共删除 %d 个广告", count), c)
	//绑一个钩子函数

}
