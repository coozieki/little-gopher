package menu

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"go-snake/internal/components"
	"go-snake/internal/config"
	"go-snake/internal/geom"
	"go-snake/internal/state"
)

type buttonContext struct {
	m    *menuState
	data *state.Data
}

type menuLayer struct {
	image         *ebiten.Image
	startButton   *components.Button
	optionsButton *components.Button
	exitButton    *components.Button
}

func (m *menuLayer) render() {
	m.image.Clear()
	m.image.Fill(config.BGColor)

	m.startButton.Render("New Game", config.ScreenWidth/2-components.ButtonWidth/2, config.ScreenHeight/2-components.ButtonHeight/2)
	m.optionsButton.Render("Options", config.ScreenWidth/2-components.ButtonWidth/2, config.ScreenHeight/2-components.ButtonHeight/2+components.ButtonHeight+10)
	m.exitButton.Render("Exit", config.ScreenWidth/2-components.ButtonWidth/2, config.ScreenHeight/2-components.ButtonHeight/2+(components.ButtonHeight+10)*2)

	draw := func(b *components.Button) {
		geoM := ebiten.GeoM{}
		geoM.Translate(float64(b.Rect.X), float64(b.Rect.Y))
		m.image.DrawImage(b.Image, &ebiten.DrawImageOptions{GeoM: geoM})
	}
	draw(m.startButton)
	draw(m.optionsButton)
	draw(m.exitButton)
}

type menuState struct {
	menuLayer menuLayer
}

var Menu *menuState

func (m *menuState) Draw(screen *ebiten.Image, data *state.Data) {
	screen.DrawImage(m.menuLayer.image, &ebiten.DrawImageOptions{})
}

func (m *menuState) Update(data *state.Data) {
	x, y := ebiten.CursorPosition()
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		data.CurrentState = data.States.Play
		return
	}
	buttons := []*components.Button{m.menuLayer.startButton, m.menuLayer.optionsButton, m.menuLayer.exitButton}
	for _, b := range buttons {
		if b.Rect.Contains(geom.Point{X: x, Y: y}) {
			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
				b.Action(buttonContext{m: m, data: data})
			} else {
				b.Active = true
				m.menuLayer.render()
			}
			return
		}
	}
	for _, b := range buttons {
		if b.Active {
			b.Active = false
			m.menuLayer.render()
		}
	}
}
