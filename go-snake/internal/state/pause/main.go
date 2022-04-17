package pause

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"go-snake/internal/config"
	"go-snake/internal/state"
	"go-snake/internal/state/play"
	"image/color"
)

type pauseState struct {
}

var Pause = &pauseState{}

func (p *pauseState) Update(data *state.Data) {
	if inpututil.IsKeyJustPressed(config.PauseKey) {
		data.CurrentState = data.States.Play
	}
}

func (p *pauseState) Draw(screen *ebiten.Image, data *state.Data) {
	play.Play.Draw(screen, data)

	shadow := ebiten.NewImage(config.ScreenWidth, config.ScreenHeight)
	shadow.Fill(color.RGBA{R: 0, G: 0, B: 0, A: 200})
	screen.DrawImage(shadow, &ebiten.DrawImageOptions{})
}
