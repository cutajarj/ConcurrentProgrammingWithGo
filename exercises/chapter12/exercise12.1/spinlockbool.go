package exercise12_1

import (
    "runtime"
    "sync"
    "sync/atomic"
)

type SpinLock struct {
    isLocked atomic.Bool
}

func (s *SpinLock) Lock() {
    for !s.isLocked.CompareAndSwap(false, true) {
        runtime.Gosched()
    }
}

func (s *SpinLock) Unlock() {
    s.isLocked.Store(false)
}

func NewSpinLock() sync.Locker {
    var lock SpinLock
    return &lock
}

