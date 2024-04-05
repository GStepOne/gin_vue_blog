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
