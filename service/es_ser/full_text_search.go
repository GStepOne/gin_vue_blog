package es_ser

import (
	"blog/gin/global"
	"blog/gin/models"
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/olivere/elastic/v7"
	"github.com/russross/blackfriday"
	"github.com/sirupsen/logrus"
	"strings"
)

type SearchData struct {
	Body  string `json:"body"`
	Slug  string `json:"slug"` //跳转地址
	Title string `json:"title"`
	//Id    string `json:"id"`
	Key string `json:"key"`
}

func GetSearchIndexContent(id, title string, content string) []SearchData {
	var headList, bodyList []string
	headList = append(headList, GetHeader(title))
	dataList := strings.Split(content, "\n")
	var newContent string
	var isCode = false
	for _, l := range dataList {
		if strings.HasPrefix(l, "```") {
			isCode = !isCode
		}
		if strings.HasPrefix(l, "#") && !isCode {
			headList = append(headList, GetHeader(l))
			bodyList = append(bodyList, GetBody(content))
			content = ""
			continue
		}
		fmt.Println(content)
		fmt.Println(l)
		newContent += l
	}

	fmt.Println("我content的现在内容为", newContent)

	bodyList = append(bodyList, newContent)
	ln := len(headList)
	var searchDataList []SearchData
	for i := 0; i < ln; i++ {
		searchDataList = append(searchDataList, SearchData{
			Title: headList[i],
			Body:  bodyList[i],
			Slug:  id + GetSlug(headList[i]),
			Key:   id,
		})
	}
	return searchDataList
}

func GetHeader(head string) string {
	strings.ReplaceAll(head, "#", "")
	strings.ReplaceAll(head, " ", "")
	return "?id=" + head
}

func GetBody(body string) string {
	unsafe := blackfriday.MarkdownCommon([]byte(body))
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(unsafe)))
	return doc.Text()
}

func GetSlug(slug string) string {
	return "#" + slug
}

func AsyncArticleByFullText(id, title, content string) {
	indexList := GetSearchIndexContent(id, title, content)
	bulk := global.EsClient.Bulk()
	for _, indexData := range indexList {
		req := elastic.NewBulkIndexRequest().Index(models.FullTextModel{}.Index()).Doc(indexData)
		bulk.Add(req)
	}

	result, err := bulk.Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}

	logrus.Infof("%s 添加成功", "共", len(result.Succeeded()), "条")

}

func DeleteFullTextByArticleId(id string) {
	//先删掉再添加
	boolSearch := elastic.NewTermQuery("key", id)
	res, _ := global.EsClient.DeleteByQuery().
		Index(models.FullTextModel{}.Index()).
		Query(boolSearch).
		Do(context.Background())

	logrus.Infof("成功删除 %d 条记录", res.Deleted)
}
