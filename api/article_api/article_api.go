package article_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/ctype"
	"blog/gin/models/res"
	"blog/gin/utils/jwt"
	"github.com/gin-gonic/gin"
	"time"
)

type ArticleRequest struct {
	ID       string `json:"id" binding:"required" msg:"文章id必填"`
	Title    string `json:"title"`
	Abstract string `json:"abstract"` //简介
	Content  string `json:"content" binding:"required" msg:"文章内容必填"`

	Category string `json:"category"`
	Source   string `json:"source"`
	Link     string `json:"link"`

	BannerID uint        `json:"banner_id"`
	Tags     ctype.Array `json:"tags"`
}

func (ArticleApi) ArticleCreateView(c *gin.Context) {
	var cr ArticleRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	userID := claims.UserId
	userNickname := claims.Nickname
	//校验content xss 攻击
	if cr.Abstract == "" {
		//从内容里面去选择30个字符
		abs := []rune(cr.Content) //汉字的截取不一样
		if len(abs) > 100 {
			cr.Abstract = string(abs[:100])
		} else {
			cr.Abstract = string(abs[:])
		}
	}

	//查banner_id下的banner_url

	var bannerUrl string

	err = global.DB.Model(models.BannerModel{}).Where("id =?", cr.BannerID).Select("path").Scan(&bannerUrl).Error
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	var avatar string
	err = global.DB.Model(models.UserModel{}).Where("id =?", userID).Select("avatar").Scan(&avatar).Error
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	now := time.Now().Format("2006-01-02 15:04:05")

	article := models.ArticleModel{
		CreatedAt:    now,
		UpdatedAt:    now,
		Title:        cr.Title,
		Abstract:     cr.Abstract,
		Content:      cr.Content,
		UserId:       userID,
		UserNickName: userNickname,
		UserAvatar:   avatar,
		Category:     cr.Category,
		Source:       cr.Source,
		Link:         cr.Link,
		BannerID:     cr.BannerID,
		BannerUrl:    bannerUrl,
		Tags:         cr.Tags,
	}

	article.Create()

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}

	res.OKWithMessage("文章发布成功", c)
}
