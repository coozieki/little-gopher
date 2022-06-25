package pause

import (
	"github.com/hajimehoshi/ebiten/v2"
	"go-snake/internal/config"
	"go-snake/internal/events"
	"go-snake/internal/state"
	"image/color"
)

type pauseState struct {
	events.EventListener
	stateData *state.Data
}

var Pause *pauseState

func init() {
	Pause = &pauseState{EventListener: events.NewEventListener()}

	Pause.ButtonPress(&events.ButtonPressHandler{
		Key: config.PauseKey,
		Pressed: func(ctx events.Context) {
			Pause.stateData.CurrentState = Pause.stateData.States.Play
		},
	})
}

func (p *pauseState) Update(data *state.Data) {
	p.stateData = data
	p.ProcessEvents()
}

func (p *pauseState) Draw(screen *ebiten.Image, data *state.Data) {
	data.States.Play.Draw(screen, data)

	shadow := ebiten.NewImage(config.ScreenWidth, config.ScreenHeight)
	shadow.Fill(color.RGBA{R: 0, G: 0, B: 0, A: 200})
	screen.DrawImage(shadow, &ebiten.DrawImageOptions{})
}
