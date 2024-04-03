package article_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"time"
)

type CalendarsResponse struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

type Buckets struct {
	Buckets []struct {
		KeyAsString string `json:"key_as_string"`
		Key         int64  `json:"key"`
		DocCount    int    `json:"doc_count"`
	} `json:"buckets"`
}

var DataCount = map[string]int{}

func (ArticleApi) ArticleCalendarView(c *gin.Context) {
	//按照天来分组（时间聚合）
	agg := elastic.NewDateHistogramAggregation().Field("created_at").CalendarInterval("day")
	//时间段搜索 今天开始
	now := time.Now()
	oneYearAgo := now.AddDate(-1, 0, 0) //去年今日

	format := "2006-01-02 15:04:05"
	//hour minute
	query := elastic.NewRangeQuery("created_at").Lte(now.Format(format)).Gte(oneYearAgo.Format(format))
	result, err := global.EsClient.Search(models.ArticleModel{}.Index()).
		Query(query).
		Aggregation("calendar", agg).
		Size(0).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("查询失败", c)
		return
	}

	var resList = make([]CalendarsResponse, 0) //赋一个控制

	//days := int(now.Sub(oneYearAgo).Hours() / 24)
	//for i := 0; i <= days; i++ {
	//	oneYearAgo.AddDate(0, 0, i)
	//	fmt.Println(oneYearAgo)
	//}

	var data Buckets

	_ = json.Unmarshal(result.Aggregations["calendar"], &data)

	for _, bucket := range data.Buckets {
		Time, _ := time.Parse(format, bucket.KeyAsString)
		DataCount[Time.Format("2006-01-02")] = bucket.DocCount
	}

	days := int(now.Sub(oneYearAgo).Hours() / 24)
	for i := 0; i <= days; i++ {
		day := oneYearAgo.AddDate(0, 0, i).Format("2006-01-02")
		count, _ := DataCount[day]
		resList = append(resList, CalendarsResponse{
			Date:  day,
			Count: count,
		})
	}
	res.OKWithData(resList, c)
}
