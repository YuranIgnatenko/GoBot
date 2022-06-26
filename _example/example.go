package main

import . "go_robot"

func main() {
	bot := Bot{}
	bot.MouseMove(10, 10)
	x := 0
	y := 200
	for x < 800 {
		x++
		bot.MouseMove(x, y)
	}
	for x > 0 {
		x -= 2
		bot.MouseRelativeMove(-2, 0)
	}
}
