package main

import (
	"email/smtpEmail"
	"fmt"
)

func main() {
	fmt.Println("main")
	//smtpEmail.SendEmail()
	receivers := []string{"389403710@qq.com"}
	subject := "我是主题"
	body := "你好呀，很高兴见到你，我是body"
	attach := "./hello.txt"
	msg := &smtpEmail.Message{
		Subject: subject,
		Body:    body,
		Attach:  attach,
	}
	emailInfo := &smtpEmail.Email{
		SMTPHost: "smtp.qq.com",
		SMTPPort: 587,
		From:     "389403710@qq.com",
		Password: "sybagzuwgjstbjgi",
		To:       receivers,
		Msg:      msg,
	}
	smtpEmail.SendEmail(emailInfo)

	//err := shortMsg.Send(tea.StringSlice(os.Args[1:]))
	//if err != nil {
	//	fmt.Println(err)
	//	panic(err)
	//}
	//fmt.Println("success")

}
