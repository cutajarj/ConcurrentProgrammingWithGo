package main

import (
    "sync"
)

func doWork(cond *sync.Cond) {
    println("Work started")
    println("Work finished")
    cond.L.Lock()
    cond.Signal()
    cond.L.Unlock()
}

func main() {
    cond := sync.NewCond(&sync.Mutex{})
    cond.L.Lock()
    for i := 0; i < 50000; i++ {
        go doWork(cond)
        println("Waiting for child goroutine")
        cond.Wait()
        println("Child goroutine finished")
    }
    cond.L.Unlock()
}