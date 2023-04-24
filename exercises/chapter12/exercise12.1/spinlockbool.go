package exercise12_1

import (
    "runtime"
    "sync"
    "sync/atomic"
)

type SpinLock atomic.Bool

func (s *SpinLock) Lock() {
    for !(*atomic.Bool)(s).CompareAndSwap(false, true) {
        runtime.Gosched()
    }
}

func (s *SpinLock) Unlock() {
    (*atomic.Bool)(s).Store(false)
}

func NewSpinLock() sync.Locker {
    var lock SpinLock
    return &lock
}

