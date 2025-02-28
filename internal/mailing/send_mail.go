package mailing

import (
	"fmt"
	"log"

	"github.com/go-gomail/gomail"
	"github.com/rikidas/go-email-service/config"
)

func Send_mail(to string, subject string, message string) {

	fmt.Println("Inciando")
	m := gomail.NewMessage()

	conf, err := config.Load_config()
	if err != nil {
		log.Panic("Error al cargar configuracion: ")
	}

	m.SetHeader("From", conf.USER)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", message)

	d := gomail.NewDialer(conf.SMTP, conf.PORT, conf.USER, conf.PASS)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}
