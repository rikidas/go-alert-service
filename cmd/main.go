package main

import (
	"flag"
	"fmt"

	"github.com/rikidas/go-email-service/internal/logging"
	"github.com/rikidas/go-email-service/internal/mailing"
)

func main() {
	email := flag.String("email", "", "Correo al que enviar las alertas")
	subject := flag.String("subject", "", "Asunto")
	message := flag.String("message", "", "Mensaje")
	mode := flag.String("mode", "both", "Modo de ejecución: log, mail o both")

	flag.Parse()

	// Verificamos qué modo de ejecución eligió el usuario
	switch *mode {
	case "log":
		logging.SaveLog(*message)
		fmt.Println("Mensaje guardado en logs.")
	case "mail":
		mailing.Send_mail(*email, *subject, *message)
		fmt.Println("Correo enviado.")
	case "both":
		logging.SaveLog(*message)
		mailing.Send_mail(*email, *subject, *message)
		fmt.Println("Mensaje guardado en logs y correo enviado.")
	default:
		fmt.Println("Modo inválido. Usa 'log', 'mail' o 'both'.")
	}
}
