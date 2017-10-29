package tool

import (
	"net/smtp"
	"net/textproto"

	"github.com/jordan-wright/email"
)

// SMTPIdentity identity
const SMTPAddr = "smtp.163.com:25"

// SMTPAccount account
const SMTPAccount = "15522634982@163.com"

// SMTPAccountPwd password
const SMTPAccountPwd = "shizhan214"

// SMTPHost host
const SMTPHost = "smtp.163.com"

func SendEmail(emailAccont string, content string, subject string) (bool, error) {
	e := &email.Email{
		To:   []string{emailAccont},
		From: "Smartdo <15522634982@163.com>",
		// From:    "Smartdo <noreply@smartdo.io>",
		Subject: subject,
		HTML:    []byte(content),
		Headers: textproto.MIMEHeader{},
	}
	err := e.Send(SMTPAddr, smtp.PlainAuth("", SMTPAccount, SMTPAccountPwd, SMTPHost))
	if err == nil {
		return true, nil
	}
	return false, err
}
