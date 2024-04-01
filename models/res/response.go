package res

import (
	"blog/gin/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS = 0
	ERROR   = 1
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

type ListResponse[T any] struct {
	Count int64 `json:"count"`
	List  T     `json:"list"`
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
	Result(SUCCESS, map[string]any{}, msg, c)
}

func Fail(data any, msg string, c *gin.Context) {
	Result(ERROR, data, msg, c)
}

func FailWithMessage(msg string, c *gin.Context) {
	Result(ERROR, map[string]any{}, msg, c)
}

func FailWithCode(code ErrorCode, c *gin.Context) {
	msg, ok := ErrorMap[code]
	if ok {
		Result(int(code), map[string]any{}, msg, c)
		return
	}
	Result(ERROR, map[string]any{}, "未知错误", c)
}

func FailWithError(err error, request any, c *gin.Context) {
	errString := utils.GetValidMsg(err, request)
	//GetValidMsg
	Result(ERROR, map[string]any{}, errString, c)
}

func OkWithList(list any, count int64, c *gin.Context) {
	OKWithData(ListResponse[any]{
		List:  list,
		Count: count,
	}, c)
}
