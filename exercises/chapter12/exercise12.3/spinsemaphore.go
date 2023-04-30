package exercise12_3

import (
    "sync/atomic"
)

type SpinSemaphore int32

func (s *SpinSemaphore) Acquire() {
    for {
        v := atomic.LoadInt32((*int32)(s))
        if v != 0 && atomic.CompareAndSwapInt32((*int32)(s), v, v - 1) {
            break
        }
    }
}

func (s *SpinSemaphore) Release() {
    atomic.AddInt32((*int32)(s), 1)
}

func NewSpinSemaphore(permits int32) *SpinSemaphore {
    var sema SpinSemaphore
    atomic.StoreInt32((*int32)(&sema), permits)
    return &sema
}
