package cron_ser

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/service/redis_ser"
	"gorm.io/gorm"
)

func syncCommentData() {
	info := redis_ser.NewCommentCount().GetInfo()

	if len(info) <= 0 {
		global.Log.Info("今日没有评论数")
		return
	}
	//循环更新

	for id, val := range info {
		err := global.DB.Model(&models.CommentModel{}).Where("id = ?", id).Update("comment_count", gorm.Expr("comment_count + ?", val)).Error
		if err != nil {
			global.Log.Errorf("评论id:%d,评论数：%s同步失败了", id, val)
			continue
		}
	}

	redis_ser.NewCommentCount().Clear()

	global.Log.Infof("所有的评论数量都已经更新成功")

	diggInfo := redis_ser.NewCommentDigg().GetInfo()

	if len(diggInfo) <= 0 {
		global.Log.Info("今日没有评论点赞数")
		return
	}
	//循环更新

	for id, val := range info {
		err := global.DB.Model(&models.CommentModel{}).Where("id = ?", id).Update("digg_count", gorm.Expr("digg_count + ?", val)).Error
		if err != nil {
			global.Log.Errorf("评论id:%d,点赞数：%s同步失败了", id, val)
			continue
		}
	}

	redis_ser.NewCommentDigg().Clear()
	global.Log.Infof("所有的评论点赞量都已经更新成功")
	return
}
