package main

import (
    "github.com/cutajarj/ConcurrentProgrammingWithGo/chapter4/listing4.12"
    "time"
)

func main() {
    rwMutex := listing4_12.ReadWriteMutex{}
    for i := 0; i < 10; i++ {
        go func() {
            rwMutex.ReadLock()
            println("Read started")
            time.Sleep(5 * time.Second)
            println("Read done")
            rwMutex.ReadUnlock()
        }()
    }
    time.Sleep(1 * time.Second)
    println("Write started")
    rwMutex.WriteLock()
    println("Write finished")
}

