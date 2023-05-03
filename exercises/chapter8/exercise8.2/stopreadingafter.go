package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateNumbers() chan int {
	output := make(chan int)
	go func() {
		for {
			output <- rand.Intn(10)
			time.Sleep(200 * time.Millisecond)
		}
	}()
	return output
}

func main() {
	numbers := generateNumbers()
	timeout := time.After(5 * time.Second)
	for {
		select {
		case n := <-numbers:
			fmt.Println(n)
		case <-timeout:
			fmt.Println("Stopping reading after timeout")
			return
		}
	}
}
