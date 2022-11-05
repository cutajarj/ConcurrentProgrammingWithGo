package main

import (
    "sync"
    "time"
)

func playerHandler(cond *sync.Cond, playersRemaining *int,
    playerId int, cancel *bool) {
    cond.L.Lock()
    println(playerId, ": Connected")
    *playersRemaining--
    if *playersRemaining == 0 {
        cond.Broadcast()
    }
    for *playersRemaining > 0 && !*cancel {
        println(playerId, ": Waiting for more players")
        cond.Wait()
    }
    cond.L.Unlock()
    if *cancel {
        println(playerId, ": Game cancelled")
    } else {
        println("All players connected. Ready player", playerId)
    }
}

func timeout(cond *sync.Cond, cancel *bool) {
    time.Sleep(10 * time.Second)
    cond.L.Lock()
    *cancel = true
    cond.Broadcast()
    cond.L.Unlock()
}

func main() {
    cond := sync.NewCond(&sync.Mutex{})
    cancel := false
    go timeout(cond, &cancel)
    playersInGame := 5
    for i := 0; i < 4; i++ {
        go playerHandler(cond, &playersInGame, i, &cancel)
        time.Sleep(1 * time.Second)
    }
    time.Sleep(60 * time.Second)
}
