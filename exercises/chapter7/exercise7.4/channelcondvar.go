package main

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

type Channel[M any] struct {
	cond        *sync.Cond
	maxCapacity int
	buffer      *list.List
}

func NewChannel[M any](capacity int) *Channel[M] {
	return &Channel[M]{
		cond:        sync.NewCond(&sync.Mutex{}),
		maxCapacity: capacity,
		buffer:      list.New(),
	}
}

func (c *Channel[M]) Send(message M) {
	c.cond.L.Lock()
	for c.buffer.Len() == c.maxCapacity {
		c.cond.Wait()
	}
	c.buffer.PushBack(message)
	c.cond.Broadcast()
	c.cond.L.Unlock()
}

func (c *Channel[M]) Receive() M {
	c.cond.L.Lock()
	c.maxCapacity++
	c.cond.Broadcast()
	for c.buffer.Len() == 0 {
		c.cond.Wait()
	}
	c.maxCapacity--
	v := c.buffer.Remove(c.buffer.Front()).(M)
	c.cond.L.Unlock()
	return v
}

func receiver(messages *Channel[int], wGroup *sync.WaitGroup) {
	msg := 0
	for msg != -1 {
		time.Sleep(1 * time.Second)
		msg = messages.Receive()
		fmt.Println("Received:", msg)
	}
	wGroup.Done()
}

func main() {
	channel := NewChannel[int](2)
	wGroup := sync.WaitGroup{}
	wGroup.Add(1)
	go receiver(channel, &wGroup)
	for i := 1; i <= 6; i++ {
		fmt.Println("Sending: ", i)
		channel.Send(i)
	}
	channel.Send(-1)
	wGroup.Wait()
}
