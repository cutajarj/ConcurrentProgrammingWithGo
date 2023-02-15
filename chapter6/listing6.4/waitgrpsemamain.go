package main

import (
    "fmt"
    "github.com/cutajarj/ConcurrentProgrammingWithGo/chapter6/listing6.3"
)

func doWork(id int, wg *listing6_3.WaitGrp) {
    fmt.Println(id, "Done working ")
    wg.Done()
}

func main() {
    wg := listing6_3.NewWaitGrp(4)
    for i := 1; i <= 4; i++ {
        go doWork(i, wg)
    }
    wg.Wait()
    fmt.Println("All complete")
}
