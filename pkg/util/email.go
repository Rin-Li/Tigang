package util

import (
	"Tigang/conf"
	"fmt"
	"gopkg.in/gomail.v2"
)

func SendEmail(to string, subject string, body string) error {
	cfg := conf.InitEmail()
	
	m := gomail.NewMessage()
	m.SetHeader("From", cfg.SmtpEmail)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(cfg.SmtpHost, cfg.SmtpPort, cfg.SmtpEmail, cfg.SmtpPass)

	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Send email failed: ", err)
		return err
	}

	fmt.Println("Send email success")
	return nil
}