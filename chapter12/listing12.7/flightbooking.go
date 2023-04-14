package listing12_7

import (
    "github.com/cutajarj/ConcurrentProgrammingWithGo/chapter12/listing12.6"
    "sort"
)

func Book(flights []*listing12_6.Flight, seatsToBook int) bool {
    bookable := true
    sort.Slice(flights, func(a, b int) bool {
        flightA := flights[a].Origin + flights[a].Dest
        flightB := flights[b].Origin + flights[b].Dest
        return flightA < (flightB)
    })
    for _, f := range flights {
        f.Locker.Lock()
    }
    for i := 0; i < len(flights) && bookable; i++ {
        if flights[i].SeatsLeft < seatsToBook {
            bookable = false
        }
    }
    for i := 0; i < len(flights) && bookable; i++ {
        flights[i].SeatsLeft-=seatsToBook
    }
    for _, f := range flights {
        f.Locker.Unlock()
    }
    return bookable
}
