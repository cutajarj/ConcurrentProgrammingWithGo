package main

import (
    "github.com/cutajarj/ConcurrentProgrammingWithGo/chapter4/listing4.10"
    "time"
)

func main() {
    rwMutex := listing4_10.ReadWriteMutex{}
    for i := 0; i < 2; i++ {
        go func() {
            for {
                rwMutex.ReadLock()
                time.Sleep(1 * time.Second)
                println("Read done")
                rwMutex.ReadUnlock()
            }
        }()
    }
    time.Sleep(1 * time.Second)
    rwMutex.WriteLock()
    println("Write finished")
}

