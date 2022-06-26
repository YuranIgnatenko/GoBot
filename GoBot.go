package go_bot

import (
	"os/exec"
	"strconv"
)

//need init struct for using methods
type Bot struct {
	/*
		Only linux-Ubuntu
		used package xautomation
		help:
		sudo apt-get install xautomation
	*/
	Home      string "Home"
	Left      string "Left"
	Up        string "Up"
	Right     string "Right"
	Down      string "Down"
	Page_Up   string "Page_Up"
	Page_Down string "Page_Down"
	End       string "End"
	Return    string "Return"
	BackSpace string "BackSpace"
	Tab       string "Tab"
	Escape    string "Escape"
	Delete    string "Delete"
	Shift_L   string "Shift_L"
	Shift_R   string "Shift_R"
	Control_L string "Control_L"
	Control_R string "Control_R"
	Meta_L    string "Meta_L"
	Meta_R    string "Meta_R"
	Alt_L     string "Alt_L"
	Alt_R     string "Alt_R"
	Multi_key string "Multi_key"
	Super_L   string "Super_L"
	Super_R   string "Super_R"
}

// inner method call func from terminal
func (gb *Bot) cmd(line string) error {
	_, err := exec.Command("xte", line).Output()
	return err
}

//set cursor position x y
func (gb *Bot) MouseMove(x, y int) error {
	return gb.cmd("mousemove" + " " + strconv.Itoa(x) + " " + strconv.Itoa(y))
}

// set cursor position and click
func (gb *Bot) MouseClick(x, y int) error {
	return gb.cmd("mouseclick" + " " + strconv.Itoa(x) + " " + strconv.Itoa(y))
}

//set curson position x y relative
func (gb *Bot) MouseRelativeMove(x, y int) error {
	return gb.cmd("mousermove" + " " + strconv.Itoa(x) + " " + strconv.Itoa(y))
}

//mouse key down :
//1 - left key
//2 - midle key
//3 - right key
func (gb *Bot) MouseDown(keyMouse int) error {
	return gb.cmd("mousedown" + " " + strconv.Itoa(keyMouse))
}

//mouse key up :
//1 - left key
//2 - midle key
//3 - right key
func (gb *Bot) MouseUp(keyMouse int) error {
	return gb.cmd("mouseup" + " " + strconv.Itoa(keyMouse))
}

//key press
//using special key:
//KeyPress(GoBot.Escape)
//KeyPress(GoBot.Home)
func (gb *Bot) KeyPress(key string) error {
	// ru key -- false (?)
	return gb.cmd("key" + " " + key)
}

//key down - long
//using special key:
//KeyPress(GoBot.Escape)
//KeyPress(GoBot.Home)
func (gb *Bot) KeyDown(key string) error {
	return gb.cmd("keydown" + " " + key)
}

//key up
//using special key:
//KeyPress(GoBot.Escape)
//KeyPress(GoBot.Home)
func (gb *Bot) KeyUp(key string) error {
	return gb.cmd("keyup" + " " + key)
}

//sleeping seconds
func (gb *Bot) SleepSec(sec int) error {
	return gb.cmd("sleep" + " " + strconv.Itoa(sec))
}

//sleeping milliseconds
func (gb *Bot) SleepMicro(sec int) error {
	return gb.cmd("usleep" + " " + strconv.Itoa(sec))
}

//write string line in field input
func (gb *Bot) WriteLine(line string) error {
	return gb.cmd("str" + " " + line)
}
