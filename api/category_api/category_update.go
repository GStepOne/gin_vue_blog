package category_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

func (CategoryApi) CategoryUpdateView(c *gin.Context) {
	var request CategoryRequest

	id := c.Param("id")

	fmt.Println(id)
	err := c.ShouldBindJSON(&request)

	if err != nil {
		res.FailWithError(err, &request, c)
		return
	}

	var tag models.CategoryModel

	err = global.DB.Debug().Take(&tag, id).Error
	if err != nil {
		res.FailWithMessage("分类不存在", c)
		return
	}

	maps := structs.Map(&request)

	err = global.DB.Debug().Model(&tag).Updates(maps).Error

	//结构体转map的第三方包

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改分类失败", c)
		return
	}

	res.OKWithMessage("修改分类成功", c)
}
