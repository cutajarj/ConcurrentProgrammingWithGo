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
    for i := 2000; i <= 2200; i++ {
        url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
        go listing4_5.CountLetters(url, frequency, &mutex)
    }
    for i := 0; i < 100; i++ {
        time.Sleep(100 * time.Millisecond)
        if mutex.TryLock() {
            for i, c := range listing4_5.AllLetters {
                fmt.Printf("%c-%d ", c, frequency[i])
            }
            fmt.Println()
            mutex.Unlock()
        } else {
            fmt.Println("Mutex already being used")
        }
    }
}
