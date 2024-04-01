package main

import (
	"blog/gin/core"
	"blog/gin/global"
	"gopkg.in/gomail.v2"
)

type Subject string

const (
	Code  Subject = "平台验证码"
	Note  Subject = "操作通知"
	Alarm Subject = "告警通知"
)

type Api struct {
	Subject Subject
}

func (a Api) Send(name, body string) error {
	return send(name, string(a.Subject), body)
}

func NewCode() Api {
	return Api{
		Subject: Code,
	}
}

func NewNote() Api {
	return Api{
		Subject: Note,
	}
}

func NewAlarm() Api {
	return Api{
		Subject: Alarm,
	}
}

// send 邮件发送，发给谁，主题，正文
func send(name, Subject, body string) error {
	e := global.Config.Email
	return sendMail(
		e.User,
		e.Password,
		e.Host,
		e.Port,
		name,
		e.DefaultFromEmail, Subject,
		body,
	)
}

func sendMail(username, authCode, host string, port int, mailTo, sendName string, Subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(username, sendName))
	m.SetHeader("To", mailTo)
	m.SetHeader("Subject", Subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(host, port, username, authCode)
	err := d.DialAndSend(m)
	return err
}

func main() {
	core.InitCoreConf()
	core.InitLogger()
	NewCode().Send("867700123@qq.com", "验证码是：2379")
}
