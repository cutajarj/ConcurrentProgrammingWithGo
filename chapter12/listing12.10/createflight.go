package listing12_10

import (
    "github.com/cutajarj/ConcurrentProgrammingWithGo/chapter12/listing12.6"
    "github.com/cutajarj/ConcurrentProgrammingWithGo/chapter12/listing12.9"
)

func NewFlight(origin, dest string) *listing12_6.Flight {
    return &listing12_6.Flight{
        Origin:    origin,
        Dest:      dest,
        SeatsLeft: 200,
        Locker:    listing12_9.NewSpinLock(),
    }
}
