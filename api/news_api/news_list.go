package news_api

import (
	"blog/gin/models/res"
	"blog/gin/utils/requests"
	"encoding/json"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

type Params struct {
	ID   string `json:"id"`
	Size uint   `json:"size"`
}

type NewsData struct {
	Index    int    `json:"index"`
	Title    string `json:"title"`
	HotValue string `json:"hot_value"`
	Link     string `json:"link"`
}

type Header struct {
	Signaturekey string `form:"signaturekey" structs:"signaturekey"`
	Version      string `form:"version" structs:"version"`
	UserAgent    string `form:"user-agent" structs:"user-agent"`
}

type NewsResponse struct {
	Code int        `json:"code"`
	Data []NewsData `json:"data"`

	Msg string `json:"msg"`
}

const API = "https://api.codelife.cc/api/top/list"
const TIMEOUT = 2 * time.Second

func (NewsApi) NewListView(c *gin.Context) {

	var cr Params
	var header Header

	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	httpResponse, err := requests.Post(API, cr, structs.Map(header), TIMEOUT)
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

	res.OKWithData(response, c)

	return

}
