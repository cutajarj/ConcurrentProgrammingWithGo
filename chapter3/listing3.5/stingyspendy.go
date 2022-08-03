package main

import (
    "time"
)

func stingy(money *int) {
    for i := 0; i < 1000000; i++ {
        *money += 10
    }
    println("Stingy Done")
}

func spendy(money *int) {
    for i := 0; i < 1000000; i++ {
        *money -= 10
    }
    println("Spendy Done")
}

func main() {
    money := 100
    go stingy(&money)
    go spendy(&money)
    time.Sleep(2 * time.Second)
    print("Money in bank account: ", money)
}
