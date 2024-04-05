package article_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

func (ArticleApi) FullTextSearch(c *gin.Context) {
	var cr models.PageView
	_ = c.ShouldBindQuery(&cr)
	boolQuery := elastic.NewBoolQuery()
	if cr.Key != "" {
		boolQuery.Must(elastic.NewMultiMatchQuery(cr.Key, "title", "body"))
	}

	result, err := global.EsClient.Search(models.FullTextModel{}.Index()).
		//Query(elastic.NewMultiMatchQuery(cr.Key, "title", "body")).
		Query(boolQuery).
		Highlight(elastic.NewHighlight().Field("body")).Size(100).Do(context.Background())
	if err != nil {
		return
	}

	count := result.Hits.TotalHits.Value
	fullTextList := []models.FullTextModel{}

	for _, hit := range result.Hits.Hits {
		var model models.FullTextModel
		json.Unmarshal(hit.Source, &model)

		body, ok := hit.Highlight["body"]
		if ok {
			model.Body = body[0]
		}
		fullTextList = append(fullTextList, model)
	}

	res.OkWithList(fullTextList, count, c)
}
