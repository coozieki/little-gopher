package state

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Data struct {
	CurrentState State
	States       struct {
		Menu     State
		Play     State
		Pause    State
		Options  State
		Gameover State
	}
	Actions struct {
		Gameover func()
	}
}

type State interface {
	Update(data *Data)
	Draw(screen *ebiten.Image, data *Data)
}
