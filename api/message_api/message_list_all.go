package message_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"blog/gin/service/common"
	"github.com/gin-gonic/gin"
)

func (MessageApi) MessageListAllView(c *gin.Context) {
	var cr models.PageView

	err := c.ShouldBindQuery(&cr)

	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	//_claims, _ := c.Get("claims")
	//claims := _claims.(*jwt.CustomClaims)

	//role := claims.Role
	//if uint(role) == ctype.PermissionAdmin {
	//
	//}

	list, count, _ := common.ComList(&models.MessageModel{}, common.Option{
		PageView: cr,
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("消息发送失败", c)
		return
	}

	res.OkWithList(list, count, c)
	return
}
