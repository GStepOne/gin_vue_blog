package article_api

import (
	"blog/gin/models"
	"blog/gin/models/res"
	"blog/gin/service/es_ser"
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
)

func (ArticleApi) ArticleUpdateView(c *gin.Context) {
	var cr models.PageView

	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, err := es_ser.CommonList(cr.Key, cr.Page, cr.Limit)

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
