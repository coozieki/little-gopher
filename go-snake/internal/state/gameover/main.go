package gameover

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"go-snake/internal/config"
	"go-snake/internal/state"
	"image/color"
)

type gameoverState struct {
}

var Gameover = &gameoverState{}

func (p *gameoverState) Update(data *state.Data) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		data.Actions.Gameover()
		data.CurrentState = data.States.Menu
	}
}

func (p *gameoverState) Draw(screen *ebiten.Image, data *state.Data) {
	data.States.Play.Draw(screen, data)

	shadow := ebiten.NewImage(config.ScreenWidth, config.ScreenHeight)
	shadow.Fill(color.RGBA{R: 0, G: 0, B: 0, A: 200})
	screen.DrawImage(shadow, &ebiten.DrawImageOptions{})
}
