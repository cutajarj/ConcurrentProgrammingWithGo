package main

import (
    "fmt"
    "time"
)

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
    time.Sleep(5 * time.Second)
    fmt.Println("Receiver slept for 5 seconds")
}
