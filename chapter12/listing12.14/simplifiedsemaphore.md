# Simplifying Go's internal semaphore to understand how it works

Note: This code will not compile as since it is pseudocode. It is meant to illustrate how
Go's internal semaphore works without the extra noise of the runtime functionality

```go
package listing12_14

import "sync/atomic"

func semaphoreAcquire(permits *int32, queueAtTheBack bool) {
    for {
        v := atomic.LoadInt32(permits)
        if v != 0 && atomic.CompareAndSwapInt32(permits, v, v-1) {
            break
        }
        //The queue functions will only queue and park the
        //goroutine if the permits atomic variable is zero
        if queueAtTheBack {
            queueAndSuspendGoroutineAtTheEnd(permits)
        } else {
            queueAndSuspendGoroutineInFront(permits)
        }
    }
}
```
