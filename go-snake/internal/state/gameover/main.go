package gameover

import (
	"github.com/hajimehoshi/ebiten/v2"
	"go-snake/internal/config"
	"go-snake/internal/events"
	"go-snake/internal/state"
	"image/color"
)

type gameoverState struct {
	events.EventListener
	stateData *state.Data
}

var Gameover *gameoverState

func init() {
	Gameover = &gameoverState{EventListener: events.NewEventListener()}

	Gameover.ButtonPress(&events.ButtonPressHandler{
		Key: ebiten.KeyEscape,
		Pressed: func(ctx events.Context) {
			Gameover.stateData.Actions.Gameover()
			Gameover.stateData.CurrentState = Gameover.stateData.States.Menu
		},
	})
}

func (p *gameoverState) Update(data *state.Data) {
	p.stateData = data
	p.ProcessEvents()
}

func (p *gameoverState) Draw(screen *ebiten.Image, data *state.Data) {
	data.States.Play.Draw(screen, data)

	shadow := ebiten.NewImage(config.ScreenWidth, config.ScreenHeight)
	shadow.Fill(color.RGBA{R: 0, G: 0, B: 0, A: 200})
	screen.DrawImage(shadow, &ebiten.DrawImageOptions{})
}
