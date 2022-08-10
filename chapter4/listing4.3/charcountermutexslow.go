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

/*
  Note: this program us locking the entire goroutine with mutex on purpose to demonstrate
  bad placement of the lock and unlock. We fix this in the next listing
*/
func countLetters(url string, frequency []int, mutex *sync.Mutex) {
    mutex.Lock()
    resp, _ := http.Get(url)
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    for _, b := range body {
        c := strings.ToLower(string(b))
        cIndex := strings.Index(allLetters, c)
        if cIndex >= 0 {
            frequency[cIndex] += 1
        }
    }
    fmt.Println("Completed:", url)
    mutex.Unlock()
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
