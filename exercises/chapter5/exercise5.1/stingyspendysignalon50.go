package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	money := 100
	mutex := sync.Mutex{}
	cond := sync.NewCond(&mutex)
	go stingy(&money, cond)
	go spendy(&money, cond)
	time.Sleep(2 * time.Second)
	mutex.Lock()
	fmt.Println("Money in bank account: ", money)
	mutex.Unlock()
}

func stingy(money *int, cond *sync.Cond) {
	for i := 0; i < 1000000; i++ {
		cond.L.Lock()
		*money += 10
		if *money >= 50 {
			cond.Signal()
		}
		cond.L.Unlock()
	}
	fmt.Println("Stingy Done")
}

func spendy(money *int, cond *sync.Cond) {
	for i := 0; i < 200000; i++ {
		cond.L.Lock()
		for *money < 50 {
			cond.Wait()
		}
		*money -= 50
		if *money < 0 {
			fmt.Println("Money is negative!")
			os.Exit(1)
		}
		cond.L.Unlock()
	}
	fmt.Println("Spendy Done")
}
