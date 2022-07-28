package main

import (
    "fmt"
    "time"
)

func doWork(id int) {
    time.Sleep(1 * time.Second)
    fmt.Println(id, "Work finished")
}

func main() {
    for i := 0; i < 5; i++ {
        doWork(i)
    }
}
