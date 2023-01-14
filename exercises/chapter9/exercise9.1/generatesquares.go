package exercise9_1

func GenerateSquares(quit <-chan int) <-chan int {
	squares := make(chan int)
	go func() {
		i := 0
		defer close(squares)
		for {
			i += 1
			select {
			case squares <- i * i:
			case <-quit:
				return
			}
		}
	}()
	return squares
}
