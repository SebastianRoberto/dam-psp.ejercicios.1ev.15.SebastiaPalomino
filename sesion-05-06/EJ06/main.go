package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

func contarLetra(nombreArchivo string, letra string, ch chan<- int) {
    contenido, err := ioutil.ReadFile(nombreArchivo)
    if err != nil {
        ch <- -1 // Usamos -1 para indicar un error
        return
    }
    texto := strings.ToLower(string(contenido))
    letraBuscada := strings.ToLower(letra)
    cuenta := strings.Count(texto, letraBuscada)
    ch <- cuenta
}

func main() {
    if len(os.Args) < 3 {
        fmt.Println("Uso: go run main.go <letra> <nombreArchivo1> <nombreArchivo2> ...")
        return
    }

    letra := os.Args[1]
    if len(letra) != 1 {
        fmt.Println("Por favor, ingrese una sola letra como primer argumento.")
        return
    }

    ch := make(chan int)
    total := 0

    for _, nombreArchivo := range os.Args[2:] {
        go contarLetra(nombreArchivo, letra, ch)
    }

    for i, nombreArchivo := range os.Args[2:] {
        cuenta := <-ch
        if cuenta == -1 {
            fmt.Printf("Error al abrir el archivo: %s\n", nombreArchivo)
        } else {
            total += cuenta
            fmt.Printf("La letra '%s' aparece %d veces en el archivo %s.\n", letra, cuenta, nombreArchivo)
        }
        
        if i == len(os.Args[2:])-1 {
            close(ch)
        }
    }

    fmt.Printf("Total de apariciones de la letra '%s': %d\n", letra, total)
}



