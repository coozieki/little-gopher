package menu

import (
	"github.com/hajimehoshi/ebiten/v2"
	"go-snake/internal/components"
	"go-snake/internal/config"
	"go-snake/internal/events"
	"go-snake/internal/state"
)

type menuLayer struct {
	image         *ebiten.Image
	startButton   *components.Button
	optionsButton *components.Button
	exitButton    *components.Button
}

func (m *menuLayer) Render() {
	m.image.Clear()
	m.image.Fill(config.BGColor)

	m.startButton.Render(
		"New Game",
		config.ScreenWidth/2-components.ButtonWidth/2,
		config.ScreenHeight/2-components.ButtonHeight/2,
	)
	m.optionsButton.Render(
		"Options",
		config.ScreenWidth/2-components.ButtonWidth/2,
		config.ScreenHeight/2-components.ButtonHeight/2+components.ButtonHeight+10,
	)
	m.exitButton.Render(
		"Exit",
		config.ScreenWidth/2-components.ButtonWidth/2,
		config.ScreenHeight/2-components.ButtonHeight/2+(components.ButtonHeight+10)*2,
	)

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
	events.EventListener
	menuLayer *menuLayer
	stateData *state.Data
}

var Menu *menuState

func (m *menuState) Draw(screen *ebiten.Image, _ *state.Data) {
	screen.DrawImage(m.menuLayer.image, &ebiten.DrawImageOptions{})
}

func (m *menuState) Update(data *state.Data) {
	m.stateData = data
	m.ProcessEvents()
}
