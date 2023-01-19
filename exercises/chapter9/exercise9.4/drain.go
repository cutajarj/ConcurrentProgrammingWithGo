package exercise9_4

func Drain[T any](quit <-chan int, input <-chan T) {
	go func() {
		for {
			select {
			case <-input:
			case <-quit:
				return
			}
		}
	}()
}
