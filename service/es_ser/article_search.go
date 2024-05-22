package es_ser

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/service/redis_ser"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"strings"
)

type Option struct {
	models.PageView
	Tag    string   `form:"tag"`
	Fields []string `form:"fields"`
}

func (o Option) GetForm() int {
	if o.Page == 0 {
		o.Page = 1
	}

	if o.Limit == 0 {
		o.Limit = 1
	}
	return (o.Page - 1) * o.Limit
}

func CommonList(option Option) (list []models.ArticleModel, count int, err error) {
	boolSearch := elastic.NewBoolQuery()
	if option.Key != "" {
		boolSearch.Must(
			elastic.NewMultiMatchQuery(option.Key, option.Fields...),
		)
	} else {
		// 如果没有搜索关键词，则匹配所有文章
		boolSearch.Must(elastic.NewMatchAllQuery())
	}

	if option.Tag != "" {
		boolSearch.Must(
			elastic.NewMultiMatchQuery(option.Tag, "tags"),
		)
	}

	client := global.EsClient
	sourceContext := elastic.NewFetchSourceContext(true).Exclude("content")

	// 高亮显示 highlight
	type SortField struct {
		Field     string
		Ascending bool
		Unmapped  string
	}
	sortField := SortField{
		Field:     "created_at",
		Ascending: false,
		Unmapped:  "date",
	}

	if option.Sort != "" {
		_list := strings.Split(option.Sort, " ")
		if len(_list) == 2 && (_list[1] == "desc" || _list[1] == "asc") {
			sortField.Field = _list[0]
			sortField.Ascending = _list[1] == "asc"
		}
	}

	// 构建排序对象
	//sort := elastic.NewFieldSort(sortField.Field).
	//	UnmappedType(sortField.Unmapped).
	//	Order(sortField.Ascending)

	// 使用 SearchSource 构建搜索请求
	searchSource := elastic.NewSearchSource().
		Query(boolSearch).
		FetchSourceContext(sourceContext).
		Highlight(elastic.NewHighlight().Field("title")).
		From(option.GetForm()).
		Size(option.Limit).Sort(sortField.Field, sortField.Ascending)
	//SortBy(sort)

	// 打印搜索请求的 JSON 表示
	searchSourceJSON, err := searchSource.Source()
	if err != nil {
		logrus.Error("Failed to get search source JSON:", err)
	} else {
		sourceJSON, _ := json.Marshal(searchSourceJSON)
		fmt.Printf("Elasticsearch Search Request Body: %s\n", sourceJSON)
	}

	// 构建搜索请求
	searchReq := client.Search(models.ArticleModel{}.Index()).SearchSource(searchSource)

	// 执行搜索请求
	res, err := searchReq.Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	// 查询到的总条数
	num := int(res.Hits.TotalHits.Value)
	for _, hit := range res.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)
		if err != nil {
			logrus.Error(err)
			continue
		}
		article.ID = hit.Id
		title, ok := hit.Highlight["title"]
		if ok {
			article.Title = title[0]
		}
		digg := redis_ser.NewDigg().Get(hit.Id)
		look := redis_ser.NewArticleLook().Get(hit.Id)
		commentCount := redis_ser.NewCommentCount().Get(hit.Id)
		article.DiggCount += digg
		article.LookCount += look
		article.CommentCount += commentCount
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
	model.LookCount += redis_ser.NewArticleLook().Get(id)
	return model, nil
}

func CommonDetailByKeyWord(keyword string) (model models.ArticleModel, err error) {
	res, err := global.EsClient.Search().Index(models.ArticleModel{}.Index()).
		Query(elastic.NewTermQuery("keyword", keyword)).
		Size(1).
		Do(context.Background())
	if err != nil {
		return
	}
	if res.Hits.TotalHits.Value == 0 {
		return model, errors.New("文章不存在")
	}
	hit := res.Hits.Hits[0]
	err = json.Unmarshal(hit.Source, &model)
	if err != nil {
		return
	}
	return

}

func ArticleUpdate(id string, data map[string]any) error {
	_, err := global.EsClient.Update().
		Index(models.ArticleModel{}.Index()).
		Id(id).
		Doc(data).
		Refresh("true"). //立马更新，否则有延迟
		Do(context.Background())
	return err
}
