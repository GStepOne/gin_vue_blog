package requests

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
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

func Get(fullURI string, data interface{}, headers map[string]interface{}, timeout time.Duration) (*http.Response, error) {
	// 将数据转换为反射值
	v := reflect.ValueOf(data)

	// 确保传入的是结构体类型
	if v.Kind() != reflect.Struct {
		return nil, errors.New("data must be a struct")
	}
	fullURI += "?"
	// 遍历结构体的属性
	for i := 0; i < v.NumField(); i++ {
		// 获取结构体字段的名称和值
		fieldName := v.Type().Field(i).Name
		fieldValue := v.Field(i)

		// 将字段名称和值添加到 URI 中
		fullURI += fmt.Sprintf("&%s=%v", fieldName, fieldValue.Interface())
	}

	fmt.Println(fullURI)

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
