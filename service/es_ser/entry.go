package es_ser

import (
	"blog/gin/global"
	"blog/gin/models"
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

func CommonList(key string, page, limit int) (list []models.ArticleModel, count int, err error) {
	boolSearch := elastic.NewBoolQuery()
	from := page
	if key != "" {
		boolSearch.Must(
			elastic.NewMatchQuery("title", key),
		)
	}

	if limit == 0 {
		limit = 10
	}

	if from == 0 {
		from = 1
	}

	client := global.EsClient
	sourceContext := elastic.NewFetchSourceContext(true).Exclude("content")
	res, err := client.Search(models.ArticleModel{}.Index()).
		Query(boolSearch).FetchSourceContext(sourceContext).
		From((from - 1) * limit).Size(limit).Do(context.Background()) // 指定返回字段
	if err != nil {
		logrus.Error(err.Error())
		return
	}

	//查询到的总条数
	num := int(res.Hits.TotalHits.Value)
	//list = []models.ArticleModel{}
	for _, hit := range res.Hits.Hits {
		var article models.ArticleModel
		data, err := hit.Source.MarshalJSON() //es里面的_source
		if err != nil {
			logrus.Error(err)
			continue
		}
		err = json.Unmarshal(data, &article)
		if err != nil {
			logrus.Error(err)
			continue
		}
		article.ID = hit.Id
		list = append(list, article)
	}

	return list, num, nil

}

func CommonDetail(id string) (model models.ArticleModel, err error) {
	res, err := global.EsClient.Get().Index(models.ArticleModel{}.Index()).Id(id).Do(context.Background())
	if err != nil {
		return
	}
	err = json.Unmarshal(res.Source, &model)
	if err != nil {
		return
	}
	model.ID = res.Id
	return model, nil
}
