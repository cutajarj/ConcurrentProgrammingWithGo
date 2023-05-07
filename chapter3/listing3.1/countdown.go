package main

import (
    "fmt"
    "time"
)
/*
Note: this program has a race condition for demonstration purposes
This is then fixed later as an exercise in chapter 4
 */
func countdown(seconds *int) {
    for *seconds > 0 {
        time.Sleep(1 * time.Second)
        *seconds -= 1
    }
}

func main() {
    count := 5
    go countdown(&count)
    for count > 0 {
        time.Sleep(500 * time.Millisecond)
        fmt.Println(count)
    }
}
