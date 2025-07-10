package main

import (
    "fmt"
    "sync"
    "time"
)

const SIZE = 10

type ThreadSafeQueue struct {
    items    []int
    lock     sync.Mutex
    notEmpty *sync.Cond
    notFull  *sync.Cond
}

func NewQueue() *ThreadSafeQueue {
    q := &ThreadSafeQueue{items: make([]int, 0, SIZE)}
    q.notEmpty = sync.NewCond(&q.lock)
    q.notFull = sync.NewCond(&q.lock)
    return q
}

func (q *ThreadSafeQueue) Enqueue(item int) {
    q.lock.Lock()
    for len(q.items) == SIZE {
        q.notFull.Wait()
    }
    q.items = append(q.items, item)
    q.notEmpty.Signal()
    q.lock.Unlock()
}

func (q *ThreadSafeQueue) Dequeue() int {
    q.lock.Lock()
    for len(q.items) == 0 {
        q.notEmpty.Wait()
    }
    item := q.items[0]
    q.items = q.items[1:]
    q.notFull.Signal()
    q.lock.Unlock()
    return item
}

func main() {
    q := NewQueue()

    go func() {
        for i := 0; i < 20; i++ {
            fmt.Println("Producer enqueuing", i)
            q.Enqueue(i)
            time.Sleep(100 * time.Millisecond)
        }
    }()

    go func() {
        for i := 0; i < 20; i++ {
            val := q.Dequeue()
            fmt.Println("Consumer dequeued", val)
            time.Sleep(150 * time.Millisecond)
        }
    }()

    time.Sleep(5 * time.Second)
}
