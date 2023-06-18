package main

import (
    "fmt"
    "os"
    "strconv"
    "time"
)

func sendMsgAfter(seconds time.Duration) <-chan string {
    messages := make(chan string)
    go func() {
        time.Sleep(seconds)
        messages <- "Hello"
    }()
    return messages
}

func main() {
    t, _ := strconv.Atoi(os.Args[1])
    messages := sendMsgAfter(3 * time.Second)
    timeoutDuration := time.Duration(t) * time.Second
    fmt.Printf("Waiting for message for %d seconds...\n", t)
    select {
    case msg := <-messages:
        fmt.Println("Message received:", msg)
    case tNow := <-time.After(timeoutDuration):
        fmt.Println("Timed out. Waited until:", tNow.Format("15:04:05"))
    }
}
