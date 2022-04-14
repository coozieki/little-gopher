package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"go-snake/internal/snake"
	"golang.org/x/image/colornames"
	"image/color"
)

// Game implements ebiten.Game interface.
type game struct {
	snake   snake.Snake
	offsetX float64
	offsetY float64
}

func NewGame(snake snake.Snake) *game {
	return &game{snake: snake}
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *game) Update() error {
	for _, key := range []ebiten.Key{ebiten.KeyS, ebiten.KeyD, ebiten.KeyW, ebiten.KeyA, ebiten.KeyF} {
		if inpututil.IsKeyJustPressed(key) {
			switch key {
			case ebiten.KeyS:
				g.snake.Move(snake.DirectionDown)
			case ebiten.KeyD:
				g.snake.Move(snake.DirectionRight)
			case ebiten.KeyW:
				g.snake.Move(snake.DirectionUp)
			case ebiten.KeyA:
				g.snake.Move(snake.DirectionLeft)
			case ebiten.KeyF:
				g.snake.PushBlock()
			}
		}
	}
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Aquamarine)
	for i, block := range g.snake.GetBlocks() {
		rectInner := ebiten.NewImage(18, 18)
		if i == len(g.snake.GetBlocks())-1 {
			rectInner.Fill(colornames.Green)
		} else {
			rectInner.Fill(color.White)
		}
		geoM := ebiten.GeoM{}
		geoM.Translate(1, 1)
		rect := ebiten.NewImage(20, 20)
		rect.Fill(color.Black)
		rect.DrawImage(rectInner, &ebiten.DrawImageOptions{GeoM: geoM})
		geoM = ebiten.GeoM{}
		geoM.Translate(block.X*19, block.Y*19)
		screen.DrawImage(rect, &ebiten.DrawImageOptions{GeoM: geoM})
	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
