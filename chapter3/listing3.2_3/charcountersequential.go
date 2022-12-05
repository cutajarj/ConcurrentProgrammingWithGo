package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func countLetters(url string, frequency []int) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	for _, b := range body {
		c := strings.ToLower(string(b))
		cIndex := strings.Index(allLetters, c)
		if cIndex >= 0 {
			frequency[cIndex] += 1
		}
	}
	fmt.Println("Completed:", url)
}

func main() {
	var frequency = make([]int, 26)
	for i := 1000; i <= 1200; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		countLetters(url, frequency)
	}
	for i, c := range "abcdefghijklmnopqrstuvwxyz" {
		fmt.Printf("%c-%d ", c, frequency[i])
	}
}
