# GoBot

> For emulate move system

+ cursor 
+ keyboard
+ pause
+ terminal

***

> Requirements
``` bash
sudo apt install xautomation # install xte tool
```

> Install 
```
go get guthub.com/YuranIgnatenko/GoBot
```

> Use
``` go
package main

import . "github.com/YuranIgnatenko/GoBot"

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
```

***

> Constants special keys
```
Home
Left
Up
Right
Down
Page_Up
Page_Down
End
Return
BackSpace
Tab
Escape
Delete
Shift_L
Shift_R
Control_L
Control_R
Meta_L
Meta_R
Alt_L
Alt_R
Multi_key
Super_L
Super_R

```
***