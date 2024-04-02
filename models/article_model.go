package models

import (
	"blog/gin/global"
	"blog/gin/models/ctype"
	"context"
	"github.com/sirupsen/logrus"
)

type ArticleModel struct {
	ID            string `json:"id"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	Title         string `json:"title"`
	Abstract      string `json:"abstract"`
	Content       string `json:"content"`
	LookCount     int    `json:"look_count"`
	CommentCount  int    `json:"comment_count"`
	DiggCount     int    `json:"digg_count"`
	CollectsCount int    `json:"collects_count"`

	UserId       uint   `json:"user_id"`
	UserNickName string `json:"user_nick_name"`
	UserAvatar   string `json:"user_avatar"`

	Category string `json:"category"`
	Source   string `json:"source"`
	Link     string `json:"link"`

	BannerID  uint   `json:"banner_id"`
	BannerUrl string `json:"banner_url"`

	Tags ctype.Array `json:"tags"`
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
        "type": "text"
      },
      "user_avatar": {
        "type": "text"
      },
      "category": {
        "type": "text"
      },
      "source": {
        "type": "text"
      },
      "banner_id": {
        "type": "integer"
      },
      "link": {
        "type": "text"
      },
      "banner_url": {
        "type": "text"
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
		logrus.Info("索引存在,删除中")
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

	logrus.Infof("%s 索引创建成功", article.Index())

	return nil
}

// 添加
func (data ArticleModel) Create() (err error) {
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
