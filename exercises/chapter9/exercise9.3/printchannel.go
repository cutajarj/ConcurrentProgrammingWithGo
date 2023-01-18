package exercise9_3

import "fmt"

func Print[T any](quit <-chan int, input <-chan T) <-chan T {
	output := make(chan T)
	go func() {
		defer close(output)
		moreData := true
		var msg T
		for moreData {
			select {
			case msg, moreData = <-input:
				if moreData {
					fmt.Println(msg)
					output <- msg
				}
			case <-quit:
				return
			}
		}
	}()
	return output
}
