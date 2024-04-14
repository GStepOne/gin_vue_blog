package user_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (UserApi) UserRemoveView(c *gin.Context) {
	//批量删除
	var query models.RemoveRequest

	err := c.ShouldBindJSON(&query)

	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		fmt.Println(err)
		return
	}
	var userList []models.UserModel
	count := global.DB.Find(&userList, query.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("用户不存在", c)
		return
	}

	err = global.DB.Transaction(func(tx *gorm.DB) error {
		//删除 用户、消息表、评论表、用户收藏的文章，用户发布的文章
		err = global.DB.Delete(&userList).Error
		if err != nil {
			global.Log.Error()
			return err
		}
		return nil
	})

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("删除用户失败", c)
		return
	}

	res.OKWithMessage(fmt.Sprintf("共删除 %d 个用户", count), c)
}
