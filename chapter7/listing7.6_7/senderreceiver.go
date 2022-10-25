package main

import (
    "fmt"
    "sync"
)

func sender(messages chan<- string) {
    messages <- "HELLO"
    messages <- "THERE"
    messages <- "STOP"
}

func receiver(messages <-chan string, waitGroup *sync.WaitGroup) {
    msg := ""
    for msg != "STOP" {
        msg = <-messages
        fmt.Println("Received:", msg)
    }
    waitGroup.Done()
}

func main() {
    msgChannel := make(chan string)
    wGroup := sync.WaitGroup{}
    wGroup.Add(1)
    go sender(msgChannel)
    go receiver(msgChannel, &wGroup)
    wGroup.Wait()
}
