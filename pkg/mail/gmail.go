package commons

import (
	"crypto/tls"

	"github.com/armnerd/go-skeleton/config"

	"gopkg.in/gomail.v2"
)

func SendMail(subject string, body string, mailTo ...string) error {
	emailConfig := config.MailConfig
	if mailTo == nil {
		mailTo = config.AlertEmailList
	}
	m := gomail.NewMessage()
	m.SetHeader("From", "<"+emailConfig.User+">")
	m.SetHeader("To", mailTo...)    // 发送给多个用户
	m.SetHeader("Subject", subject) // 设置邮件主题
	m.SetBody("text/html", body)    // 设置邮件正文
	d := gomail.NewDialer(
		emailConfig.Host,
		emailConfig.Port,
		emailConfig.User,
		emailConfig.Password,
	)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err := d.DialAndSend(m)
	return err
}
