package main

import (
    "fmt"
    "sync"
)

func multiplyByFactor(num int, factor int, wg *sync.WaitGroup, ch chan int) {
    defer wg.Done()
    ch <- num * factor
}

func main() {
    array := []int{2, 4, 6, 8}
    factor := 3

    ch := make(chan int, len(array))
    var wg sync.WaitGroup

    for _, num := range array {
        wg.Add(1)
        go multiplyByFactor(num, factor, &wg, ch)
    }

    wg.Wait()
    close(ch)

    for result := range ch {
        fmt.Println("Resultado:", result)
    }
}
