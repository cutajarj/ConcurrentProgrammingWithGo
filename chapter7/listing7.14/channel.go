package listing7_14

import (
    "container/list"
    "github.com/cutajarj/ConcurrentProgrammingWithGo/chapter5/listing5.16"
    "sync"
)

type Channel[M any] struct {
    capacitySema *listing5_16.Semaphore
    sizeSema     *listing5_16.Semaphore
    mutex        sync.Mutex
    buffer       *list.List
}

func NewChannel[M any](capacity int) *Channel[M] {
    return &Channel[M]{
        capacitySema: listing5_16.NewSemaphore(capacity),
        sizeSema:     listing5_16.NewSemaphore(0),
        buffer:       list.New(),
    }
}

func (c *Channel[M]) Send(message M) {
    c.capacitySema.Acquire()
    c.mutex.Lock()
    c.buffer.PushBack(message)
    c.mutex.Unlock()
    c.sizeSema.Release()
}

func (c *Channel[M]) Receive() M {
    c.capacitySema.Release()
    c.sizeSema.Acquire()
    c.mutex.Lock()
    v := c.buffer.Remove(c.buffer.Front()).(M)
    c.mutex.Unlock()
    return v
}
