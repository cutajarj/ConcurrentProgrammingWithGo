package main

import "fmt"

func primeMultipleFilter(numbers <-chan int, quit chan<- int) {
	var right chan int
	p := <-numbers
	fmt.Println(p)
	for n := range numbers {
		if n%p != 0 {
			if right == nil {
				right = make(chan int)
				go primeMultipleFilter(right, quit)
			}
			right <- n
		}
	}
	if right == nil {
		close(quit)
	} else {
		close(right)
	}
}

func main() {
	numbers := make(chan int)
	quit := make(chan int)
	go primeMultipleFilter(numbers, quit)
	for i := 2; i < 100000; i++ {
		numbers <- i
	}
	close(numbers)
	<-quit
}
