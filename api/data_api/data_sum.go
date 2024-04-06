package data_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

type DataSumResponse struct {
	UserCount      int `json:"user_count"`
	ArticleCount   int `json:"article_count"`
	MessageCount   int `json:"message_count"`
	ChatGroupCount int `json:"chat_group_count"`
	NowLoginCount  int `json:"now_login_count"`
	NowSignCount   int `json:"now_sign_count"`
}

func (DataApi) DataSumView(c *gin.Context) {
	var userCount, articleCount, messageCount, chatGroupCount, nowSignCount, nowLoginCount int

	result, err := global.EsClient.Search(models.ArticleModel{}.Index()).
		Query(elastic.NewMatchAllQuery()).
		Size(10000).Do(context.Background())
	if err != nil {
		global.Log.Error("获取文章数据失败")
		articleCount = 0
	}

	articleCount = int(result.Hits.TotalHits.Value)

	global.DB.Model(models.UserModel{}).Select("count(id)").Scan(&userCount)
	global.DB.Model(models.MessageModel{}).Select("count(id)").Scan(&messageCount)
	global.DB.Model(models.ChatModel{IsGroup: true}).Select("count(id)").Scan(&chatGroupCount)

	global.DB.Model(models.LoginDataModel{}).Where("to_days(created_at) = to_days(now())").Select("count(id)").Scan(&nowLoginCount)
	global.DB.Model(models.UserModel{}).Where("to_days(created_at) = to_days(now())").Select("count(id)").Scan(&nowSignCount) //今日注册人数

	res.OKWithData(DataSumResponse{
		UserCount:      userCount,
		ArticleCount:   articleCount,
		MessageCount:   messageCount,
		ChatGroupCount: chatGroupCount,
		NowLoginCount:  nowLoginCount,
		NowSignCount:   nowSignCount,
	}, c)
}
