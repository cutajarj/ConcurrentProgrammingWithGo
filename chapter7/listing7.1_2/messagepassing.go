package main

import "fmt"

func main() {
    msgChannel := make(chan string)
    go receiver(msgChannel)
    fmt.Println("Sending HELLO...")
    msgChannel <- "HELLO"
    fmt.Println("Sending THERE...")
    msgChannel <- "THERE"
    fmt.Println("Sending STOP...")
    msgChannel <- "STOP"
}

func receiver(messages chan string) {
    msg := ""
    for msg != "STOP" {
        msg = <-messages
        fmt.Println("Received:", msg)
    }
}
