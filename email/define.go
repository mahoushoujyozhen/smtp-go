package email

type Email struct {
	SMTPHost string   //smtp server Host;like smtp.qq.com
	SMTPPort int      // smtp server Host Port;like smtp.qq.com expose port 465 or 587
	From     string   //sender email
	Password string   //authorization code for smtp server
	To       []string // receivers email
	Msg      *Message //email msg body
}
type Message struct {
	Subject string //email subject
	Body    string //email body
	Attach  string //email attach file
}
