package main

import (
    "fmt"
    "github.com/cutajarj/ConcurrentProgrammingWithGo/chapter10/listing10.13"
)

func AddOnPipe[X, Y any](f func(X) Y, q <-chan int, in <-chan X) chan Y {
    output := make(chan Y)
    go func() {
        defer close(output)
        for {
            select {
            case <-q:
                return
            case input := <-in:
                output <- f(input)
            }
        }
    }()
    return output
}

func main() {
    input := make(chan int)
    quit := make(chan int)
    output := AddOnPipe(listing10_13.Box, quit,
        AddOnPipe(listing10_13.AddToppings, quit,
            AddOnPipe(listing10_13.Bake, quit,
                AddOnPipe(listing10_13.Mixture, quit,
                    AddOnPipe(listing10_13.PrepareTray, quit, input)))))
    go func() {
        for i := 0; i < 10; i++ {
            input <- i
        }
    }()
    for i := 0; i < 10; i++ {
        fmt.Println(<-output, "received")
    }
}
