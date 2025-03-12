package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/rikidas/go-email-service/config"
	"github.com/rikidas/go-email-service/internal/logging"
	"github.com/rikidas/go-email-service/internal/mailing"
)

func main() {
	email := flag.String("email", "", "Correo al que enviar las alertas")
	subject := flag.String("subject", "", "Asunto")
	message := flag.String("message", "", "Mensaje")
	mode := flag.String("mode", "both", "Modo de ejecución: log, mail o both")
	ruta_env := flag.String("ruta_env", "", "Ruta donde se encuentra el .env con las claves")
	log_filename := flag.String("log_filename", "", "Ruta donde se guardaran los logs")

	flag.Parse()

	// Verificamos qué modo de ejecución eligió el usuario
	switch *mode {
	case "log":
		logging.SaveLog(*message, *log_filename)
		fmt.Println("Mensaje guardado en logs.")
	case "mail":

		conf, err := config.Load_config(*ruta_env)
		if err != nil {
			log.Panic("Error al cargar configuracion: ")
		}
		mailing.Send_mail(*email, *subject, *message, conf.USER, conf.PASS, conf.SMTP, conf.PORT)
		fmt.Println("Correo enviado.")
	case "both":

		conf, err := config.Load_config(*ruta_env)
		if err != nil {
			log.Panic("Error al cargar configuracion: ")
		}

		logging.SaveLog(*message, *log_filename)
		mailing.Send_mail(*email, *subject, *message, conf.USER, conf.PASS, conf.SMTP, conf.PORT)
		fmt.Println("Mensaje guardado en logs y correo enviado.")
	default:
		fmt.Println("Modo inválido. Usa 'log', 'mail' o 'both'.")
	}
}
