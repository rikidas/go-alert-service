package logging

import (
	"fmt"
	"os"
	"time"
)

// SaveLog guarda el log en un archivo con timestamp
func SaveLog(log_string string, log_filename string) {

	// Crear o abrir el archivo de logs
	file, err := os.OpenFile(log_filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error al abrir el archivo de logs:", err)
		return
	}
	defer file.Close()

	// Formatear el log con fecha y hora
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logEntry := fmt.Sprintf("[%s] %s\n", timestamp, log_string)

	// Escribir en el archivo
	if _, err := file.WriteString(logEntry); err != nil {
		fmt.Println("Error al escribir en el archivo de logs:", err)
	}
}
