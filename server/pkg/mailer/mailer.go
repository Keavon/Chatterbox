package mailer

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

// SendMail sends emails to an address with the contents message.
// It is configured by the following environmental variables:
// MAILER_HOST: SMTP server address
// MAILER_USERNAME: Email username
// MAILER_PASSWORD: Email password
// MAILER_SMTP_PORT: SMPT port to send emails over (usually port 25)
// MAILER_FROM: Address to send mail from
func SendMail(to, subj, msg string) error {
	msg = fmt.Sprintf("To: %s\r\n"+
		"Subject: %s\r\n"+
		"\r\n"+
		"%s", to, subj, strings.Replace(msg, "\n", "\r\n", -1))

	auth := smtp.PlainAuth("", os.Getenv("MAILER_USERNAME"), os.Getenv("MAILER_PASSWORD"),
		os.Getenv("MAILER_HOST"))
	host := os.Getenv("MAILER_HOST") + ":" + os.Getenv("MAILER_SMTP_PORT")
	fmt.Println(host)
	err := smtp.SendMail(host, auth, os.Getenv("MAILER_FROM"), []string{to}, []byte(msg))
	return err
}
