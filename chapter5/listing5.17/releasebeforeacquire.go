package main

import (
    "github.com/cutajarj/ConcurrentProgrammingWithGo/chapter5/listing5.16"
)

func main() {
    semaphore := listing5_16.NewSemaphore(0)
    for i := 0; i < 50000; i++ {
        go doWork(semaphore)
        println("Waiting for child goroutine")
        semaphore.Acquire()
        println("Child goroutine finished")
    }
}

func doWork(semaphore *listing5_16.Semaphore) {
    println("Work started")
    println("Work finished")
    semaphore.Release()
}

