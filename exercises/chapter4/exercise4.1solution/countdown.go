package main

import (
    "fmt"
    "sync"
    "time"
)

func countdown(seconds *int, mutex *sync.Mutex) {
    mutex.Lock()
    remaining := *seconds
    mutex.Unlock()
    for remaining > 0 {
        time.Sleep(1 * time.Second)
        mutex.Lock()
        *seconds -= 1
        remaining = *seconds
        mutex.Unlock()
    }
}

func main() {
    mutex := sync.Mutex{}
    count := 5
    go countdown(&count, &mutex)
    remaining := count
    for remaining > 0 {
        time.Sleep(500 * time.Millisecond)
        mutex.Lock()
        fmt.Println(count)
        remaining = count
        mutex.Unlock()
    }
}

