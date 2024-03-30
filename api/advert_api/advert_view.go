package advert_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"blog/gin/service/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

type AdvertView struct {
	Title  string `gorm:"size:32" binding:"required" max:"20" msg:"请输入标题"  json:"title"`
	Href   string `json:"href" binding:"required,url" msg:"请输入跳转链接"`
	Images string `json:"images" binding:"required,url" msg:"请输入一个合法的图片地址"`
	IsShow bool   `json:"is_show" binding:"required"  msg:"请选择是否展示"`
}

// AdvertListView 广告列表
// @Tags 广告列表
// @Summary 广告列表
// @Description 广告列表
// @Param limit query models.PageView false ""
// @Param page query models.PageView false "表示多个参数"
// @Router /api/advert [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.AdvertModel]}
func (AdvertApi) AdvertListView(c *gin.Context) {
	var request models.PageView

	err := c.ShouldBindQuery(&request)
	referer := c.GetHeader("Referer")
	isShow := true
	fmt.Println(referer)

	if strings.Contains(referer, "admin") {
		//admin来的
		isShow = false //等于false gorm会忽略
	}

	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	//判断referer 是否包含admin 如果是，就全部返回，不是，就返回is_show=true,
	list, count, _ := common.ComList(models.AdvertModel{IsShow: isShow}, common.Option{
		PageView: request,
	})

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("添加广告失败", c)
		return
	}
	res.OkWithList(list, count, c)
}
