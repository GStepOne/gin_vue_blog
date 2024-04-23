package category_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (CategoryApi) CategoryRemoveView(c *gin.Context) {
	//批量删除
	var query models.RemoveRequest
	err := c.ShouldBindJSON(&query)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var CategoryList []models.CategoryModel
	count := global.DB.Find(&CategoryList, query.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("分类不存在", c)
		return
	}
	global.DB.Delete(&CategoryList)
	res.OKWithMessage(fmt.Sprintf("共删除 %d 个分类", count), c)
}
