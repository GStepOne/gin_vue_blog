package cron_ser

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/service/redis_ser"
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
)

func SyncArticleData() {
	result, err := global.EsClient.Search(models.ArticleModel{}.Index()).Query(elastic.NewMatchAllQuery()).Size(10000).Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		return
	}

	diggInfo := redis_ser.NewDigg().GetInfo()
	lookInfo := redis_ser.NewArticleLook().GetInfo()
	commentInfo := redis_ser.NewArticleLook().GetInfo()

	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)
		if err != nil {
			global.Log.Error(err)
			continue
		}

		//加上缓存的数据
		digg := diggInfo[hit.Id]
		newDigg := article.DiggCount + digg

		lookDigg := lookInfo[hit.Id]
		newLook := article.LookCount + lookDigg

		commentDigg := commentInfo[hit.Id]
		newComment := article.CommentCount + commentDigg
		//要判断 是否需要改变

		if digg == 0 && lookDigg == 0 && commentDigg == 0 {
			global.Log.Info("数据无变化，无需更新")
			continue
		}

		_, err = global.EsClient.Update().Index(models.ArticleModel{}.Index()).Id(hit.Id).Doc(map[string]int{
			"look_count":    newLook,
			"comment_count": newComment,
			"digg_count":    newDigg,
		}).Do(context.Background())

		if err != nil {
			global.Log.Error(fmt.Sprintf("更新失败：%s", article.Title))
			continue
		}

		global.Log.Infof("%s 更新成功,点赞数:%d，评论数:%d,浏览量：%s", article.Title, newDigg, newComment, newComment)

	}

	//删除redis中的数据
	redis_ser.NewDigg().Clear()
	redis_ser.NewArticleLook().Clear()
	//redis_ser.NewCommentCount().Clear()
}
