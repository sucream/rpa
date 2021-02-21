package main

import (
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

const (
	HookEnabled  = 1 // iota
	HookDisabled = 2

	KeyDown = 3
	KeyHold = 4
	KeyUp   = 5

	MouseUp    = 6
	MouseHold  = 7
	MouseDown  = 8
	MouseMove  = 9
	MouseDrag  = 10
	MouseWheel = 11

	FakeEvent = 12

	// Keychar could be v
	CharUndefined = 0xFFFF
	WheelUp       = -1
	WheelDown     = 1
)

func main() {
	logging := []hook.Event{}

	if ok := robotgo.ShowAlert("Start Recording...", "Start all your action(keydown, mousemove, mousedown) "); ok {
		evChan := robotgo.EventStart()
		for e := range evChan {
			if e.Keychar == '?' {
				robotgo.EventEnd()
				robotgo.ShowAlert("Stop Recording...", "Stop Recording.")
				break
			}
			if e.Kind == KeyDown || e.Kind == MouseMove || e.Kind == MouseDown {
				logging = append(logging, e)
			}
		}
		for _, e := range logging {
			if e.Kind == KeyDown {
				if e.Keycode == 13 {
					robotgo.KeyTap("enter")
				} else if e.Keycode == 8 {
					robotgo.KeyTap("backspace")
				} else {
					robotgo.KeyTap(string(e.Keychar))
				}
			} else if e.Kind == MouseMove {
				robotgo.MoveMouse(int(e.X), int(e.Y))
			} else if e.Kind == MouseDown {

			}
		}
	} else {
		robotgo.ShowAlert("Exit program", "bye")
	}

}
