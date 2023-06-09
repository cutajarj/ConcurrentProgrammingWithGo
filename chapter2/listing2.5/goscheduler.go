package main

import (
    "fmt"
    "runtime"
)
/*
Note: this program is an example of what not to do; using go scheduler
to synchronize executions
*/
func sayHello() {
    fmt.Println("Hello")
}

func main() {
    go sayHello()
    runtime.Gosched()
    fmt.Println("Finished")
}
