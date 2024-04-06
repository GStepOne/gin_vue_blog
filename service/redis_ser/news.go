package redis_ser

import (
	"blog/gin/global"
	"encoding/json"
)

const newsIndex = "news_index"

type NewsData struct {
	Index    int    `json:"index"`
	Title    string `json:"title"`
	HotValue string `json:"hot_value"`
	Link     string `json:"link"`
}

func SetNews(key string, newData []NewsData) error {
	byteData, _ := json.Marshal(newData)
	err := global.Redis.HSet(newsIndex, key, byteData).Err()
	//global.Redis.Expire(newsIndex, news_api.TIMEOUT)
	return err
}

func GetNews(key string) (newData []NewsData) {
	res := global.Redis.HGet(newsIndex, key).Val()
	_ = json.Unmarshal([]byte(res), &newData)
	return
}
