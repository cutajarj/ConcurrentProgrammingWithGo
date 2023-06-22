package main

import (
	"github.com/cutajarj/ConcurrentProgrammingWithGo/exercises/chapter9/exercise9.1"
	"github.com/cutajarj/ConcurrentProgrammingWithGo/exercises/chapter9/exercise9.2"
	"github.com/cutajarj/ConcurrentProgrammingWithGo/exercises/chapter9/exercise9.3"
	"github.com/cutajarj/ConcurrentProgrammingWithGo/exercises/chapter9/exercise9.4"
)

func main() {
	quit := make(chan int)
	exercise9_4.Drain(quit,
		exercise9_3.Print(quit,
			exercise9_2.TakeUntil(quit, func(s int) bool { return s <= 1000000 } ,
				exercise9_1.GenerateSquares(quit))))
	<-quit
}
