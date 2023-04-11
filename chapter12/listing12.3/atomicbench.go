package main

import (
    "os"
    "sync/atomic"
)

func main() {
    total := int64(0)
    if os.Args[1] == "atomic" {
        for i := int64(0); i < 1000000000; i++ {
            atomic.AddInt64(&total, i)
        }
    } else {
        for i := int64(0); i < 1000000000; i++ {
            total += i
        }
    }
}
