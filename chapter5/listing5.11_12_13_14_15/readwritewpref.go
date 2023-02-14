package main

import (
    "fmt"
    "sync"
    "time"
)

type ReadWriteMutex struct {
    readersCounter int
    writersWaiting int
    writerActive   bool
    cond           *sync.Cond
}

func NewReadWriteMutex() *ReadWriteMutex {
    return &ReadWriteMutex{cond: sync.NewCond(&sync.Mutex{})}
}

func (rw *ReadWriteMutex) ReadLock() {
    rw.cond.L.Lock()
    for rw.writersWaiting > 0 || rw.writerActive {
        rw.cond.Wait()
    }
    rw.readersCounter++
    rw.cond.L.Unlock()
}

func (rw *ReadWriteMutex) WriteLock() {
    rw.cond.L.Lock()
    rw.writersWaiting++
    for rw.readersCounter > 0 || rw.writerActive {
        rw.cond.Wait()
    }
    rw.writersWaiting--
    rw.writerActive = true
    rw.cond.L.Unlock()
}

func (rw *ReadWriteMutex) ReadUnlock() {
    rw.cond.L.Lock()
    rw.readersCounter--
    if rw.readersCounter == 0 {
        rw.cond.Broadcast()
    }
    rw.cond.L.Unlock()
}

func (rw *ReadWriteMutex) WriteUnlock() {
    rw.cond.L.Lock()
    rw.writerActive = false
    rw.cond.Broadcast()
    rw.cond.L.Unlock()
}

func main() {
    rwMutex := NewReadWriteMutex()
    for i := 0; i < 2; i++ {
        go func() {
            for {
                rwMutex.ReadLock()
                time.Sleep(1 * time.Second)
                fmt.Println("Read done")
                rwMutex.ReadUnlock()
            }
        }()
    }
    time.Sleep(1 * time.Second)
    rwMutex.WriteLock()
    fmt.Println("Write finished")
}
