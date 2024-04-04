package models

import (
	"blog/gin/global"
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

type FullTextModel struct {
	ID    string `json:"id" structs:"id"`
	Key   string `json:"key"`
	Title string `json:"title" structs:"title"`
	Slug  string `json:"slug" structs:"slug"`
	Body  string `json:"body" structs:"body"`
}

func (FullTextModel) Index() string {
	return "full_text_index"
}

func (FullTextModel) Mapping() string {
	return `{
			  "settings": {
				"index": {
				  "max_result_window": "100000"
				}
			  },
			  "mappings": {
				"properties": {
				  "title": {
					"type": "text"
				  },
	  				"key": {
					"type": "keyword"
				  },
				  "slug": {
					"type": "keyword"
				  },
				  "body": {
					"type": "text"
				  }
				}
			  }
			}
`
}

func (fulltext FullTextModel) IndexExists() bool {
	client := global.EsClient
	exists, err := client.IndexExists(fulltext.Index()).Do(context.Background())
	if err != nil {
		logrus.Error("索引存在", err.Error())
		return exists
	}
	return exists
}

func (fulltext FullTextModel) RemoveIndex() error {
	client := global.EsClient
	indexDelete, err := client.DeleteIndex(fulltext.Index()).Do(context.Background())
	if err != nil {
		logrus.Error("删除索引失败")
		logrus.Error(err.Error())
		return err
	}

	if !indexDelete.Acknowledged {
		logrus.Error("索引删除失败")
		logrus.Error(err.Error())
		return err
	}

	return nil
}

func (fulltext FullTextModel) CreateIndex() error {
	client := global.EsClient
	if fulltext.IndexExists() {
		logrus.Info("索引存在,删除中1")
		err := fulltext.RemoveIndex()
		if err != nil {
			logrus.Error("删除索引失败")
			return err
		}
	}
	//没有索引创建
	createIndex, err := client.CreateIndex(fulltext.Index()).BodyString(fulltext.Mapping()).Do(context.Background())
	if err != nil {
		logrus.Error("创建索引失败")
		logrus.Error(err.Error())
		return err
	}

	if !createIndex.Acknowledged {
		logrus.Error("创建失败")
		return err
	}

	logrus.Infof("%s 索引创建成功1", fulltext.Index())

	return nil
}

// 添加
func (fulltext FullTextModel) Create() (err error) {
	//第一个index 是文档
	client := global.EsClient
	indexResponse, err := client.Index().
		Index(fulltext.Index()).
		BodyJson(fulltext).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	logrus.Infof("%#v", indexResponse)
	//传递给data
	fulltext.ID = indexResponse.Id

	return nil
}

func (fulltext FullTextModel) IsExistsData() bool {
	boolSearch := elastic.NewBoolQuery()

	boolSearch.Must(
	//elastic.NewTermQuery("keyword", fulltext.Keyword),
	)
	res, err := global.EsClient.Search(fulltext.Index()).Query(boolSearch).Size(1).Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return false
	}

	if res.Hits.TotalHits.Value > 0 {
		return true
	}
	return false
}

func (fulltext FullTextModel) GetDataById(id string) error {

	var fullText FullTextModel
	res, err := global.EsClient.Get().Index(fulltext.Index()).Id(id).Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return err
	}

	err = json.Unmarshal(res.Source, &fullText)

	return err
}
