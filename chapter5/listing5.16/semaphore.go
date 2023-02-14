package listing5_16

import (
    "sync"
)

type Semaphore struct {
    permits int
    cond    *sync.Cond
}

func NewSemaphore(n int) *Semaphore {
    return &Semaphore{
        permits: n,
        cond:    sync.NewCond(&sync.Mutex{}),
    }
}

func (rw *Semaphore) Acquire() {
    rw.cond.L.Lock()
    for rw.permits <= 0 {
        rw.cond.Wait()
    }
    rw.permits--
    rw.cond.L.Unlock()
}

func (rw *Semaphore) Release() {
    rw.cond.L.Lock()
    rw.permits++
    rw.cond.Signal()
    rw.cond.L.Unlock()
}
