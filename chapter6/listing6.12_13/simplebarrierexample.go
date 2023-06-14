package main

import (
    "fmt"
    "github.com/cutajarj/ConcurrentProgrammingWithGo/chapter6/listing6.10"
    "time"
)

func workAndWait(name string, timeToWork int, barrier *listing6_10.Barrier) {
    start := time.Now()
    for {
        fmt.Println(time.Since(start), name, "is running")
        time.Sleep(time.Duration(timeToWork) * time.Second)
        fmt.Println(time.Since(start), name, "is waiting on barrier")
        barrier.Wait()
    }
}

func main() {
    barrier := listing6_10.NewBarrier(2)
    go workAndWait("Red", 4, barrier)
    go workAndWait("Blue", 10, barrier)
    time.Sleep(100 * time.Second)
}
