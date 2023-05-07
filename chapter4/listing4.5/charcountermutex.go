package listing4_5

import (
    "fmt"
    "io"
    "net/http"
    "strings"
    "sync"
    "time"
)

const AllLetters = "abcdefghijklmnopqrstuvwxyz"

func CountLetters(url string, frequency []int, mutex *sync.Mutex) {
    resp, _ := http.Get(url)
    defer resp.Body.Close()
    if resp.StatusCode != 200 {
        panic("Server returning error code: " + resp.Status)
    }
    body, _ := io.ReadAll(resp.Body)
    mutex.Lock()
    for _, b := range body {
        c := strings.ToLower(string(b))
        cIndex := strings.Index(AllLetters, c)
        if cIndex >= 0 {
            frequency[cIndex] += 1
        }
    }
    mutex.Unlock()
    fmt.Println("Completed:", url, time.Now().Format("15:04:05"))
}
