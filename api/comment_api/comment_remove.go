package comment_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"blog/gin/service/redis_ser"
	"blog/gin/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//删除评论
// 1删除根评论
//		删除
// 2删除子评论，根评论数 删除

func (CommentApi) CommentRemoveView(c *gin.Context) {
	var cr CommentIDRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var commentModel models.CommentModel
	err = global.DB.Take(&commentModel, cr.ID).Error

	total, ids := TotalSubComment(&commentModel) //把自己算进去
	fmt.Println("原来顺序", ids)
	redis_ser.NewCommentCount().SetCount(commentModel.ArticleID, -total)
	//fmt.Printf("当前id%d,下的评论数为%d", cr.ID, sonTotal)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	//判断是否是子评论
	if commentModel.ParentCommentModel != nil {
		//子评论
		global.DB.Model(&models.CommentModel{}).
			Where("id=?", *commentModel.ParentCommentID).
			Update("comment_count", gorm.Expr("comment_count - ?", total+1)) //加1是包含自己
	}
	//父评论:删除子评论以及当前评论
	ids = append(ids, cr.ID)
	newSortIds := utils.Reverse(ids)
	for _, id := range newSortIds {
		global.DB.Model(models.CommentModel{}).Debug().Delete("id = ?", id) //循环删除。。。
	}

	res.OKWithMessage(fmt.Sprintf("删除成功%d条", len(newSortIds)), c)
}
