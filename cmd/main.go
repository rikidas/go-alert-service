package main

import (
	"flag"

	"github.com/rikidas/go-email-service/internal/mailing"
)

func main() {

	email := flag.String("email", "", "Correo al que enviar las alertas")
	subject := flag.String("subject", "", "Asunto")
	message := flag.String("message", "", "Mensaje")

	flag.Parse()

	//Mandamos el mail
	mailing.Send_mail(*email, *subject, *message)
}
