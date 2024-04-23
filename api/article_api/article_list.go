package article_api

import (
	"blog/gin/models"
	"blog/gin/models/res"
	"blog/gin/service/es_ser"
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
)

type ArticleSearchRequest struct {
	models.PageView
	Tag string `json:"tag" query:"tag" form:"tag"`
}

func (ArticleApi) ArticleListView(c *gin.Context) {
	var cr ArticleSearchRequest

	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	list, count, err := es_ser.CommonList(es_ser.Option{
		PageView: cr.PageView,
		Tag:      cr.Tag,
		Fields:   []string{"title", "abstract", "content"},
	})

	if err != nil {
		res.FailWithMessage("文章列表为空", c)
		return
	}
	NewList := filter.Omit("list", list)
	_list, _ := NewList.(filter.Filter)
	//如果它为空
	if string(_list.MustMarshalJSON()) == "{}" {
		list = make([]models.ArticleModel, 0)
		res.OkWithList(list, int64(count), c)
		return
	}
	res.OkWithList(NewList, int64(count), c)
}
