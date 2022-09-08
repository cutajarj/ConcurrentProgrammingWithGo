package main

import (
    "os"
    "sync"
    "time"
)

func stingy(money *int, mutex *sync.Mutex) {
    for i := 0; i < 1000000; i++ {
        mutex.Lock()
        *money += 10
        mutex.Unlock()
    }
    println("Stingy Done")
}

func spendy(money *int, mutex *sync.Mutex) {
    for i := 0; i < 200000; i++ {
        mutex.Lock()
        for *money < 50 {
            mutex.Unlock()
            time.Sleep(10 * time.Millisecond)
            mutex.Lock()
        }
        *money -= 50
        if *money < 0 {
            println("Money is negative!")
            os.Exit(1)
        }
        mutex.Unlock()
    }
    println("Spendy Done")
}

func main() {
    money := 100
    mutex := sync.Mutex{}
    go stingy(&money, &mutex)
    go spendy(&money, &mutex)
    time.Sleep(2 * time.Second)
    mutex.Lock()
    println("Money in bank account: ", money)
    mutex.Unlock()
}
