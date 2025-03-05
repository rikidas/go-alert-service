package mailing

import (
	"fmt"

	"github.com/go-gomail/gomail"
)

func Send_mail(to string, subject string, message string, user string, pass string, smtp string, port int) {

	fmt.Println("Inciando")
	m := gomail.NewMessage()

	m.SetHeader("From", user)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", message)

	d := gomail.NewDialer(smtp, port, user, pass)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}
