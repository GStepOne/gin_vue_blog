package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
)

func (demo DemoModel) IndexExists() bool {
	fmt.Println("client", client)
	exists, err := client.IndexExists(demo.Index()).Do(context.Background())
	if err != nil {
		logrus.Error("索引存在", err.Error())
		return exists
	}
	return exists
}

func (demo DemoModel) RemoveIndex() error {
	indexDelete, err := client.DeleteIndex(demo.Index()).Do(context.Background())
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

func (demo DemoModel) CreateIndex() error {
	if demo.IndexExists() {
		logrus.Info("索引存在,删除中")
		err := demo.RemoveIndex()
		if err != nil {
			logrus.Error("删除索引失败")
			return err
		}
	}
	//没有索引创建
	createIndex, err := client.CreateIndex(demo.Index()).BodyString(demo.Mapping()).Do(context.Background())
	if err != nil {
		logrus.Error("创建索引失败")
		logrus.Error(err.Error())
		return err
	}

	if !createIndex.Acknowledged {
		logrus.Error("创建失败")
		return err
	}

	logrus.Infof("%s 索引创建成功", demo.Index())

	return nil
}

// text 可以比对
func (DemoModel) Mapping() string {
	return `
	{
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
      "user_id": {
        "type": "integer"
      },
      "created_at": {
        "type": "date",
        "null_value": "null",
        "format": "[yyyy-MM-dd HH:mm:ss]"
      }
    }
  }
}
`
}
