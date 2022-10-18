package listing4_12

import (
    "sync"
)

// Listing 4.12
type ReadWriteMutex struct {
    readersCounter int
    readersLock    sync.Mutex
    globalLock     sync.Mutex
}

// Listing 4.13
func (rw *ReadWriteMutex) ReadLock() {
    rw.readersLock.Lock()
    rw.readersCounter++
    if rw.readersCounter == 1 {
        rw.globalLock.Lock()
    }
    rw.readersLock.Unlock()
}

func (rw *ReadWriteMutex) WriteLock() {
    rw.globalLock.Lock()
}

// Listing 4.14
func (rw *ReadWriteMutex) ReadUnlock() {
    rw.readersLock.Lock()
    rw.readersCounter--
    if rw.readersCounter == 0 {
        rw.globalLock.Unlock()
    }
    rw.readersLock.Unlock()
}

func (rw *ReadWriteMutex) WriteUnlock() {
    rw.globalLock.Unlock()
}

