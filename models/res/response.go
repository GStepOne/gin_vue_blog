package res

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS = 0
	ERROR   = 7
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func OK(data any, msg string, c *gin.Context) {
	Result(SUCCESS, data, msg, c)
}

func OKWith(c *gin.Context) {
	Result(SUCCESS, map[string]any{}, "成功", c)
}

func OKWithData(data any, c *gin.Context) {
	Result(SUCCESS, data, "success", c)
}

func OKWithMessage(msg string, c *gin.Context) {
	Result(SUCCESS, map[string]any{}, "success", c)
}

func Fail(data any, msg string, c *gin.Context) {
	Result(ERROR, data, msg, c)
}

func FailWithMessage(msg string, c *gin.Context) {
	Result(ERROR, map[string]any{}, "success", c)

}

func FailWithCode(code ErrorCode, c *gin.Context) {
	msg, ok := ErrorMap[code]
	if ok {
		Result(int(code), map[string]any{}, msg, c)
		return
	}
	Result(ERROR, map[string]any{}, "未知错误", c)
}
