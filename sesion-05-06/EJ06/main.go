package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)


func contarLetraA(nombreArchivo string, ch chan int) {
    contenido, err := ioutil.ReadFile(nombreArchivo)
    if err != nil {
        fmt.Printf("Error al abrir el archivo: %v\n", err)
        ch <- 0
        return
    }

    texto := string(contenido)
    cuenta := strings.Count(texto, "a") + strings.Count(texto, "A")
    ch <- cuenta 
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Uso: go run main.go <nombreArchivo>")
        return
    }

    nombreArchivo := os.Args[1]
    ch := make(chan int)

    
    go contarLetraA(nombreArchivo, ch)

    
    cuenta := <-ch
    fmt.Printf("La letra 'a' aparece %d veces en el archivo %s.\n", cuenta, nombreArchivo)
}

