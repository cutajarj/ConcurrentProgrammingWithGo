package main

import (
	"fmt"
	"sync"
	"time"
)

type WaitGrp struct {
	groupSize int
	cond      *sync.Cond
}

func NewWaitGrp() *WaitGrp {
	return &WaitGrp{
		cond: sync.NewCond(&sync.Mutex{}),
	}
}

func (wg *WaitGrp) Add(delta int) {
	wg.cond.L.Lock()
	wg.groupSize += delta
	wg.cond.L.Unlock()
}

func (wg *WaitGrp) TryWait() bool {
	wg.cond.L.Lock()
	result := wg.groupSize == 0
	wg.cond.L.Unlock()
	return result
}

func (wg *WaitGrp) Wait() {
	wg.cond.L.Lock()
	for wg.groupSize > 0 {
		wg.cond.Wait()
	}
	wg.cond.L.Unlock()
}

func (wg *WaitGrp) Done() {
	wg.cond.L.Lock()
	wg.groupSize--
	if wg.groupSize == 0 {
		wg.cond.Broadcast()
	}
	wg.cond.L.Unlock()
}

func main() {
	wg := NewWaitGrp()
	wg.Add(1)
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Marking Wait group as done.")
		wg.Done()
	}()
	for !wg.TryWait() {
		fmt.Println("Wait group is not done. Trying again later.")
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Wait group is done.")
}
