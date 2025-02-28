package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type configuration struct {
	SMTP         string
	PORT         int
	USER         string
	PASS         string
	LOG_FILENAME string
}

func Load_config() (configuration, error) {

	err := godotenv.Load("../.env") //Para test
	//err := godotenv.Load(".env")

	if err != nil {
		log.Println("No se pudo leer el archivo de configuracion")
	}

	smtpPortStr := os.Getenv("SMTP_PORT")

	// Convertimos a int
	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		log.Fatalf("Error al convertir SMTP_PORT a int: %v", err)
	}

	return configuration{
		SMTP:         os.Getenv("SMTP_HOST"),
		PORT:         smtpPort,
		USER:         os.Getenv("SMTP_USER"),
		PASS:         os.Getenv("SMTP_PASS"),
		LOG_FILENAME: os.Getenv("LOG_FILENAME"),
	}, err

}
