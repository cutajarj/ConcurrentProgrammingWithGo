package main

import (
    "sync"
    "time"
)

func main() {
    mutex := sync.Mutex{}
    mutex.Lock()
    go pollMutex(&mutex)
    time.Sleep(1 * time.Second)
    mutex.Unlock()
    time.Sleep(1 * time.Second)
}

func pollMutex(mutex *sync.Mutex) {
    for !mutex.TryLock() {
        println("Mutex already being used")
        time.Sleep(200 * time.Millisecond)
    }
    println("Child goroutine acquired mutex")
    mutex.Unlock()
}
