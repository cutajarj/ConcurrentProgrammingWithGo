package main

import (
    "fmt"
    "time"
)

func doWork(id int) {
    fmt.Println(id, "Work started at", time.Now().Format("15:04:05"))
    time.Sleep(1 * time.Second)
    fmt.Println(id, "Work finished at", time.Now().Format("15:04:05"))
}

func main() {
    for i := 0; i < 5; i++ {
        go doWork(i)
    }
    time.Sleep(2 * time.Second)
}
