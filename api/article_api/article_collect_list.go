package article_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"blog/gin/service/common"
	"blog/gin/utils/jwt"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

type CollResponse struct {
	models.ArticleModel
	CreatedAt string `json:"created_at"`
}

func (ArticleApi) ArticleCollectList(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	var cr models.PageView
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	//var collects []models.UserCollectModel
	var articleIDList []interface{}

	list, count, err := common.ComList(models.UserCollectModel{UserID: claims.UserId}, common.Option{
		PageView: cr,
		Debug:    true,
	})

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("获取文章收藏列表错误", c)
		return
	}

	var collMap = map[string]string{}
	for _, model := range list {
		articleIDList = append(articleIDList, model.ArticleID)
		collMap[model.ArticleID] = model.CreatedAt.Format("2006-01-02 15:03:04")
	}

	fmt.Println(articleIDList)

	//精确匹配keyword NewTermsQuery
	boolSearch := elastic.NewTermsQuery("_id", articleIDList...)

	result, err := global.EsClient.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		Size(1000).
		Do(context.Background())

	if err != nil {
		res.FailWithMessage(err.Error(), c)
	}

	var collList = make([]CollResponse, 0)

	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)
		if err != nil {
			global.Log.Error(err)
			continue
		}

		article.ID = hit.Id

		collList = append(collList, CollResponse{
			ArticleModel: article,
			CreatedAt:    collMap[hit.Id],
		})
	}

	fmt.Println(collList)
	res.OkWithList(collList, count, c)

}
