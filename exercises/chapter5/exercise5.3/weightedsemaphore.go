package main

import (
	"sync"
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
	for rw.permits <= 0 {
		rw.cond.Wait()
	}
	rw.permits--
	rw.cond.L.Unlock()
}

func (rw *WeightedSemaphore) Release(permits int) {
	rw.cond.L.Lock()
	rw.permits++
	rw.cond.Signal()
	rw.cond.L.Unlock()
}

func main() {

}
