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
        ch <- 0
        return
    }
    texto := string(contenido)
    cuenta := strings.Count(texto, "a") + strings.Count(texto, "A")
    ch <- cuenta
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Uso: go run main.go <letra> <nombreArchivo1> <nombreArchivo2> ...")
        return
    }

    letra := os.Args[1]
    ch := make(chan int)
    total := 0

    for _, nombreArchivo := range os.Args[2:] {
        go contarLetraA(nombreArchivo, ch)
    }

    for range os.Args[2:] {
        cuenta := <-ch
        total += cuenta
        if cuenta == 0 {
            fmt.Printf("Error al abrir el archivo: %s\n", nombreArchivo)
        } else {
            fmt.Printf("La letra '%s' aparece %d veces en el archivo %s.\n", letra, cuenta, nombreArchivo)
        }
    }

    fmt.Printf("Total de apariciones de la letra '%s': %d\n", letra, total)
}
