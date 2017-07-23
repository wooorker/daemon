package main

import (
	"fmt"
	"net/smtp"
	"net/textproto"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/jordan-wright/email"
)

// * * * * *
// minute 0-59
// hour 0-23
// day 1-31
// month 1-12
// week 1-7

func main() {
	for {
		select {
		case <-time.After(1 * time.Second):
			go consume()
		}
	}
}

func consume() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		return
	}
	res, err := c.Do("RPOP", "QUEUE:ACCOUNT_REGISTER")
	if err != nil {
		return
	}
	if res == nil {
		dateString := time.Now().Format("2006-01-02 08:00:01")
		fmt.Println("[" + dateString + "]" + " empty data")
		return
	}
	res, _ = redis.String(res, err)
	sendEmail(res.(string))
}

func sendEmail(emailAccont string) {
	e := &email.Email{
		To:      []string{emailAccont},
		From:    "Smartdo <15522634982@163.com>",
		Subject: "Smartdo(Smartdo.io)注册邮箱验证",
		HTML:    []byte("<h3>您好！感谢您注册Smartdo帐号，点击下面的链接即可完成激活：</h3><br><a href='http://localhost:8888/'>http://localhost:8888/</a>"),
		Headers: textproto.MIMEHeader{},
	}
	err := e.Send("smtp.163.com:25", smtp.PlainAuth("", "15522634982@163.com", "shizhan214", "smtp.163.com"))
	if err == nil {
		fmt.Println("success")
	}
}
