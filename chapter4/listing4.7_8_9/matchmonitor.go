package main

import (
    "container/list"
    "fmt"
    "strconv"
    "sync"
    "time"
)

func matchRecorder(matchEvents *list.List, mutex *sync.Mutex) {
    for i := 0; ; i++ {
        mutex.Lock()
        matchEvents.PushBack("Match event " + strconv.Itoa(i))
        mutex.Unlock()
        time.Sleep(200 * time.Millisecond)
        fmt.Println("Appended match event")
    }
}

func clientHandler(mEvents *list.List, mutex *sync.Mutex, st time.Time) {
    for i := 0; i < 100; i ++ {
        mutex.Lock()
        allEvents := copyAllEvents(mEvents)
        mutex.Unlock()
        timeTaken := time.Since(st)
        fmt.Println(len(allEvents), "events copied in", timeTaken)
    }
}

func copyAllEvents(matchEvents *list.List) []string {
    i := 0
    allEvents := make([]string, matchEvents.Len())
    for e := matchEvents.Front(); e != nil; e = e.Next() {
        allEvents[i] = e.Value.(string)
        i++
    }
    return allEvents
}

func main() {
    mutex := sync.Mutex{}
    var matchEvents = list.New()
    for j := 0; j < 10000; j++ {
        matchEvents.PushBack("Match event")
    }
    go matchRecorder(matchEvents, &mutex)
    start := time.Now()
    for j := 0; j < 5000; j++ {
        go clientHandler(matchEvents, &mutex, start)
    }
    time.Sleep(100 * time.Second)
}
