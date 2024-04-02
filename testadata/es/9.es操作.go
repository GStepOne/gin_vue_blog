package main

import (
	"blog/gin/core"
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"time"
)

var client *elastic.Client

func EsConnect() *elastic.Client {
	var err error
	sniffOpt := elastic.SetSniff(false)
	host := "http://127.0.0.1:9200"
	client, err := elastic.NewClient(
		elastic.SetURL(host),
		sniffOpt,
		elastic.SetBasicAuth("", ""),
	)
	if err != nil {
		logrus.Fatalf("es链接失败%s", err.Error())
	}

	return client
}
func init() {
	core.InitCoreConf()
	core.InitLogger()
	client = EsConnect()
}

type DemoModel struct {
	Title     string `json:"title"`
	Id        string `json:"id"`
	UserId    uint   `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

func (DemoModel) Index() string {
	return "demo_index"
}

// 添加
func Create(data *DemoModel) (err error) {
	//第一个index 是文档
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
	data.Id = indexResponse.Id

	return nil
}

func Update(id string, data *DemoModel) error {
	_, err := client.Update().Index(DemoModel{}.Index()).Id(id).
		Doc(map[string]string{
			"title": data.Title,
		}).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return err
	}

	logrus.Info("更新demo成功")
	return nil
}

func Remove(IdList []string) (count int, err error) {
	//先创建一批
	BulkService := client.Bulk().Index(DemoModel{}.Index()).Refresh("true")
	for _, id := range IdList {
		req := elastic.NewBulkDeleteRequest().Id(id)
		BulkService.Add(req)
	}
	res, err := BulkService.Do(context.Background())
	return len(res.Succeeded()), err
}

func FindList(keyword string, page, limit int) (demoList []DemoModel, count int) {
	boolSearch := elastic.NewBoolQuery()
	from := page
	if keyword != "" {
		boolSearch.Must(
			elastic.NewMatchQuery("title", keyword),
		)
	}

	if limit == 0 {
		limit = 10
	}

	if from == 0 {
		limit = 1
	}

	sourceContext := elastic.NewFetchSourceContext(true).Include("title")
	res, err := client.Search(DemoModel{}.Index()).Query(boolSearch).FetchSourceContext(sourceContext). // 指定返回字段
														From((from - 1) * limit).Size(limit).Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return
	}

	//查询到的总条数
	num := int(res.Hits.TotalHits.Value)
	for _, hit := range res.Hits.Hits {
		var demo DemoModel
		data, err := hit.Source.MarshalJSON() //es里面的_source
		if err != nil {
			logrus.Error(err)
			continue
		}
		err = json.Unmarshal(data, &demo)
		if err != nil {
			logrus.Error(err)
			continue
		}
		demo.Id = hit.Id
		demoList = append(demoList, demo)
	}

	return demoList, num

}
func main() {
	//删除索引
	//DemoModel{}.CreateIndex()
	//添加数据
	Create(&DemoModel{Title: "天天2", UserId: 1, CreatedAt: time.Now().Format("2006-01-02 15:04:05")})
	//更新数据
	//Update("mKWInY4BReI8ZXCFpxUB", &DemoModel{
	//	Title: "天天爱学习",
	//})
	// 查询数据
	//list, count := FindList("", 1, 10)
	//fmt.Println(list, count)

	//删除数据

	//Remove([]string{"mKWInY4BReI8ZXCFpxUB"})

}
