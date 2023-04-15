package main

import (
    "fmt"
    "sync/atomic"
)

func main() {
    number := int32(17)
    result := atomic.CompareAndSwapInt32(&number, 17, 19)
    fmt.Printf("17 <- swap(17,19): result %t, value: %d\n", result, number)
    number = int32(23)
    result = atomic.CompareAndSwapInt32(&number, 17, 19)
    fmt.Printf("23 <- swap(17,19): result %t, value: %d\n", result, number)
}
