package requests

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// "https://api.codelife.cc/api/top/list"
func Post(url string, data any, headers map[string]interface{}, timeout time.Duration) (*http.Response, error) {
	reqParam, _ := json.Marshal(data)
	reqBody := strings.NewReader(string(reqParam))
	httpReq, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Add("Content-Type", "application/json")

	//httpReq.Header.Add("signaturekey", "U2FsdGVkX1/dBVj3thjxNDyIRPGWjvYm94d/e6+ho98=")
	//httpReq.Header.Add("version", "1.3.21")

	for key, val := range headers {
		switch v := val.(type) {
		case string:
			httpReq.Header.Add(key, v)
		case int:
			httpReq.Header.Add(key, strconv.Itoa(v))
		}
		//httpReq.Header.Add(key, val)
	}

	client := http.Client{Timeout: timeout}
	httpRsp, err := client.Do(httpReq)
	return httpRsp, err
}

func Get(fullURI string, headers map[string]interface{}, timeout time.Duration) (*http.Response, error) {
	// 创建 HTTP 请求
	httpReq, err := http.NewRequest("GET", fullURI, nil)
	if err != nil {
		return nil, err
	}

	// 添加请求头
	httpReq.Header.Add("Content-Type", "application/json")
	for key, val := range headers {
		switch v := val.(type) {
		case string:
			httpReq.Header.Add(key, v)
		case int:
			httpReq.Header.Add(key, strconv.Itoa(v))
		}
	}

	// 创建 HTTP 客户端并发送请求
	client := http.Client{Timeout: timeout}
	httpRsp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	return httpRsp, nil
}
