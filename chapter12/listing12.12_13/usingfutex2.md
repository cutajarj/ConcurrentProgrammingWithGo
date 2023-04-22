# Developing a mutex using a futex attempt 2

Note: This code will not compile because we cannot access futexes from Go.
This code is to demonstrate the mutex algorithms. It is our second attempt to
implement a mutex using futexes.

```go
package listing12_12

import "sync/atomic"

type FutexLock int32

func (f *FutexLock) Unlock() {
    oldValue := atomic.SwapInt32((*int32)(f), 0)
    if oldValue == 2 {
        futex_wakeup((*int32)(f), 1)
    }
}

func (f *FutexLock) Lock() {
    if !atomic.CompareAndSwapInt32((*int32)(f), 0, 1) {
        for atomic.SwapInt32((*int32)(f), 2) != 0 {
            futex_wait((*int32)(f), 2)
        }
    }
}
```
