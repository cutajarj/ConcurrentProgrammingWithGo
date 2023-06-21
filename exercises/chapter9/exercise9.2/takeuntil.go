package exercise9_2

func TakeUntil[K any](quit chan int, f func(K) bool, input <-chan K) <-chan K {
    output := make(chan K)
    go func() {
        defer close(output)
        moreData := true
        fValue := true
        var msg K
        for fValue && moreData {
            select {
            case msg, moreData = <-input:
                if moreData {
                    fValue = f(msg)
                    if fValue {
                        output <- msg
                    }
                }
            case <-quit:
                return
            }
        }
        if !fValue {
            close(quit)
        }
    }()
    return output
}
