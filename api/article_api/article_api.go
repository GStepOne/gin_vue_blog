package article_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/ctype"
	"blog/gin/models/res"
	"blog/gin/utils/jwt"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
	"math/rand"
	"strings"
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
	unsafe := blackfriday.MarkdownCommon([]byte(cr.Content))
	//是否存在script 标签
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(unsafe)))
	nodes := doc.Find("script").Nodes
	if len(nodes) > 0 {
		//有script标签
		doc.Find("script").Remove()
		converter := md.NewConverter("", true, nil)
		html, _ := doc.Html()
		markdown, _ := converter.ConvertString(html)
		cr.Content = markdown
	}
	if cr.Abstract == "" {
		//从内容里面去选择30个字符
		abs := []rune(cr.Content) //汉字的截取不一样
		//将content转义为html，并且过滤xss，以及
		if len(abs) > 100 {
			cr.Abstract = string(abs[:100]) //截取0-100
		} else {
			cr.Abstract = string(abs) //小于100所有内容
		}
	}

	//查banner_id下的banner_url
	var bannerUrl string
	var bannerIdList []uint
	global.DB.Model(models.BannerModel{}).Select("id").Scan(&bannerIdList)
	if len(bannerIdList) == 0 {
		res.FailWithMessage("没有banner数据", c)
		return
	}

	rand.Seed(time.Now().UnixNano())
	cr.BannerID = bannerIdList[rand.Intn(len(bannerIdList))] //从bannerid里面随机获取一个
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
		Keyword:      cr.Title,
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

	if article.IsExistsData() {
		global.Log.Error(err)
		res.FailWithMessage("文章标题重复", c)
		return
	}

	article.Create()

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}

	res.OKWithMessage("文章发布成功", c)
}
