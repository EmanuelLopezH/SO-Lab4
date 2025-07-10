package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

const bufferSize = 5

var buffer = make([]int, 0, bufferSize)
var mutex = sync.Mutex{}
var empty = make(chan struct{}, bufferSize)
var full = make(chan struct{}, bufferSize)

func produce() int {
    return rand.Intn(100)
}

func consume(item int) {
    fmt.Println("Consumed:", item)
}

func producer() {
    for i := 0; i < 10; i++ {
        item := produce()
        empty <- struct{}{}
        mutex.Lock()
        buffer = append(buffer, item)
        fmt.Println("Produced:", item)
        mutex.Unlock()
        full <- struct{}{}
        time.Sleep(100 * time.Millisecond)
    }
}

func consumer() {
    for i := 0; i < 10; i++ {
        <-full
        mutex.Lock()
        item := buffer[0]
        buffer = buffer[1:]
        mutex.Unlock()
        <-empty
        consume(item)
        time.Sleep(150 * time.Millisecond)
    }
}

func main() {
    for i := 0; i < bufferSize; i++ {
        empty <- struct{}{}
    }
    go producer()
    go consumer()
    time.Sleep(3 * time.Second)
}
