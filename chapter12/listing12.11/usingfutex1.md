# Developing a mutex using a futex attempt 1

Note: This code will not compile because we cannot access futexes from Go. This code is to demonstrate the mutex algorithms.

```go
package listing12_11

import "sync/atomic"

type FutexLock int32

func (f *FutexLock) Lock() {
    for !atomic.CompareAndSwapInt32((*int32)(f), 0, 1) {
        futex_wait((*int32)(f), 1)
    }
}

func (f *FutexLock) Unlock() {
    atomic.StoreInt32((*int32)(f), 0)
    futex_wakeup((*int32)(f), 1)
}
```
