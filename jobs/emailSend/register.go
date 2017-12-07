package emailSend

import (
	"daemon/tool"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
)

const subject = "Smartdo(Smartdo.io)注册邮箱验证"

type RegisterInfo struct {
	Email string `json:"email"`
	Code  int64  `json:"code"`
}

func Register() {
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
	register := RegisterInfo{}
	json.Unmarshal([]byte(result), &register)
	code := strconv.FormatInt(register.Code, 10)
	if res, err := SendRegisterEmail(register.Email, code); !res {
		fmt.Println(dateString + " | " + register.Email + " | " + code + " | " + "FAIL" + " | " + err.Error())
		return
	}
	fmt.Println(dateString + " | " + register.Email + " | " + code + " | " + "SUCCESS")
	return
}

func SendRegisterEmail(email string, code string) (bool, error) {
	content := "<h3>您好！感谢您注册Smartdo帐号，你的验证码是：<h3>" + code
	return tool.SendEmail(email, content, subject)
}
