package main

import (
    "fmt"
    "github.com/cutajarj/ConcurrentProgrammingWithGo/chapter10/listing10.13"
)

func main() {
    for i := 0; i < 10; i++ {
        result := listing10_13.Box(
            listing10_13.AddToppings(
                listing10_13.Bake(
                    listing10_13.Mixture(
                        listing10_13.PrepareTray(i)))))
        fmt.Println("Accepting", result)
    }
}
