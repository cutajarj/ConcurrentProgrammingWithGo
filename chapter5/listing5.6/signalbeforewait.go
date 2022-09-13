package main

import (
    "sync"
)

/*
  Note: this program has a bug for demonstration purposes
  We demonstrate how to fix this problem in the next listing
*/
func doWork(cond *sync.Cond) {
    println("Work started")
    println("Work finished")
    cond.Signal()
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

