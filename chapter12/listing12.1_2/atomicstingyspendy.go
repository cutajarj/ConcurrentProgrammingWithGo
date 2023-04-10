package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

func stingy(money *int32) {
    for i := 0; i < 1000000; i++ {
        atomic.AddInt32(money, 10)
    }
    fmt.Println("Stingy Done")
}

func spendy(money *int32) {
    for i := 0; i < 1000000; i++ {
        atomic.AddInt32(money, -10)
    }
    fmt.Println("Spendy Done")
}

func main() {
    money := int32(100)
    wg := sync.WaitGroup{}
    wg.Add(2)
    go func() {
        stingy(&money)
        wg.Done()
    }()
    go func() {
        spendy(&money)
        wg.Done()
    }()
    wg.Wait()
    fmt.Println("Money in account: ", atomic.LoadInt32(&money))
}
