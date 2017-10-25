package main

import (
	"encoding/json"
	"fmt"
	"net/smtp"
	"net/textproto"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/jordan-wright/email"
)

// var registerUrl string = "http://localhost:8888/v1/user/activate"

var registerUrl string = "http://api.tigerb.cn/v1/user/activate"

type RegisterInfo struct {
	Email  string `json:"email"`
	SToken string `json:"s_token"`
}

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
	dateString := time.Unix(time.Now().Unix(), 0).String()
	if res == nil {
		fmt.Println(dateString + " | empty data")
		return
	}
	result, _ := redis.String(res, err)
	// fmt.Println(result)
	register := RegisterInfo{}
	json.Unmarshal([]byte(result), &register)
	// fmt.Println(register.Email)
	if res := sendEmail(register.Email, register.SToken); !res {
		fmt.Println(dateString + " | " + register.Email + " | " + register.SToken + " | " + "FAIL")
		return
	}
	fmt.Println(dateString + " | " + register.Email + " | " + register.SToken + " | " + "SUCCESS")
	return
}

func sendEmail(emailAccont string, sToken string) bool {
	wholeUrl := registerUrl + "?s_token=" + sToken
	e := &email.Email{
		To:      []string{emailAccont},
		From:    "Smartdo <15522634982@163.com>",
		Subject: "Smartdo(Smartdo.io)注册邮箱验证",
		HTML:    []byte("<h3>您好！感谢您注册Smartdo帐号，点击下面的链接即可完成激活：</h3><br><a href='" + wholeUrl + " +'>" + wholeUrl + "</a>"),
		Headers: textproto.MIMEHeader{},
	}
	err := e.Send("smtp.163.com:25", smtp.PlainAuth("", "15522634982@163.com", "shizhan214", "smtp.163.com"))
	if err == nil {
		return true
	}
	return false
}
