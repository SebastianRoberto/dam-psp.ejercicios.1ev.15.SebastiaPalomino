package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

type Resultado struct {
    nombreArchivo string
    cuenta        int
}

func contarLetra(nombreArchivo string, letra string, ch chan<- Resultado) {
    contenido, err := ioutil.ReadFile(nombreArchivo)
    if err != nil {
        ch <- Resultado{nombreArchivo, -1} // Usamos -1 para indicar un error
        return
    }
    texto := string(contenido)
    cuenta := strings.Count(strings.ToLower(texto), strings.ToLower(letra))
    ch <- Resultado{nombreArchivo, cuenta}
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

    ch := make(chan Resultado)
    total := 0

    for _, nombreArchivo := range os.Args[2:] {
        go contarLetra(nombreArchivo, letra, ch)
    }

    resultados := make([]Resultado, len(os.Args[2:]))
    for i := range os.Args[2:] {
        resultados[i] = <-ch
    }
    close(ch)

    for _, res := range resultados {
        if res.cuenta == -1 {
            fmt.Printf("Error al abrir el archivo: %s\n", res.nombreArchivo)
        } else {
            total += res.cuenta
            fmt.Printf("La letra '%s' aparece %d veces en el archivo %s.\n", letra, res.cuenta, res.nombreArchivo)
        }
    }

    fmt.Printf("Total de apariciones de la letra '%s': %d\n", letra, total)
}


