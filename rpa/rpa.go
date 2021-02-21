package rpa

import (
	"github.com/go-vgo/robotgo"
)

type Mouse struct {
}

func (Mouse) LeftClick() {
	robotgo.MouseClick("left")
}

func (Mouse) RightClick() {
	robotgo.MouseClick("right")
}

type Keyboard struct {
}
