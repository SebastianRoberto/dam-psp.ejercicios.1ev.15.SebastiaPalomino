package main

import (
    "fmt"
    "sync"
)

func sumArrayPart(array []int, wg *sync.WaitGroup, ch chan int) {
    defer wg.Done()
    sum := 0
    for _, num := range array {
        sum += num
    }
    ch <- sum
}

func main() {
    array := []int{1, 2, 3, 4, 5, 6}
    mid := len(array) / 2
    firstHalf := array[:mid]
    secondHalf := array[mid:]

    ch := make(chan int, 2)
    var wg sync.WaitGroup

    wg.Add(2)
    go sumArrayPart(firstHalf, &wg, ch)
    go sumArrayPart(secondHalf, &wg, ch)

    wg.Wait()
    close(ch)

    totalSum := 0
    for sum := range ch {
        totalSum += sum
    }

    fmt.Printf("La suma total es: %d\n", totalSum)
}
