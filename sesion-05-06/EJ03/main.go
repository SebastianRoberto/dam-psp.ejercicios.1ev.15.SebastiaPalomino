package main

import (
    "fmt"
    "strings"
    "sync"
)

func countWords(texts []string, wg *sync.WaitGroup, ch chan int) {
    defer wg.Done()
    count := 0
    for _, text := range texts {
        count += len(strings.Fields(text))
    }
    ch <- count
}

func main() {
    texts := []string{"Hola mundo", "Esto es un ejemplo", "Golang es divertido", "Las gorutinas son útiles"}
    mid := len(texts) / 2
    firstHalf := texts[:mid]
    secondHalf := texts[mid:]

    ch := make(chan int, 2)
    var wg sync.WaitGroup

    wg.Add(2)
    go countWords(firstHalf, &wg, ch)
    go countWords(secondHalf, &wg, ch)

    wg.Wait()
    close(ch)

    totalWords := 0
    for count := range ch {
        totalWords += count
    }

    fmt.Printf("El total de palabras es: %d\n", totalWords)
}
