package options

import (
	"github.com/hajimehoshi/ebiten/v2"
	"go-snake/internal/components"
	"go-snake/internal/config"
	"go-snake/internal/events"
	"go-snake/internal/state"
)

type buttonsLayer struct {
	image      *ebiten.Image
	backButton *components.Button
}

func (b *buttonsLayer) Render() {
	b.backButton.Render("Back", config.ScreenWidth/2-components.ButtonWidth/2, config.ScreenHeight/2-components.ButtonHeight/2)

	geoM := ebiten.GeoM{}
	geoM.Translate(float64(b.backButton.Rect.X), float64(b.backButton.Rect.Y))
	b.image.DrawImage(b.backButton.Image, &ebiten.DrawImageOptions{GeoM: geoM})
}

type optionsState struct {
	events.EventListener
	optionsImage *ebiten.Image
	buttonsLayer *buttonsLayer
	stateData    *state.Data
}

var Options *optionsState

func init() {
	Options = &optionsState{}
	Options.EventListener = events.NewEventListener()

	optionsImage := ebiten.NewImage(config.ScreenWidth, config.ScreenHeight)
	optionsImage.Fill(config.BGColor)

	Options.optionsImage = optionsImage

	img := ebiten.NewImage(config.ScreenWidth, config.ScreenHeight)

	Options.buttonsLayer = &buttonsLayer{
		image: img,
	}

	backButton := &components.Button{
		Parent: Options.buttonsLayer,
	}
	Options.Hover(backButton, &events.HoverHandler{})
	Options.Click(backButton, &events.ClickHandler{
		Pressed: func(ctx events.Context) {
			Options.stateData.CurrentState = Options.stateData.States.Menu
		},
	})
	Options.buttonsLayer.backButton = backButton

	Options.ButtonPress(&events.ButtonPressHandler{
		Key: ebiten.KeyEscape,
		Pressed: func(ctx events.Context) {
			Options.stateData.CurrentState = Options.stateData.States.Menu
		},
	})

	Options.buttonsLayer.Render()
}

func (o *optionsState) Update(data *state.Data) {
	o.stateData = data
	o.ProcessEvents()
}

func (o *optionsState) Draw(screen *ebiten.Image, _ *state.Data) {
	screen.DrawImage(o.optionsImage, &ebiten.DrawImageOptions{})
	screen.DrawImage(o.buttonsLayer.image, &ebiten.DrawImageOptions{})
}
