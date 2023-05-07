package main

import (
    "fmt"
    "io"
    "net/http"
    "strings"
    "sync"
    "sync/atomic"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func countLetters(url string, frequency []int32) {
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
            atomic.AddInt32(&frequency[cIndex], 1)
        }
    }
    fmt.Println("Completed:", url)
}


func main() {
    wg := sync.WaitGroup{}
    wg.Add(31)
    var frequency = make([]int32, 26)
    for i := 1000; i <= 1030; i++ {
        url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
        go func() {
            countLetters(url, frequency)
            wg.Done()
        }()
    }
    wg.Wait()
    for i, c := range allLetters {
        fmt.Printf("%c-%d ", c, atomic.LoadInt32(&frequency[i]))
    }
}

