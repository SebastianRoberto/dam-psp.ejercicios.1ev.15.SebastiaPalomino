package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

func contarLetraA(nombreArchivo string, ch chan<- int) {
    contenido, err := ioutil.ReadFile(nombreArchivo)
    if err != nil {
        ch <- -1 // Usamos -1 para indicar un error
        return
    }
    texto := strings.ToLower(string(contenido))
    cuenta := strings.Count(texto, "a")
    ch <- cuenta
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Uso: go run main.go <nombreArchivo1> <nombreArchivo2> ...")
        return
    }

    ch := make(chan int)
    total := 0

    for _, nombreArchivo := range os.Args[1:] {
        go contarLetraA(nombreArchivo, ch)
    }

    for i, nombreArchivo := range os.Args[1:] {
        cuenta := <-ch
        if cuenta == -1 {
            fmt.Printf("Error al abrir el archivo: %s\n", nombreArchivo)
        } else {
            total += cuenta
            fmt.Printf("La letra 'a' aparece %d veces en el archivo %s.\n", cuenta, nombreArchivo)
        }
        
        if i == len(os.Args[1:])-1 {
            close(ch)
        }
    }

    fmt.Printf("Total de apariciones de la letra 'a': %d\n", total)
}

