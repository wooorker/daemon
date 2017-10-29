package emailSend

import (
	"daemon/tool"
	"encoding/json"
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

const subject = "Smartdo(Smartdo.io)注册邮箱验证"

var registerUrl string = "http://api.tigerb.cn/v1/user/activate"

// var registerUrl string = "http://localhost:8888/v1/user/activate"

type RegisterInfo struct {
	Email  string `json:"email"`
	SToken string `json:"s_token"`
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
	// fmt.Println(result)
	register := RegisterInfo{}
	json.Unmarshal([]byte(result), &register)
	// fmt.Println(register.Email)
	if res, err := SendRegisterEmail(register.Email, register.SToken); !res {
		fmt.Println(dateString + " | " + register.Email + " | " + register.SToken + " | " + "FAIL" + " | " + err.Error())
		return
	}
	fmt.Println(dateString + " | " + register.Email + " | " + register.SToken + " | " + "SUCCESS")
	return
}

func SendRegisterEmail(email string, sToken string) (bool, error) {

	wholeUrl := registerUrl + "?s_token=" + sToken
	content := "<h3>您好！感谢您注册Smartdo帐号，点击下面的链接即可完成激活：</h3><br><a href='" + wholeUrl + " +'>" + wholeUrl + "</a>"
	return tool.SendEmail(email, content, subject)
}
