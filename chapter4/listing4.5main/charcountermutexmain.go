package main

import (
    "fmt"
    "github.com/cutajarj/ConcurrentProgrammingWithGo/chapter4/listing4.5"
    "sync"
    "time"
)

func main() {
    mutex := sync.Mutex{}
    var frequency = make([]int, 26)
    for i := 1000; i <= 1030; i++ {
        url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
        go listing4_5.CountLetters(url, frequency, &mutex)
    }
    time.Sleep(10 * time.Second)
    mutex.Lock()
    for i, c := range listing4_5.AllLetters {
        fmt.Printf("%c-%d ", c, frequency[i])
    }
    mutex.Unlock()
}
