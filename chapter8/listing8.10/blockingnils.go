package main

import "fmt"

func main() {
    var ch chan string = nil
    ch <- "message"
    fmt.Println("This is never printed")
}
