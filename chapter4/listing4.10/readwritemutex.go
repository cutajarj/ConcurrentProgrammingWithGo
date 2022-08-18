package main

import (
    "sync"
    "time"
)

type ReadWriteMutex struct {
    readersCounter int
    readersLock    sync.Mutex
    globalLock     sync.Mutex
}

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

func main() {
    rwMutex := ReadWriteMutex{}
    for i := 0; i < 10; i++ {
        go func() {
            rwMutex.ReadLock()
            println("Read started")
            time.Sleep(5 * time.Second)
            println("Read done")
            rwMutex.ReadUnlock()
        }()
    }
    time.Sleep(1 * time.Second)
    println("Write started")
    rwMutex.WriteLock()
    println("Write finished")
}
