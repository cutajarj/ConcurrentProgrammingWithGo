package main

import (
	"fmt"
	"time"
)

func writeEvery(msg string, seconds int) <-chan string {
	messages := make(chan string)
	go func() {
		for {
			time.Sleep(time.Duration(seconds) * time.Second)
			messages <- msg
		}
	}()
	return messages
}

func main() {
	messagesFromA := writeEvery("Tick", 1)
	messagesFromB := writeEvery("Tock", 3)
	for {
		select {
		case msg1 := <-messagesFromA:
			fmt.Println(msg1)
		case msg2 := <-messagesFromB:
			fmt.Println(msg2)
		}
	}
}
