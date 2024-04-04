package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
	"strings"
)

type SearchData struct {
	Body  string `json:"body"`
	Slug  string `json:"slug"` //跳转地址
	Title string `json:"title"`
}

func main() {
	var data = "##第一个标题\n嘻嘻嘻阿珂记录到撒酒疯\n  ```\n\n\n 创建docker的代码 import ``` ##第二个标题\n拉萨看见了大开杀戒咖啡机\n"
	GetSearchIndexContent("/article/sadf2", data)

}

func GetSearchIndexContent(title string, content string) []SearchData {
	var headList, bodyList []string
	headList = append(headList, GetHeader(title))
	dataList := strings.Split(content, "\n")
	var isCode = false
	for _, l := range dataList {
		if strings.HasPrefix(l, "```") {
			isCode = !isCode
		}
		if strings.HasPrefix(l, "#") && !isCode {
			headList = append(headList, GetHeader(l))
			content = GetBody(content)
			bodyList = append(bodyList, content)
			content = ""
			continue
		}
		content += l
	}

	bodyList = append(bodyList, content)

	ln := len(headList)
	var searchDataList []SearchData
	for i := 0; i < ln; i++ {
		searchDataList = append(searchDataList, SearchData{
			Title: headList[i],
			Body:  bodyList[i],
			Slug:  title + GetSlug(headList[i]),
		})
	}

	b, _ := json.Marshal(searchDataList)
	fmt.Println(b)

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
