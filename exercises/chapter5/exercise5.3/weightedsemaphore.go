package main

import (
	"fmt"
	"sync"
	"time"
)

type WeightedSemaphore struct {
	permits int
	cond    *sync.Cond
}

func NewSemaphore(n int) *WeightedSemaphore {
	return &WeightedSemaphore{
		permits: n,
		cond:    sync.NewCond(&sync.Mutex{}),
	}
}

func (rw *WeightedSemaphore) Acquire(permits int) {
	rw.cond.L.Lock()
	for rw.permits-permits < 0 {
		rw.cond.Wait()
	}
	rw.permits -= permits
	rw.cond.L.Unlock()
}

func (rw *WeightedSemaphore) Release(permits int) {
	rw.cond.L.Lock()
	rw.permits += permits
	rw.cond.Broadcast()
	rw.cond.L.Unlock()
}

func main() {
	sema := NewSemaphore(2)
	sema.Acquire(2)
	fmt.Println("Parent thread A acquired semaphore")
	go func() {
		sema.Acquire(1)
		fmt.Println("Child B thread acquired semaphore")
		time.Sleep(1 * time.Second)
		sema.Release(1)
		fmt.Println("Child B thread released semaphore")
	}()
	go func() {
		sema.Acquire(1)
		fmt.Println("Child C thread acquired semaphore")
		time.Sleep(1 * time.Second)
		sema.Release(1)
		fmt.Println("Child C thread released semaphore")
	}()
	time.Sleep(3 * time.Second)
	fmt.Println("Parent thread A releasing semaphore")
	sema.Release(2)
	time.Sleep(3 * time.Second)
}
