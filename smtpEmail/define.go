package smtpEmail

type Email struct {
	SMTPHost string   //smtp server Host;like smtp.qq.com
	SMTPPort int      // smtp server Host Port;like smtp.qq.com expose port 465 or 587
	From     string   //sender smtpEmail
	Password string   //authorization code for smtp server
	To       []string // receivers smtpEmail
	Msg      *Message //smtpEmail msg body
}
type Message struct {
	Subject string //smtpEmail subject
	Body    string //smtpEmail body
	Attach  string //smtpEmail attach file
}
