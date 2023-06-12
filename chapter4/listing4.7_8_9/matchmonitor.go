package main

import (
    "fmt"
    "strconv"
    "sync"
    "time"
)

func matchRecorder(matchEvents *[]string, mutex *sync.Mutex) {
    for i := 0; ; i++ {
        mutex.Lock()
        *matchEvents = append(*matchEvents,
            "Match event " + strconv.Itoa(i))
        mutex.Unlock()
        time.Sleep(200 * time.Millisecond)
        fmt.Println("Appended match event")
    }
}

func clientHandler(mEvents *[]string, mutex *sync.Mutex, st time.Time) {
    for i := 0; i < 100; i ++ {
        mutex.Lock()
        allEvents := copyAllEvents(mEvents)
        mutex.Unlock()

        timeTaken := time.Since(st)
        fmt.Println(len(allEvents), "events copied in", timeTaken)
    }
}

func copyAllEvents(matchEvents *[]string) []string {
    allEvents := make([]string, 0, len(*matchEvents))
    for _, e := range *matchEvents {
        allEvents = append(allEvents, e)
    }
    return allEvents
}

func main() {
    mutex := sync.Mutex{}
    var matchEvents = make([]string, 0, 10000)
    for j := 0; j < 10000; j++ {
        matchEvents = append(matchEvents, "Match event")
    }
    go matchRecorder(&matchEvents, &mutex)
    start := time.Now()
    for j := 0; j < 5000; j++ {
        go clientHandler(&matchEvents, &mutex, start)
    }
    time.Sleep(100 * time.Second)
}
