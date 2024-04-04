package models

import (
	"blog/gin/global"
	"blog/gin/models/ctype"
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

type ArticleModel struct {
	ID            string `json:"id" structs:"id"`
	CreatedAt     string `json:"created_at" structs:"created_at"`
	UpdatedAt     string `json:"updated_at" structs:"updated_at"`
	Title         string `json:"title" structs:"title"`
	Keyword       string `json:"keyword,omit(list)" structs:"keyword"`
	Abstract      string `json:"abstract" structs:"abstract"`
	Content       string `json:"content,omit(list)" structs:"content"` //在list的情况下 不返回content
	LookCount     int    `json:"look_count" structs:"look_count"`
	CommentCount  int    `json:"comment_count" structs:"comment_count"`
	DiggCount     int    `json:"digg_count" structs:"digg_count"`
	CollectsCount int    `json:"collects_count" structs:"collects_count"`

	UserId       uint   `json:"user_id" structs:"user_id"`
	UserNickName string `json:"user_nick_name" structs:"user_nick_name"`
	UserAvatar   string `json:"user_avatar" structs:"user_avatar"`

	Category string `json:"category" structs:"category"`
	Source   string `json:"source" structs:"source"`
	Link     string `json:"link" structs:"link"`

	BannerID  uint   `json:"banner_id" structs:"banner_id"`
	BannerUrl string `json:"banner_url" structs:"banner_url"`

	Tags ctype.Array `json:"tags" structs:"tags"`
}

func (ArticleModel) Index() string {
	return "article_search"
}

func (ArticleModel) Mapping() string {
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
      "content": {
        "type": "text"
      },
      "keyword": {
        "type": "keyword"
      },
      "abstract": {
        "type": "text"
      },
      "look_count": {
        "type": "integer"
      },
      "comment_count": {
        "type": "integer"
      },
      "digg_count": {
        "type": "integer"
      },
      "collects_count": {
        "type": "integer"
      },
      "user_id": {
        "type": "integer"
      },
      "user_nick_name": {
        "type": "keyword"
      },
      "user_avatar": {
        "type": "keyword"
      },
      "category": {
        "type": "keyword"
      },
      "source": {
        "type": "keyword"
      },
      "banner_id": {
        "type": "integer"
      },
      "link": {
        "type": "keyword"
      },
      "banner_url": {
        "type": "keyword"
      },
      "tags": {
        "type": "keyword"
      },
      "created_at": {
        "type": "date",
        "null_value": "null",
        "format": "[yyyy-MM-dd HH:mm:ss]"
      },
      "updated_at": {
        "type": "date",
        "null_value": "null",
        "format": "[yyyy-MM-dd HH:mm:ss]"
      }
    }
  }
}
`
}

func (article ArticleModel) IndexExists() bool {
	client := global.EsClient
	exists, err := client.IndexExists(article.Index()).Do(context.Background())
	if err != nil {
		logrus.Error("索引存在", err.Error())
		return exists
	}
	return exists
}

func (article ArticleModel) RemoveIndex() error {
	client := global.EsClient
	indexDelete, err := client.DeleteIndex(article.Index()).Do(context.Background())
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

func (article ArticleModel) CreateIndex() error {
	client := global.EsClient
	if article.IndexExists() {
		logrus.Info("索引存在,删除中1")
		err := article.RemoveIndex()
		if err != nil {
			logrus.Error("删除索引失败")
			return err
		}
	}
	//没有索引创建
	createIndex, err := client.CreateIndex(article.Index()).BodyString(article.Mapping()).Do(context.Background())
	if err != nil {
		logrus.Error("创建索引失败")
		logrus.Error(err.Error())
		return err
	}

	if !createIndex.Acknowledged {
		logrus.Error("创建失败")
		return err
	}

	logrus.Infof("%s 索引创建成功1", article.Index())

	return nil
}

// 添加
func (data *ArticleModel) Create() (err error) {
	//第一个index 是文档
	client := global.EsClient
	indexResponse, err := client.Index().
		Index(data.Index()).
		BodyJson(data).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	logrus.Infof("%#v", indexResponse)
	//传递给data
	data.ID = indexResponse.Id

	return nil
}

func (article ArticleModel) IsExistsData() bool {
	boolSearch := elastic.NewBoolQuery()

	boolSearch.Must(
		elastic.NewTermQuery("keyword", article.Keyword),
	)
	res, err := global.EsClient.Search(article.Index()).Query(boolSearch).Size(1).Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return false
	}

	if res.Hits.TotalHits.Value > 0 {
		return true
	}
	return false
}

func (article *ArticleModel) GetDataById(id string) error {

	res, err := global.EsClient.Get().Index(article.Index()).Id(id).Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	err = json.Unmarshal(res.Source, &article)
	if err != nil {
		return err
	}
	return nil
}
