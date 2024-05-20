package es_ser

import (
	"blog/gin/global"
	"blog/gin/models"
	"bytes"
	"github.com/olivere/elastic/v7"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// 自定义Transport
type LoggingTransport struct {
	Transport http.RoundTripper
}

func (t *LoggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// 记录请求体
	if req.Body != nil {
		var buf bytes.Buffer
		tee := io.TeeReader(req.Body, &buf)
		req.Body = ioutil.NopCloser(tee)
		bodyBytes, _ := ioutil.ReadAll(&buf)
		log.Printf("Request Body: %s", string(bodyBytes))
	}

	// 执行请求
	resp, err := t.Transport.RoundTrip(req)

	// 记录响应体
	if resp != nil && resp.Body != nil {
		var buf bytes.Buffer
		tee := io.TeeReader(resp.Body, &buf)
		resp.Body = ioutil.NopCloser(tee)
		bodyBytes, _ := ioutil.ReadAll(&buf)
		log.Printf("Response Body: %s", string(bodyBytes))
	}

	return resp, err
}

func createEsClientWithLogging() (*elastic.Client, error) {
	transport := &LoggingTransport{
		Transport: http.DefaultTransport,
	}
	client, err := elastic.NewClient(
		elastic.SetHttpClient(&http.Client{Transport: transport}),
		elastic.SetURL("http://localhost:9200"), // 替换为你的Elasticsearch地址
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func main() {
	client, err := createEsClientWithLogging()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// 设置全局Elasticsearch客户端
	global.EsClient = client

	// 这里调用你的函数，例如：CommonList
	option := Option{
		//Key:    "example",
		Fields: []string{"title", "content"},
		PageView: models.PageView{
			Page:  1,
			Limit: 10,
		},
		//Tag:  "sample-tag",
		//Sort: "created_at desc",
	}

	list, count, err := CommonList(option)
	if err != nil {
		log.Fatalf("Error in CommonList: %s", err)
	}
	log.Printf("Total Count: %d", count)
	log.Printf("Articles: %v", list)
}
