package mail

import (
	"log"
	"strings"

	"net/smtp"
)

func SendMail(to []string, subject string, body string) error {
	// Set up authentication information.
	auth := smtp.PlainAuth("", "kittabun.sk@gmail.com", "qpmj miit cjpn hshe", "smtp.gmail.com")

	// Compose the email message.
	msg := []byte("To: " + strings.Join(to, ",") + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body)

	// Send the email.
	err := smtp.SendMail("smtp.gmail.com:587", auth, "kittabun.sk@gmail.com", to, msg)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
