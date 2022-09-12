package main

import (
    "os"
    "sync"
    "time"
)

func main() {
    money := 100
    mutex := sync.Mutex{}
    cond := sync.NewCond(&mutex)
    go stingy(&money, cond)
    go spendy(&money, cond)
    time.Sleep(2 * time.Second)
    mutex.Lock()
    println("Money in bank account: ", money)
    mutex.Unlock()
}

func stingy(money *int, cond *sync.Cond) {
    for i := 0; i < 1000000; i++ {
        cond.L.Lock()
        *money += 10
        cond.Signal()
        cond.L.Unlock()
    }
    println("Stingy Done")
}

func spendy(money *int, cond *sync.Cond) {
    for i := 0; i < 200000; i++ {
        cond.L.Lock()
        for *money < 50 {
            cond.Wait()
        }
        *money -= 50
        if *money < 0 {
            println("Money is negative!")
            os.Exit(1)
        }
        cond.L.Unlock()
    }
    println("Spendy Done")
}

