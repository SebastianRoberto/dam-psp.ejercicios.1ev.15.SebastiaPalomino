package main

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "strings"
    "sync"
)

type Resultado struct {
    nombreArchivo string
    cuenta        int
}

func contarLetra(nombreArchivo string, letra string, ch chan<- Resultado) {
    cmd := exec.Command("grep", "-oi", letra, nombreArchivo)
    output, err := cmd.StdoutPipe()
    if err != nil {
        ch <- Resultado{nombreArchivo, -1}
        return
    }

    if err := cmd.Start(); err != nil {
        ch <- Resultado{nombreArchivo, -1}
        return
    }

    scanner := bufio.NewScanner(output)
    count := 0
    for scanner.Scan() {
        count++
    }

    if err := cmd.Wait(); err != nil {
        if exitErr, ok := err.(*exec.ExitError); ok {
            if exitErr.ExitCode() != 1 {
                ch <- Resultado{nombreArchivo, -1}
                return
            }
        } else {
            ch <- Resultado{nombreArchivo, -1}
            return
        }
    }

    ch <- Resultado{nombreArchivo, count}
}

func main() {
    if len(os.Args) < 3 {
        fmt.Println("Uso: go run main.go <letra> <nombreArchivo1> <nombreArchivo2> ...")
        return
    }

    letra := strings.ToLower(os.Args[1])
    if len(letra) != 1 {
        fmt.Println("Por favor, ingrese una sola letra como primer argumento.")
        return
    }

    ch := make(chan Resultado)
    var wg sync.WaitGroup
    total := 0

    for _, nombreArchivo := range os.Args[2:] {
        wg.Add(1)
        go func(archivo string) {
            defer wg.Done()
            contarLetra(archivo, letra, ch)
        }(nombreArchivo)
    }

    go func() {
        wg.Wait()
        close(ch)
    }()

    for resultado := range ch {
        if resultado.cuenta == -1 {
            fmt.Printf("Error al procesar el archivo: %s\n", resultado.nombreArchivo)
        } else {
            total += resultado.cuenta
            fmt.Printf("La letra '%s' aparece %d veces en el archivo %s.\n", letra, resultado.cuenta, resultado.nombreArchivo)
        }
    }

    fmt.Printf("Total de apariciones de la letra '%s': %d\n", letra, total)
}
