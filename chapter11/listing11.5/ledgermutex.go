package main

import (
    "fmt"
    "github.com/cutajarj/ConcurrentProgrammingWithGo/chapter11/listing11.3_4"
    "math/rand"
    "time"
)

func main() {
    accounts := []listing11_3_4.BankAccount{
        *listing11_3_4.NewBankAccount("Sam"),
        *listing11_3_4.NewBankAccount("Paul"),
        *listing11_3_4.NewBankAccount("Amy"),
        *listing11_3_4.NewBankAccount("Mia"),
    }
    total := len(accounts)
    for i := 0; i < 4; i++ {
        go func(eId int) {
            for j := 1; j < 1000; j++ {
                from, to := rand.Intn(total), rand.Intn(total)
                for from == to {
                    to = rand.Intn(total)
                }
                accounts[from].Transfer(&accounts[to], 10, eId)
            }
            fmt.Println(eId, "COMPLETE")
        }(i)
    }
    time.Sleep(60 * time.Second)
}
