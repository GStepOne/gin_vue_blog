package main

import (
	"blog/gin/core"
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/service/es_ser"
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

func main() {
	core.InitCoreConf()
	core.InitLogger()
	global.EsClient = core.EsConnect()

	boolSearch := elastic.NewMatchAllQuery()

	res, _ := global.EsClient.Search(models.ArticleModel{}.Index()).Query(boolSearch).Size(1000).Do(context.Background())

	for _, hit := range res.Hits.Hits {
		var article models.ArticleModel
		_ = json.Unmarshal(hit.Source, &article)

		indexList := es_ser.GetSearchIndexContent(hit.Id, article.Title, article.Content)

		bulk := global.EsClient.Bulk()

		for _, indexData := range indexList {
			req := elastic.NewBulkIndexRequest().Index(models.FullTextModel{}.Index()).Doc(indexData)
			bulk.Add(req)
		}

		result, err := bulk.Do(context.Background())
		if err != nil {
			logrus.Error(err)
			continue
		}

		fmt.Println(article.Title, "添加成功", "共", len(result.Succeeded()), "条")
	}
}
