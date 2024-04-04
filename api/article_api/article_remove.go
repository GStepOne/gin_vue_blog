package article_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"blog/gin/service/es_ser"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

type ArticleIDRequest struct {
	IdList []string `json:"id_list"`
}

func (ArticleApi) ArticleRemoveView(c *gin.Context) {
	var cr ArticleIDRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error()
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	//如果文章删除了，用户收藏这篇文章怎么办
	// 顺带把与这个文章关联的收藏也删除了
	// 标记文章删除，不让收藏的继续看
	bulkService := global.EsClient.Bulk().Index(models.ArticleModel{}.Index()).Refresh("true")
	for _, id := range cr.IdList {
		req := elastic.NewBulkDeleteRequest().Id(id)
		bulkService.Add(req)
		//把全文索引也删掉
		go es_ser.DeleteFullTextByArticleId(id)
	}

	result, err := bulkService.Do(context.Background())

	if err != nil {
		logrus.Error("删除失败", err.Error())
		res.FailWithMessage("删除失败", c)
		return
	}

	res.OKWithMessage(fmt.Sprintf("成功删除%d篇文章", len(result.Succeeded())), c)
}
