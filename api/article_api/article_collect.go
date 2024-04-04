package article_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"blog/gin/service/es_ser"
	"blog/gin/utils/jwt"
	"github.com/gin-gonic/gin"
)

func (ArticleApi) ArticleCollectCreateView(c *gin.Context) {
	var cr models.ESIDRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	model, err := es_ser.CommonDetail(cr.ID)
	if err != nil {
		res.FailWithMessage("文章不存在", c)
		return
	}

	var num = -1
	var coll models.UserCollectModel
	err = global.DB.Take(&coll, "user_id=? and article_id=?", claims.UserId, cr.ID).Error
	if err != nil {
		global.DB.Create(&models.UserCollectModel{
			UserID:    claims.UserId,
			ArticleID: cr.ID,
		})
		num = 1
	}
	//取消收藏
	//文章数-1
	//fmt.Println("coll", coll)
	global.DB.Debug().Delete(&coll) //如果上面找不到 这里是删除不了的
	err = es_ser.ArticleUpdate(cr.ID, map[string]any{
		"collects_count": model.CollectsCount + num,
	})
	if err != nil {
		res.FailWithMessage("取消收藏失败", c)
		return
	}
	if num != 1 {
		res.OKWithMessage("取消收藏成功", c)
	} else {
		res.OKWithMessage("收藏成功", c)
	}
}
