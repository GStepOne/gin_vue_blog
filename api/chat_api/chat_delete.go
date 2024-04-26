package chat_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

type chatDeleteRequest struct {
	IdList []int `json:"id_list" form:"id_list" query:"id_list"`
}

func (ChatApi) ChatDeleteView(c *gin.Context) {
	var cr chatDeleteRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	fmt.Println(cr.IdList)
	var chatList []models.ChatModel
	err = global.DB.Where("id in ?", cr.IdList).Find(&chatList).Error
	if err != nil {
		res.FailWithError(err, cr, c)
		return
	}

	fmt.Println(chatList)
	if len(chatList) > 0 {
		//err = global.DB.Where("id in ?", cr.IdList).Delete(&chatList).Error
		err = global.DB.Delete(&chatList).Error
		if err != nil {
			res.FailWithMessage("删除失败", c)
			return
		}
	}

	res.OKWithMessage("删除成功", c)
}
