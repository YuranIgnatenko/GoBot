package main

import (
	. "GoBot"
	"fmt"
)

func main() {
	bot := Bot{}
	bot.MouseMove(10, 10)
	x := 0
	y := 200
	fmt.Println("1", x, y)
	// for x < 800 {
	// 	fmt.Println("2")

	// 	x++
	// 	bot.MouseMove(x, y)
	// 	er := bot.SleepSec(1)
	// 	fmt.Println(er)
	// }
	// for x > 0 {
	// 	x -= 2
	// 	bot.MouseRelativeMove(-2, 0)
	// }
}
