package main

import (
    "fmt"
    "runtime"
)

func sayHello() {
    fmt.Println("Hello")
}

func main() {
    go sayHello()
    runtime.Gosched()
    fmt.Println("Finished")
}
