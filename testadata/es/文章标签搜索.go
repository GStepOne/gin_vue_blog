package main

import (
	"blog/gin/core"
	"blog/gin/global"
	"blog/gin/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
)

type TagsResponse struct {
	Tag           string   `json:"tag"`
	Count         int      `json:"count"`
	ArticleIDList []string `json:"article_id_list"`
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

func main() {
	core.InitCoreConf()
	global.Log = core.InitLogger()
	global.EsClient = core.EsConnect()

	agg := elastic.NewTermsAggregation().Field("tags")
	agg.SubAggregation("article", elastic.NewTermsAggregation().Field("keyword"))
	agg.SubAggregation("article_id", elastic.NewTermsAggregation().Field("_id"))
	query := elastic.NewBoolQuery()

	result, err := global.EsClient.
		Search(models.ArticleModel{}.Index()).
		Query(query).
		Aggregation("tags", agg).
		Size(0).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		return
	}

	var tagType TagsType
	var tagList = make([]TagsResponse, 0)
	//fmt.Println(result.Aggregations["tags"])
	_ = json.Unmarshal(result.Aggregations["tags"], &tagType)
	fmt.Println(string(result.Aggregations["tags"]))

	fmt.Println(tagType)
	for _, bucket := range tagType.Buckets {
		var articleList []string
		for _, s := range bucket.Article.Buckets {
			articleList = append(articleList, s.Key)
		}

		tagList = append(tagList, TagsResponse{
			Tag:           bucket.Key,
			Count:         bucket.DocCount,
			ArticleIDList: articleList,
		})
	}

	//fmt.Println(tagList)
}
