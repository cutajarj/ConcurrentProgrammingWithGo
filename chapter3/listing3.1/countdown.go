package main

import (
    "time"
)

func countdown(seconds *int) {
    for *seconds > 0 {
        time.Sleep(1 * time.Second)
        *seconds -= 1
    }
}

func main() {
    count := 5
    countdown(&count)
    for count > 0 {
        time.Sleep(500 * time.Millisecond)
        println(count)
    }
}

