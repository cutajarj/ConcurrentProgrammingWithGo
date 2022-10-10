package main

import (
    "github.com/cutajarj/ConcurrentProgrammingWithGo/chapter6/listing6.7"
)

func doWork(id int, wg *listing6_7.WaitGrp) {
    println(id, "Done working ")
    wg.Done()
}

func main() {
    wg := listing6_7.NewWaitGrp()
    for i := 1; i <= 4; i++ {
        wg.Add(2)
        go doWork(i, wg)
        go doWork(i, wg)
    }
    wg.Wait()
    println("All complete")
}
