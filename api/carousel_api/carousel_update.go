package carousel_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

type CarouselUpdateRequest struct {
	ID   uint   `query:"id" json:"id" binding:"required" msg:"请选择文件id""`
	Name string `query:"name" json:"name" binding:"required" msg:"请输入文件名称"`
}

func (CarouselApi) CarouselListUpdate(c *gin.Context) {
	var query CarouselUpdateRequest

	err := c.ShouldBindJSON(&query)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		fmt.Println(err)
		return
	}

	var imageModel models.CarouselModel
	err = global.DB.Take(&imageModel, query.ID).Error
	if err != nil {
		res.FailWithMessage("图片不存在", c)
		return
	}

	err = global.DB.Model(&imageModel).Update("name", query.Name).Error
	if err != nil {
		res.FailWithMessage("更新失败", c)
		return
	}

	res.OKWithMessage("图片名称修改成功", c)

}
