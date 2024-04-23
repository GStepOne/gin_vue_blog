package article_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"blog/gin/service/es_ser"
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"time"
)

type ArticleUpdateRequest struct {
	ID       string   `json:"id"`
	Content  string   `json:"content"`
	Title    string   `json:"title"`
	Abstract string   `json:"abstract"` //简介
	Category string   `json:"category"`
	Source   string   `json:"source"`
	Link     string   `json:"link"`
	BannerID uint     `json:"banner_id"`
	Tags     []string `json:"tags"`
}

func (ArticleApi) ArticleUpdateView(c *gin.Context) {
	var cr ArticleUpdateRequest
	err := c.ShouldBindJSON(&cr)

	if err != nil {
		global.Log.Error()
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var bannerUrl string
	if cr.BannerID != 0 {
		//查询一下banner是否还存在
		err = global.DB.Model(models.BannerModel{}).Where("id = ?", cr.BannerID).Select("path").Scan(&bannerUrl).Error
		if err != nil {
			res.FailWithMessage("banner不存在", c)
			return
		}
	}

	article := models.ArticleModel{
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
		Title:     cr.Title,
		Keyword:   cr.Title,
		Abstract:  cr.Abstract,
		Content:   cr.Content, //在list的情况下 不返回content
		Category:  cr.Category,
		Source:    cr.Source,
		Link:      cr.Link,
		BannerID:  cr.BannerID,
		BannerUrl: bannerUrl,
		Tags:      cr.Tags,
	}

	maps := structs.Map(&article)
	var DataMap = map[string]any{}
	for key, v := range maps {
		switch val := v.(type) {
		case string:
			if val == "" {
				continue
			}
		case uint:
			if val == 0 {
				continue
			}
		case []string:
			if len(val) == 0 {
				continue
			}
		case int:
			if val == 0 {
				continue
			}
			//case ctype.Array:
			//	if len(val) == 0 {
			//		continue
			//	}
		}

		DataMap[key] = v
	}

	//这里的article会变化因为是传递引用
	err = article.GetDataById(cr.ID)
	if err != nil {
		global.Log.Error(err.Error())
		res.FailWithMessage("文章不存在", c)
		return
	}
	fmt.Println("DataMap", DataMap)
	err = es_ser.ArticleUpdate(cr.ID, DataMap)
	if err != nil {
		global.Log.Error(err.Error())
		res.FailWithMessage("更新失败", c)
		return
	}

	newArticle, _ := es_ser.CommonDetail(cr.ID)

	if article.Content != newArticle.Content || article.Title != newArticle.Title {
		//先把原来的全文索引删掉
		es_ser.DeleteFullTextByArticleId(cr.ID)
		//再同步一下新的全文索引
		es_ser.AsyncArticleByFullText(cr.ID, cr.Title, cr.Content)
	}

	res.OKWithMessage("更新成功", c)
}
