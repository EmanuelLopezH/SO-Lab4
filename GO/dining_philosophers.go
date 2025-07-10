package main

import (
    "fmt"
    "sync"
    "time"
)

const N = 5

var forks [N]sync.Mutex
var room = make(chan struct{}, N-1)

func philosopher(id int) {
    for {
        fmt.Println("Philosopher", id, "thinking...")
        time.Sleep(time.Second)

        room <- struct{}{}
        forks[id].Lock()
        forks[(id+1)%N].Lock()

        fmt.Println("Philosopher", id, "eating...")
        time.Sleep(time.Second)

        forks[id].Unlock()
        forks[(id+1)%N].Unlock()
        <-room
    }
}

func main() {
    for i := 0; i < N; i++ {
        go philosopher(i)
    }
    select {}
}
