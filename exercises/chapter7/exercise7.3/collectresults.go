package main

import (
	"fmt"
	"math/rand"
)

func findFactors(number int) []int {
	result := make([]int, 0)
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			result = append(result, i)
		}
	}
	return result
}

func main() {
	resultChs := make([]chan []int, 10)
	for i := 0; i < 10; i++ {
		resultChs[i] = make(chan []int)
		go func(n int) {
			resultChs[n] <- findFactors(rand.Intn(1000000000))
		}(i)
	}
	for i := 0; i < 10; i++ {
		fmt.Println(<-resultChs[i])
	}
}
