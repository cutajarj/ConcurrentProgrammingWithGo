package listing4_4

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
    "sync"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func CountLetters(url string, frequency []int, mutex *sync.Mutex) {
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


