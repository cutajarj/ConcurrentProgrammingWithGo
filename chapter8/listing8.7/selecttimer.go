package main

import (
	"fmt"
	"os"
	"strconv"
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
	t, _ := strconv.Atoi(os.Args[1])
	messages := sendMsgAfter(3)
	timeoutDuration := time.Duration(t) * time.Second
	select {
	case msg := <-messages:
		fmt.Println("Message received:", msg)
	case tNow := <-time.After(timeoutDuration):
		fmt.Println("Timed out. Waited until:", tNow.Format("15:04:05"))
	}
}
