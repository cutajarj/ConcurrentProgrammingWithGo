package main

import (
    "fmt"
    "time"
)

func sendMsgAfter(seconds int) <-chan string {
    messages := make(chan string)
    go func() {
        time.Sleep(time.Duration(seconds) * time.Second)
        messages <- "Hello"
    }()
    return messages
}

func main() {
    messages := sendMsgAfter(3)
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
