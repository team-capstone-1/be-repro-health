package config

import (
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendMail(email_to, subject, template string) error {
    mailer := gomail.NewMessage()
    mailer.SetHeader("From", SENDER_NAME)
    mailer.SetHeader("To", email_to)
    mailer.SetHeader("Subject", subject)
    mailer.SetBody("text/html", template)

	port, err := strconv.Atoi(SMTP_PORT)
	if err != nil {
		return err
	}

    dialer := gomail.NewDialer(
        SMTP_HOST,
        port,
        AUTH_EMAIL,
        AUTH_PASSWORD,
    )

    err = dialer.DialAndSend(mailer)
    if err != nil {
		return err
    }

	return nil
}
