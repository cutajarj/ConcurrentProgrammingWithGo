package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const matrixSize = 3

func generateRandMatrix(matrix *[matrixSize][matrixSize]int) {
	for row := 0; row < matrixSize; row++ {
		for col := 0; col < matrixSize; col++ {
			matrix[row][col] = rand.Intn(10) - 5
		}
	}
}

func rowMultiply(matrixA, matrixB, result *[matrixSize][matrixSize]int, row int, wg *sync.WaitGroup) {
	for col := 0; col < matrixSize; col++ {
		sum := 0
		for i := 0; i < matrixSize; i++ {
			sum += matrixA[row][i] * matrixB[i][col]
		}
		result[row][col] = sum
	}
	wg.Done()
}

func main() {
	var matrixA, matrixB, result [matrixSize][matrixSize]int
	for i := 0; i < 4; i++ {
		generateRandMatrix(&matrixA)
		generateRandMatrix(&matrixB)
		wg := sync.WaitGroup{}
		wg.Add(matrixSize)
		for row := 0; row < matrixSize; row++ {
			go rowMultiply(&matrixA, &matrixB, &result, row, &wg)
		}
		wg.Wait()
		for i := 0; i < matrixSize; i++ {
			fmt.Println(matrixA[i], matrixB[i], result[i])
		}
		fmt.Println()
	}
}
