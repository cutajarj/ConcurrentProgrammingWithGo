package listing9_14

func Broadcast[K any](quit <-chan int, input <-chan K, n int) []chan K {
    outputs := CreateAll[K](n)
    go func() {
        defer CloseAll(outputs...)
        var msg K
        moreData := true
        for moreData {
            select {
            case msg, moreData = <-input:
                if moreData {
                    for _, output := range outputs {
                        output <- msg
                    }
                }
            case <-quit:
                return
            }
        }
    }()
    return outputs
}

func CreateAll[K any](n int) []chan K {
    channels := make([]chan K, n)
    for i, _ := range channels {
        channels[i] = make(chan K)
    }
    return channels
}

func CloseAll[K any](channels ...chan K) {
    for _, output := range channels {
        close(output)
    }
}
