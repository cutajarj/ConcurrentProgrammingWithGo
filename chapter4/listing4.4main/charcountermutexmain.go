package main

import (
    "fmt"
    "github.com/cutajarj/ConcurrentProgrammingWithGo/chapter4/listing4.4"
    "sync"
    "time"
)

func main() {
    mutex := sync.Mutex{}
    var frequency = make([]int, 26)
    for i := 1000; i <= 1200; i++ {
        url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
        go listing4_4.CountLetters(url, frequency, &mutex)
    }
    time.Sleep(10 * time.Second)
    mutex.Lock()
    fmt.Println(frequency)
    mutex.Unlock()
}
