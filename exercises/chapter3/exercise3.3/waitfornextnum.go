package main

import (
	"fmt"
	"time"
)

/*
  Note: this program has a race condition for demonstration purposes
  You might need to run this multiple times to trigger the race condition
*/

func addNextNumber(nextNum *[101]int) {
	i := 0
	for nextNum[i] != 0 {
		i++
	}
	nextNum[i] = nextNum[i-1] + 1
}

func main() {
	nextNum := [101]int{1}
	for i := 0; i < 100; i++ {
		go addNextNumber(&nextNum)
	}
	for nextNum[100] == 0 {
		fmt.Println("Waiting for goroutines to complete")
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Println(nextNum)
}
