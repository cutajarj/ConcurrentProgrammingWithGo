package listing9_18

func Take[K any](quit chan int, n int, input <-chan K) <-chan K {
    output := make(chan K)
    go func() {
        defer close(output)
        moreData := true
        var msg K
        for n > 0 && moreData {
            select {
            case msg, moreData = <-input:
                if moreData {
                    output <- msg
                    n--
                }
            case <-quit:
                return
            }
        }
        if n == 0 {
            close(quit)
        }
    }()
    return output
}
