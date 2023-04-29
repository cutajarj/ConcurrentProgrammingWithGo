package main

import (
    "fmt"
    "github.com/cutajarj/ConcurrentProgrammingWithGo/exercises/chapter12/exercise12.2"
)

func main() {
    spin := exercise12_2.SpinLock{}

    spin.Lock()

    fmt.Println("This should be false: ", spin.TryLock())

    spin.Unlock()

    fmt.Println("This should be true: ", spin.TryLock())

    spin.Unlock()
}
