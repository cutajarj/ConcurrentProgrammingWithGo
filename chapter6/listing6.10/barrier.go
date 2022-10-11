package listing6_10

import "sync"

type Barrier struct {
    size      int
    waitCount int
    cond      *sync.Cond
}

func NewBarrier(size int) *Barrier {
    condVar := sync.NewCond(&sync.Mutex{})
    return &Barrier{size, 0, condVar}
}

func (b *Barrier) Wait() {
    b.cond.L.Lock()
    b.waitCount += 1
    if b.waitCount == b.size {
        b.waitCount = 0
        b.cond.Broadcast()
    } else {
        b.cond.Wait()
    }
    b.cond.L.Unlock()
}
