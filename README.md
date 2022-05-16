# gobot 

Step 1: (in terminal/shell/bash)
sudo apt-get install xautomation

Step 2: (in terminal/shell/bash)
go install github.com/YuranIgnatenko/gobot

Step 3: (in main file)
// init structure
bot := GoBot{}

// use functions structure
bot.MouseMove(100, 10)
bot.MouseClick(100, 10)
bot.MouseRelativeMove(10, 10)
bot.MouseDown(3)
bot.KeyPress("K")
bot.WriteLine("Hello world")

Thank you !
