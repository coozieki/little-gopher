package components

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"go-snake/internal/events"
	"go-snake/internal/geom"
	"golang.org/x/image/font/basicfont"
	"image/color"
)

const (
	ButtonWidth  = 100
	ButtonHeight = 20
)

type Button struct {
	rectangleShape
	Image  *ebiten.Image
	Active bool
	Parent Renderer
}

func (b *Button) Render(txt string, x, y int) {
	buttonRect := geom.Rect{
		X:      x,
		Y:      y,
		Width:  ButtonWidth,
		Height: ButtonHeight,
	}

	textRect := text.BoundString(basicfont.Face7x13, txt)

	button := ebiten.NewImage(ButtonWidth, ButtonHeight)
	if b.Active {
		button.Fill(color.RGBA{R: 255, G: 169, B: 0, A: 220})
	} else {
		button.Fill(color.RGBA{R: 255, G: 169, B: 0, A: 255})
	}
	text.Draw(button, txt, basicfont.Face7x13, ButtonWidth/2-textRect.Size().X/2, ButtonHeight/2+textRect.Size().Y/2, color.Black)

	b.Image = button
	b.Rect = buttonRect
}

func (b *Button) OnHover(_ events.Context) {
	b.Active = true
	b.Parent.Render()
}

func (b *Button) OnUnhover(_ events.Context) {
	b.Active = false
	b.Parent.Render()
}
