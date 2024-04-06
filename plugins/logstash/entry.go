package logstash

import (
	"blog/gin/global"
	"blog/gin/utils"
	"blog/gin/utils/jwt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Log struct {
	IP     string `json:"ip"`
	Token  string `json:"token"`
	Addr   string `json:"addr"`
	UserId uint   `json:"user_id"`
}

var (
	std = new(Log)
)

func New(ip string, token string) *Log {
	claims, err := jwt.ParseToken(token)
	var UserId uint
	if err == nil {
		UserId = claims.UserId
	}
	addr := utils.GetAddr(ip)
	return &Log{
		IP:     ip,
		Token:  token,
		Addr:   addr,
		UserId: UserId,
	}
}

func NewLogByGin(c *gin.Context) *Log {
	ip := c.ClientIP()
	token := c.Request.Header.Get("token")
	return New(ip, token)
}
func (l Log) Debug(content string) {
	l.Send(DEBUG_LEVEL, content)
}

func (l Log) Info(content string) {
	l.Send(INFO_LEVEL, content)
}
func (l Log) Warn(content string) {
	l.Send(WARN_LEVEL, content)
}

func (l Log) Error(content string) {
	l.Send(ERROR_LEVEL, content)
}

func (l Log) SendMsg() {

}

// Send 入库
func (l Log) Send(level Level, content string) {

	err := global.DB.Create(&LogStashModel{
		IP:      l.IP,
		Addr:    l.Addr,
		Level:   level,
		Content: content,
		UserID:  l.UserId,
	}).Error
	if err != nil {
		logrus.Error(err)
	}
}

func Debug(content string) {
	std.Debug(content)
}

func Info(content string) {
	std.Info(content)
}

func Warn(content string) {
	std.Warn(content)
}

func Error(content string) {
	std.Error(content)
}
