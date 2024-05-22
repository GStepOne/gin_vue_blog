package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

// WordPress represents the root element of the RSS feed.
type WordPress struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

// Channel contains the channel element of the RSS feed.
type Channel struct {
	Title string `xml:"title"`
	Items []Item `xml:"item"`
}

// Item represents each item element in the RSS feed.
type Item struct {
	Title       string   `xml:"title"`
	Content     string   `xml:"http://purl.org/rss/1.0/modules/content/ encoded"`
	Link        string   `xml:"link"`
	PubDate     string   `xml:"pubDate"`
	Creator     string   `xml:"http://purl.org/dc/elements/1.1/ creator"`
	Category    []string `xml:"category"`
	Description string   `xml:"description"`
	CreatedAt   string   `xml:"created_at"`
	UpdatedAt   string   `xml:"updated_at"`
}

func main() {
	xmlFile, err := os.Open("/Users/dudu/www/golang/src/blog/gin/testadata/7-2.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)
	var wordpress WordPress
	err = xml.Unmarshal(byteValue, &wordpress)
	if err != nil {
		fmt.Println("Error unmarshalling:", err)
		return
	}

	for k, item := range wordpress.Channel.Items {
		fmt.Println("Title:", item.Title)
		fmt.Println("序号:", k)
		//fmt.Println("Content:", item.Content)
		//fmt.Println("Link:", item.Link)
		//fmt.Println("Publication Date:", item.PubDate)
		//fmt.Println("Creator:", item.Creator)
		//fmt.Println("Categories:", strings.Join(item.Category, ", "))
		//fmt.Println("Description:", item.Description)
		//fmt.Println()
	}

	saveToElasticsearch(wordpress.Channel.Items)
}

func saveToElasticsearch(items []Item) {
	var err error
	sniffOpt := elastic.SetSniff(false)
	host := "http://127.0.0.1:9200"
	es, err := elastic.NewClient(
		elastic.SetURL(host),
		sniffOpt,
		elastic.SetBasicAuth("", ""),
	)
	if err != nil {
		logrus.Fatalf("es链接失败: %s", err.Error())
	}

	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 生成 1 到 10000 之间的随机整数

	for _, item := range items {
		collectsCount := rand.Intn(10) + 1
		lookCount := rand.Intn(1000) + 1
		diggCount := rand.Intn(100) + 1
		doc := map[string]interface{}{
			"title":          item.Title,
			"content":        item.Content,
			"link":           item.Link,
			"created_at":     formatDate(item.PubDate),
			"updated_at":     formatDate(item.PubDate),
			"user_id":        2, // Placeholder, adjust as needed
			"user_nick_name": "Jack Lee",
			"category":       strings.Join(item.Category, ", "),
			"tags":           item.Category, // Placeholder, adjust as needed
			"keyword":        "",            // Placeholder, adjust as needed
			"abstract":       item.Description,
			"look_count":     lookCount,                       // Placeholder, adjust as needed
			"comment_count":  0,                               // Placeholder, adjust as needed
			"digg_count":     diggCount,                       // Placeholder, adjust as needed
			"collects_count": collectsCount,                   // Placeholder, adjust as needed
			"banner_id":      0,                               // Placeholder, adjust as needed
			"banner_url":     "/uploads/file/WechatIMG35.jpg", // Placeholder, adjust as needed
			"source":         "",                              // Placeholder, adjust as needed
		}

		req := elastic.NewBulkIndexRequest().Index("article_search").Doc(doc)
		bulkRequest := es.Bulk().Add(req)

		bulkResponse, err := bulkRequest.Do(context.Background())
		if err != nil {
			log.Println("Error indexing item:", err)
		}

		if bulkResponse.Errors {
			log.Println("Bulk request contains errors")
			for _, failed := range bulkResponse.Failed() {
				log.Printf("Error: %v", failed.Error)
			}
		}
	}
}

func formatDate(pubDate string) string {
	t, err := time.Parse(time.RFC1123Z, pubDate)
	if err != nil {
		log.Printf("Error parsing date: %v", err)
		return time.Now().Format("2006-01-02 15:04:05")
	}
	return t.Format("2006-01-02 15:04:05")
}
