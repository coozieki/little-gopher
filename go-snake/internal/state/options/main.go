package options

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"go-snake/internal/components"
	"go-snake/internal/config"
	"go-snake/internal/geom"
	"go-snake/internal/state"
)

type buttonContext struct {
	o    *optionsState
	data *state.Data
}

type buttonsLayer struct {
	image      *ebiten.Image
	backButton *components.Button
}

func (b *buttonsLayer) render() {
	b.backButton.Render("Back", config.ScreenWidth/2-components.ButtonWidth/2, config.ScreenHeight/2-components.ButtonHeight/2)

	geoM := ebiten.GeoM{}
	geoM.Translate(float64(b.backButton.Rect.X), float64(b.backButton.Rect.Y))
	b.image.DrawImage(b.backButton.Image, &ebiten.DrawImageOptions{GeoM: geoM})
}

type optionsState struct {
	optionsImage *ebiten.Image
	buttonsLayer *buttonsLayer
}

var Options *optionsState

func init() {
	Options = &optionsState{}

	optionsImage := ebiten.NewImage(config.ScreenWidth, config.ScreenHeight)
	optionsImage.Fill(config.BGColor)

	Options.optionsImage = optionsImage

	img := ebiten.NewImage(config.ScreenWidth, config.ScreenHeight)
	Options.buttonsLayer = &buttonsLayer{
		image: img,
		backButton: &components.Button{
			Active: false,
			Action: func(ctx interface{}) {
				buttonCtx := ctx.(buttonContext)
				buttonCtx.data.CurrentState = buttonCtx.data.States.Menu
			},
		},
	}

	Options.buttonsLayer.render()
}

func (o *optionsState) Update(data *state.Data) {
	x, y := ebiten.CursorPosition()
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		data.CurrentState = data.States.Menu
		return
	}
	if o.buttonsLayer.backButton.Rect.Contains(geom.Point{X: x, Y: y}) {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			o.buttonsLayer.backButton.Action(buttonContext{o: o, data: data})
		} else {
			o.buttonsLayer.backButton.Active = true
			o.buttonsLayer.render()
		}
		return
	}
	if o.buttonsLayer.backButton.Active {
		o.buttonsLayer.backButton.Active = false
		o.buttonsLayer.render()
	}
}

func (o *optionsState) Draw(screen *ebiten.Image, data *state.Data) {
	screen.DrawImage(o.optionsImage, &ebiten.DrawImageOptions{})

	screen.DrawImage(o.buttonsLayer.image, &ebiten.DrawImageOptions{})
}
