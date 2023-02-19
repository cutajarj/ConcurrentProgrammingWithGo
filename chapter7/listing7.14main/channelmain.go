package main

import (
    "fmt"
    "github.com/cutajarj/ConcurrentProgrammingWithGo/chapter7/listing7.14"
    "sync"
    "time"
)

func receiver(messages *listing7_14.Channel[int], wGroup *sync.WaitGroup) {
    msg := 0
    for msg != -1 {
        time.Sleep(1 * time.Second)
        msg = messages.Receive()
        fmt.Println("Received:", msg)
    }
    wGroup.Done()
}

func main() {
    channel := listing7_14.NewChannel[int](10)
    wGroup := sync.WaitGroup{}
    wGroup.Add(1)
    go receiver(channel, &wGroup)
    for i := 1; i <= 6; i++ {
        fmt.Println("Sending: ", i)
        channel.Send(i)
    }
    channel.Send(-1)
    wGroup.Wait()
}
