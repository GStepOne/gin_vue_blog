package article_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"blog/gin/service/es_ser"
	"blog/gin/utils/jwt"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

func (ArticleApi) ArticleCollBatchRemoveView(c *gin.Context) {
	var cr models.ESIDListRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	var collects []models.UserCollectModel
	var articleIDList []string
	global.DB.Debug().
		Find(&collects, "user_id = ? and article_id in ?", claims.UserId, cr.IDList).
		Select("article_id").
		Scan(&articleIDList)

	var idList []interface{}

	for _, v := range articleIDList {
		idList = append(idList, v)
	}

	if err != nil {
		global.Log.Error("删除文章失败", err)
		res.OKWithMessage("删除文章失败", c)
		return
	}

	//去更新es中的数据
	boolSearch := elastic.NewTermsQuery("_id", idList...)

	result, err := global.EsClient.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).Size(1000).Do(context.Background())

	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	fmt.Println(result.Hits.Hits)
	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)
		if err != nil {
			global.Log.Error(err)
			continue
		}
		//更新
		err = es_ser.ArticleUpdate(hit.Id, map[string]any{
			"collects_count": article.CollectsCount - 1,
		})
		if err != nil {
			global.Log.Error(err.Error())
			continue
		}
	}

	if len(articleIDList) == 0 {
		res.OKWithMessage("请求非法", c)
		return
	}
	err = global.DB.Delete(&collects).Error
	if err != nil {
		global.Log.Error("删除文章时失败", err)
		res.OKWithMessage("删除文章时失败", c)
		return
	}

	res.OKWithMessage(fmt.Sprintf("成功取消收藏 %d", len(articleIDList)), c)
}
