package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateTemp() chan int {
	output := make(chan int)
	go func() {
		temp := 50 //fahrenheit
		for {
			output <- temp
			temp += rand.Intn(3) - 1
			time.Sleep(200 * time.Millisecond)
		}
	}()
	return output
}

func outputTemp(input chan int) {
	go func() {
		for {
			fmt.Println("Current temp:", <-input)
			time.Sleep(2 * time.Second)
		}
	}()
}

func main() {
	temps := generateTemp()
	display := make(chan int)
	outputTemp(display)
	t := <-temps
	for {
		select {
		case t = <-temps:
		case display <- t:
		}
	}
}
