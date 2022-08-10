package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
    "sync"
    "time"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func countLetters(url string, frequency []int, mutex *sync.Mutex) {
    resp, _ := http.Get(url)
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    mutex.Lock()
    for _, b := range body {
        c := strings.ToLower(string(b))
        cIndex := strings.Index(allLetters, c)
        if cIndex >= 0 {
            frequency[cIndex] += 1
        }
    }
    mutex.Unlock()
    fmt.Println("Completed:", url)
}

func main() {
    mutex := sync.Mutex{}
    var frequency = make([]int, 26)
    for i := 1000; i <= 1200; i++ {
        url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
        go countLetters(fmt.Sprintf(url), frequency, &mutex)
    }
    time.Sleep(10 * time.Second)
    mutex.Lock()
    fmt.Println(frequency)
    mutex.Unlock()
}
