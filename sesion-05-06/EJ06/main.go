package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)


func contarLetraA(nombreArchivo string) (int, error) {
    contenido, err := ioutil.ReadFile(nombreArchivo)
    if err != nil {
        return 0, err 
    }

    texto := string(contenido)
    
    cuenta := strings.Count(texto, "a") + strings.Count(texto, "A")
    return cuenta, nil
}

func main() {
    
    if len(os.Args) < 2 {
        fmt.Println("Uso: go run main.go <nombreArchivo>")
        return
    }

    nombreArchivo := os.Args[1]

    
    cuenta, err := contarLetraA(nombreArchivo)
    if err != nil {
        fmt.Printf("Error al abrir el archivo: %v\n", err)
        return
    }

    fmt.Printf("La letra 'a' aparece %d veces en el archivo %s\n", cuenta, nombreArchivo)
}
