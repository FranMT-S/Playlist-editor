package helpers

import (
	"bufio"
	"fmt"
	"os"
)

func NewLog(name string, lines []string) {

	// Crear el nombre del archivo usando la fecha y hora (formato: YYYY-MM-DD_HH-MM-SS.txt)
	filename := name + ".txt"

	// Abrir el archivo en modo de escritura (crear si no existe, agregar al final si existe)
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error al abrir el archivo: %v\n", err)
	}
	defer file.Close()

	// Crear un escritor para agregar líneas al archivo
	writer := bufio.NewWriter(file)

	for _, line := range lines {
		_, err := writer.WriteString(line + "\n") // Escribe la línea con un salto de línea
		if err != nil {
			fmt.Printf("\nError al escribir e	n el archivo: %v\n", err)
		}

	}

	// Asegurar que todo se escriba en el archivo
	if err := writer.Flush(); err != nil {
		fmt.Printf("Error al finalizar la escritura: %v\n", err)
	}

}
