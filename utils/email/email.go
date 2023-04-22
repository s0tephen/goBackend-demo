package email

import (
	"crypto/tls"
	"github.com/jordan-wright/email"
	"math/rand"
	"net/smtp"
)

func TestSendMail(ToEmail string) error {
	e := email.NewEmail()

	mailUserName := "w2477284535@163.com" //邮箱账号
	mailPassword := "FBBUWAVMRDHYSNZC"    //邮箱授权码
	code := RandCode()                    //发送的验证码
	Subject := "验证码发送测试"                  //发送的主题

	e.From = "B盘 <w2477284535@163.com>"
	e.To = []string{ToEmail}
	e.Subject = Subject
	e.HTML = []byte("<p>你好！</p>你的验证码为：<h1>" + code + "</h1>" + "<p>***该验证码5分钟内有效***</p>")
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", mailUserName, mailPassword, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		return err
	}
	return nil
}

func RandCode() string {
	s := "1234567890"
	code := ""
	for i := 0; i < 6; i++ {
		code += string(s[rand.Intn(len(s))])
	}
	return code
}
