package news_api

import (
	"blog/gin/models/res"
	"blog/gin/service/redis_ser"
	"blog/gin/utils/requests"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

type Params struct {
	ID   string `json:"id" query:"id" form:"id"`
	Size uint   `json:"size" query:"size" form:"size"`
}

type NewsData = any

//type NewsData =struct {
//	//Index    int    `json:"index"`
//	//Title    string `json:"title"`
//	//HotValue string `json:"hot_value"`
//	//Link     string `json:"link"`
//}

type Header struct {
	Signaturekey string `form:"signaturekey" structs:"signaturekey"`
	Version      string `form:"version" structs:"version"`
	UserAgent    string `form:"user-agent" structs:"user-agent"`
}

type NewsResponse struct {
	Code int                  `json:"code"`
	Data []redis_ser.NewsData `json:"data"`
	Msg  string               `json:"msg"`
}

const API = "https://api.codelife.cc/api/top/list"
const TIMEOUT = 2 * time.Second

func (NewsApi) NewListView(c *gin.Context) {

	var cr Params
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	key := fmt.Sprintf("%s-%d", cr.ID, cr.Size)
	var list []redis_ser.NewsData
	list = redis_ser.GetNews(key)
	if len(list) != 0 {
		fmt.Println("from cache")
		res.OKWithData(list, c)
		return
	}
	finalUrl := fmt.Sprintf("%s?id=%s&size=%d", API, cr.ID, cr.Size)

	httpResponse, err := requests.Get(finalUrl, nil, TIMEOUT)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	var response NewsResponse
	byteData, err := io.ReadAll(httpResponse.Body)
	err = json.Unmarshal(byteData, &response)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	if response.Code != 200 {
		res.FailWithMessage(response.Msg, c)
		return
	}
	redis_ser.SetNews(key, response.Data)
	res.OKWithData(response.Data, c)

	return
}
