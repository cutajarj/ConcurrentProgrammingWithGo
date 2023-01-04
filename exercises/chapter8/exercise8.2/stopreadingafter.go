package main

import (
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
			println(n)
		case <-timeout:
			println("Stopping reading after timeout")
			return
		}
	}
}
