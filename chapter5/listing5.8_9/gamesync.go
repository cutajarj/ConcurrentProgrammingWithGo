package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    cond := sync.NewCond(&sync.Mutex{})
    playersInGame := 4
    for playerId := 0; playerId < 4; playerId++ {
        go playerHandler(cond, &playersInGame, playerId)
        time.Sleep(1 * time.Second)
    }
}

func playerHandler(cond *sync.Cond, playersRemaining *int, playerId int) {
    cond.L.Lock()
    fmt.Println(playerId, ": Connected")
    *playersRemaining--
    if *playersRemaining == 0 {
        cond.Broadcast()
    }
    for *playersRemaining > 0 {
        fmt.Println(playerId, ": Waiting for more players")
        cond.Wait()
    }
    cond.L.Unlock()
    fmt.Println("All players connected. Ready player", playerId)
    //Game started
}
