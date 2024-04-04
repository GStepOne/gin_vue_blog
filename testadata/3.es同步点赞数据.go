package main

import (
	"blog/gin/core"
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/service/redis_ser"
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

func main() {
	core.InitCoreConf()
	global.Redis = core.ConnectRedis()
	global.EsClient = core.EsConnect()

	res, err := global.EsClient.Search(models.ArticleModel{}.Index()).
		Query(elastic.NewMatchAllQuery()).
		Size(10000).
		Do(context.Background()) // 指定返回字段

	if err != nil {
		logrus.Error(err)
		return
	}

	diggInfo := redis_ser.GetDiggInfo()
	lookInfo := redis_ser.GetLookInfo()

	for _, hit := range res.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)

		digg := diggInfo[hit.Id]
		newDigg := article.DiggCount + digg

		look := lookInfo[hit.Id]
		newLook := article.LookCount + look

		newDiggCount := article.DiggCount + digg
		newLookCount := article.LookCount + look

		if newDiggCount == article.DiggCount && newLookCount == article.LookCount {
			logrus.Info(article.Title, "点赞数与浏览数无变化")
		}

		_, err := global.EsClient.Update().Index(article.Index()).Id(hit.Id).
			Doc(map[string]int{
				"digg_count": newDigg,
				"look_count": newLook,
			}).
			Do(context.Background())

		if err != nil {
			logrus.Error(err.Error())
			continue
		}

		num := fmt.Sprintf("点赞、浏览数据同步成功，点赞数%d,浏览数%d", newDigg, newLook)

		logrus.Info(article.Title+":", num)
	}

	//删除点赞和浏览量
	redis_ser.DiggClear()
	redis_ser.LookClear()

}
