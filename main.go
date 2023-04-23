package main

import (
	"email/email"
	"fmt"
)

func main() {
	fmt.Println("main")
	//email.SendEmail()
	receivers := []string{"389403710@qq.com"}
	subject := "我是主题"
	body := "你好呀，很高兴见到你，我是body"
	attach := "./hello.txt"
	msg := &email.Message{
		Subject: subject,
		Body:    body,
		Attach:  attach,
	}
	emailInfo := &email.Email{
		SMTPHost: "smtp.qq.com",
		SMTPPort: 587,
		From:     "389403710@qq.com",
		Password: "123456",
		To:       receivers,
		Msg:      msg,
	}
	email.SendEmail(emailInfo)
}
