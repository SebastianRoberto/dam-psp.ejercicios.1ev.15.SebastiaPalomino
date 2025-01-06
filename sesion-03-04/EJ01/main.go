package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Tienes que poner go run main.go <texto> <archivo_entrada> <archivo_salida>")
		os.Exit(1)
	}

	texto := os.Args[1]
	archivoEntrada := os.Args[2]
	archivoSalida := os.Args[3]

	entrada, err := os.Open(archivoEntrada)
	if err != nil {
		fmt.Printf("Error al abrir el archivo de entrada: %v\n", err)
		os.Exit(1)
	}
	defer entrada.Close()

	salida, err := os.OpenFile(archivoSalida, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("Error al abrir el archivo de salida: %v\n", err)
		os.Exit(1)
	}
	defer salida.Close()

	comando := exec.Command("grep", texto)
	comando.Stdin = entrada
	comando.Stdout = salida
	comando.Stderr = os.Stderr

	if err := comando.Run(); err != nil {
		fmt.Printf("Error al ejecutar el comando grep: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("El filtrado se completó correctamente.")
}
