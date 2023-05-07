package main

import (
    "fmt"
    "io"
    "net/http"
    "strings"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func countLetters(url string) <-chan []int {
    result := make(chan []int)
    go func() {
        defer close(result)
        frequency := make([]int, 26)
        resp, _ := http.Get(url)
        defer resp.Body.Close()
        if resp.StatusCode != 200 {
            panic("Server returning error code: " + resp.Status)
        }
        body, _ := io.ReadAll(resp.Body)
        for _, b := range body {
            c := strings.ToLower(string(b))
            cIndex := strings.Index(allLetters, c)
            if cIndex >= 0 {
                frequency[cIndex] += 1
            }
        }
        fmt.Println("Completed:", url)
        result <- frequency
    }()
    return result
}

func main() {
    results := make([]<-chan []int, 0)
    totalFrequencies := make([]int, 26)
    for i := 1000; i <= 1030; i++ {
        url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
        results = append(results, countLetters(url))
    }
    for _, c := range results {
        frequencyResult := <-c
        for i := 0; i < 26; i++ {
            totalFrequencies[i] += frequencyResult[i]
        }
    }
    for i, c := range allLetters {
        fmt.Printf("%c-%d ", c, totalFrequencies[i])
    }
}
