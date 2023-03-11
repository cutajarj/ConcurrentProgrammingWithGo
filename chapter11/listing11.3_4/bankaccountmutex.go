package listing11_3_4

import (
    "fmt"
    "sync"
)

type BankAccount struct {
    id      string
    balance int
    mutex   sync.Mutex
}

func NewBankAccount(id string) *BankAccount {
    return &BankAccount{
        id:      id,
        balance: 100,
        mutex:   sync.Mutex{},
    }
}

func (src *BankAccount) Transfer(to *BankAccount, amount int, exId int) {
    fmt.Printf("%d Locking %s's account\n", exId, src.id)
    src.mutex.Lock()
    fmt.Printf("%d Locking %s's account\n", exId, to.id)
    to.mutex.Lock()
    src.balance -= amount
    to.balance += amount
    to.mutex.Unlock()
    src.mutex.Unlock()
    fmt.Printf("%d Unlocked %s and %s\n", exId, src.id, to.id)
}

