package listing12_6

import (
    "sync"
)

type Flight struct {
    Origin, Dest string
    SeatsLeft int
    Locker    sync.Locker
}

