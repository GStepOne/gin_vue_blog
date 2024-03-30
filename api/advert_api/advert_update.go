package advert_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

// AdvertUpdateView 广告更新
// @Tags 广告更新
// @Summary 广告更新
// @Description 广告更新
// @Param data body AdvertRequest false "更新广告的参数"
// @Router /api/advert/:id [put]
// @Produce json
// @Success 200 {object} res.Response{data=string}
func (AdvertApi) AdvertUpdateView(c *gin.Context) {
	var request AdvertRequest

	id := c.Param("id")

	fmt.Println(id)
	err := c.ShouldBindJSON(&request)

	if err != nil {
		res.FailWithError(err, &request, c)
		return
	}

	var advert models.AdvertModel

	err = global.DB.Debug().Take(&advert, id).Error
	if err != nil {
		res.FailWithMessage("广告不存在", c)
		return
	}

	maps := structs.Map(&request)

	err = global.DB.Debug().Model(&advert).Updates(maps).Error

	//结构体转map的第三方包

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改广告失败", c)
		return
	}

	res.OKWithMessage("修改广告成功", c)
}
