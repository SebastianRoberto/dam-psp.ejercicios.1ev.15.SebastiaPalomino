package main

import (
    "fmt"
    "sync"
)

func findMax(array []int, wg *sync.WaitGroup, ch chan int) {
    defer wg.Done()
    max := array[0]
    for _, num := range array {
        if num > max {
            max = num
        }
    }
    ch <- max
}

func main() {
    array := []int{3, 5, 9, 2, 8, 6}
    mid := len(array) / 2
    firstHalf := array[:mid]
    secondHalf := array[mid:]

    ch := make(chan int, 2)
    var wg sync.WaitGroup

    wg.Add(2)
    go findMax(firstHalf, &wg, ch)
    go findMax(secondHalf, &wg, ch)

    wg.Wait()
    close(ch)

    overallMax := <-ch
    for max := range ch {
        if max > overallMax {
            overallMax = max
        }
    }

    fmt.Printf("El máximo valor es: %d\n", overallMax)
}
