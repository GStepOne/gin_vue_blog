package category_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"blog/gin/service/common"
	"fmt"
	"github.com/gin-gonic/gin"
)

type CategoryView struct {
	Title string `gorm:"size:32" binding:"required" max:"20" msg:"请输入标题"  json:"title"`
	Value string `json:"value"`
}

func (CategoryApi) CategoryListView(c *gin.Context) {
	var request models.PageView

	err := c.ShouldBindQuery(&request)
	referer := c.GetHeader("Referer")
	fmt.Println(referer)

	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	//判断referer 是否包含admin 如果是，就全部返回，不是，就返回is_show=true,
	list, count, _ := common.ComList(models.CategoryModel{}, common.Option{
		PageView: request,
	})

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("获取标签失败", c)
		return
	}
	res.OkWithList(list, count, c)
}
