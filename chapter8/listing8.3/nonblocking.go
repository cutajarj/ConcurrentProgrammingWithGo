package main

import (
    "fmt"
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
    messages := sendMsgAfter(3 * time.Second)
    for {
        select {
        case msg := <-messages:
            fmt.Println("Message received:", msg)
            return
        default:
            fmt.Println("No messages waiting")
            time.Sleep(1 * time.Second)
        }
    }
}
