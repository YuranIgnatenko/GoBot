package main

import (
	"os/exec"
	"strconv"
)

type GoBot struct {
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

func (gb *GoBot) cmd(line string) error {
	_, err := exec.Command("xte", line).Output()
	return err
}

func (gb *GoBot) MouseMove(x, y int) error {
	return gb.cmd("mousemove" + " " + strconv.Itoa(x) + " " + strconv.Itoa(y))
}

func (gb *GoBot) MouseClick(x, y int) error {
	return gb.cmd("mouseclick" + " " + strconv.Itoa(x) + " " + strconv.Itoa(y))
}

func (gb *GoBot) MouseRelativeMove(x, y int) error {
	return gb.cmd("mousermove" + " " + strconv.Itoa(x) + " " + strconv.Itoa(y))
}

func (gb *GoBot) MouseDown(keyMouse int) error {
	return gb.cmd("mousedown" + " " + strconv.Itoa(keyMouse))
}

func (gb *GoBot) MouseUp(keyMouse int) error {
	return gb.cmd("mouseup" + " " + strconv.Itoa(keyMouse))
}

func (gb *GoBot) KeyPress(key string) error {
	// ru key -- false (?)
	return gb.cmd("key" + " " + key)
}

func (gb *GoBot) KeyDown(key string) error {
	return gb.cmd("keydown" + " " + key)
}

func (gb *GoBot) KeyUp(key string) error {
	return gb.cmd("keyup" + " " + key)
}

func (gb *GoBot) SleepSec(sec int) error {
	return gb.cmd("sleep" + " " + strconv.Itoa(sec))
}

func (gb *GoBot) SleepMicro(sec int) error {
	return gb.cmd("usleep" + " " + strconv.Itoa(sec))
}

func (gb *GoBot) WriteLine(line string) error {
	return gb.cmd("str" + " " + line)
}
