package main

import (
    "math/rand"
    "sync"
    "time"
)

func doWork(id int, wg *sync.WaitGroup) {
    i := rand.Intn(5)
    time.Sleep(time.Duration(i) * time.Second)
    println(id, "Done working after", i, "seconds")
    wg.Done()
}

func main() {
    wg := sync.WaitGroup{}
    wg.Add(4)
    for i := 1; i <= 4; i++ {
        go doWork(i, &wg)
    }
    wg.Wait()
    println("All complete")
}
