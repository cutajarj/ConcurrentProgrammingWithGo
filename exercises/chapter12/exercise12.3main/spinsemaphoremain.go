package main

import (
    "fmt"
    "github.com/cutajarj/ConcurrentProgrammingWithGo/exercises/chapter12/exercise12.3"
    "sync"
    "time"
)

func acquireAndWait(id int, sema *exercise12_3.SpinSemaphore) {
    sema.Acquire()
    fmt.Println(id, "has acquired the semaphore")
    time.Sleep(2 * time.Second)
    fmt.Println(id, "releasing the semaphore")
    sema.Release()
}

func main() {
    sema := exercise12_3.NewSpinSemaphore(2)
    wg := sync.WaitGroup{}
    wg.Add(10)
    for i := 0; i < 10; i++ {
        go func(id int) {
            acquireAndWait(id, sema)
            wg.Done()
        }(i)
    }
    wg.Wait()
}
