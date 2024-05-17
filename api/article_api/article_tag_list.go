package article_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

type TagsResponse struct {
	Tag           string   `json:"tag"`
	Count         int      `json:"count"`
	ArticleIDList []string `json:"article_id_list"`
	CreatedAt     string   `json:"created_at"`
}
type TagsType struct {
	DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
	SumOtherDocCount        int `json:"sum_other_doc_count"`
	Buckets                 []struct {
		Key      string `json:"key"`
		DocCount int    `json:"doc_count,omitempty"`
		Article  struct {
			DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
			SumOtherDocCount        int `json:"sum_other_doc_count"`
			Buckets                 []struct {
				Key      string `json:"key"`
				DocCount int    `json:"doc_count,omitempty"`
				Dot      int    `json:"dot,omitempty"`
			} `json:"buckets"`
		} `json:"article"`
		Doct int `json:"doct,omitempty"`
	} `json:"buckets"`
}

func (ArticleApi) ArticleTagListView(c *gin.Context) {

	var cr models.PageView
	_ = c.ShouldBindQuery(&cr)
	if cr.Limit == 0 {
		cr.Limit = 10
	}

	offset := (cr.Page - 1) * cr.Limit
	if offset < 0 {
		offset = 0
	}

	//Cardinality 基数 获取总条数 (这个NewCardinalityAggregation去重 NewValueCountAggregation 不去重)
	result, err := global.EsClient.
		Search(models.ArticleModel{}.Index()).
		Aggregation("tags", elastic.NewCardinalityAggregation().Field("tags")).
		//Aggregation("tags", elastic.NewValueCountAggregation().Field("tags")).
		Size(0).
		Do(context.Background())

	//获取总条数
	cTag, _ := result.Aggregations.Cardinality("tags")
	count := *cTag.Value
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("查询总数失败", c)
		return
	}

	agg := elastic.NewTermsAggregation().Field("tags")
	agg.SubAggregation("article", elastic.NewTermsAggregation().Field("keyword"))
	agg.SubAggregation("page", elastic.NewBucketSortAggregation().From(offset).Size(cr.Limit))
	query := elastic.NewBoolQuery()

	result, err = global.EsClient.
		Search(models.ArticleModel{}.Index()).
		Query(query).
		Aggregation("tags", agg).
		Size(0).
		Do(context.Background())

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("查询失败", c)
		return
	}

	var tagType TagsType
	var tagList = make([]*TagsResponse, 0)

	var tagIdList []string

	_ = json.Unmarshal(result.Aggregations["tags"], &tagType)

	for _, bucket := range tagType.Buckets {
		var articleList []string
		for _, s := range bucket.Article.Buckets {
			articleList = append(articleList, s.Key)
		}

		tagList = append(tagList, &TagsResponse{
			Tag:           bucket.Key,
			Count:         bucket.DocCount,
			ArticleIDList: articleList,
		})

		tagIdList = append(tagIdList, bucket.Key)
	}

	var tagModelList []models.TagModel
	global.DB.Find(&tagModelList, "title in ?", tagIdList)

	var tagData = map[string]string{}
	fmt.Println("%T", tagData)
	for _, model := range tagModelList {
		tagData[model.Title] = model.CreatedAt.Format("2006-01-02 15:04:05")
	}

	for _, k := range tagList {
		k.CreatedAt = tagData[k.Tag]
	}

	res.OkWithList(tagList, int64(count), c)
}
