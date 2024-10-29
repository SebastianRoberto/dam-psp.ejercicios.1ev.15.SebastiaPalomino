package main

import (
    "fmt"
    "sync"
)

func countEvens(array []int, wg *sync.WaitGroup, ch chan int) {
    defer wg.Done()
    count := 0
    for _, num := range array {
        if num%2 == 0 {
            count++
        }
    }
    ch <- count
}

func main() {
    array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    mid := len(array) / 2
    firstHalf := array[:mid]
    secondHalf := array[mid:]

    ch := make(chan int, 2)
    var wg sync.WaitGroup

    wg.Add(2)
    go countEvens(firstHalf, &wg, ch)
    go countEvens(secondHalf, &wg, ch)

    wg.Wait()
    close(ch)

    totalEvens := 0
    for count := range ch {
        totalEvens += count
    }

    fmt.Printf("Total de números pares: %d\n", totalEvens)
}
