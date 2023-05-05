package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
)

/*
Note: this program has a race condition for demonstration purposes
Additionally we have a timer at the end which you might need to adjust
depending on how fast your internet connection is.
In later chapters we cover how to wait for threads to complete their work
*/
func countLetters(url string, frequency map[string]int) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic("Server's error: " + resp.Status)
	}
	body, _ := io.ReadAll(resp.Body)
	wordRegex := regexp.MustCompile(`[a-zA-Z]+`)
	for _, word := range wordRegex.FindAllString(string(body), -1) {
		wordLower := strings.ToLower(word)
		frequency[wordLower] += 1
	}
	fmt.Println("Completed:", url)
}

func main() {
	var frequency = make(map[string]int)
	for i := 1000; i <= 1020; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		go countLetters(url, frequency)
	}
	time.Sleep(10 * time.Second)
	for k, v := range frequency {
		fmt.Println(k, "->", v)
	}
}
