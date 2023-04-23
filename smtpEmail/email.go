package smtpEmail

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

func SendEmail(email *Email) {

	m := gomail.NewMessage()
	//发送人
	m.SetHeader("From", email.From)
	//接受人
	//m.SetHeader("To", "huangzhen7484@gmail.com")
	m.SetHeader("To", email.To...)

	//	抄送人
	//m.SetAddressHeader("Cc", "389403710@qq.com", "389403710@qq.com")

	//主题
	m.SetHeader("Subject", email.Msg.Subject)
	//内容
	m.SetBody("text/html", email.Msg.Body)

	//	附件
	m.Attach(email.Msg.Attach)

	//连接qq的smtp服务器（发送邮件的服务器），第二个参数是ssl端口号（465或者587）第四个参数是授权码
	d := gomail.NewDialer(email.SMTPHost, email.SMTPPort, email.From, email.Password)

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("smtpEmail send success!")
}
