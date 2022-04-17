package play

import (
	"github.com/hajimehoshi/ebiten/v2"
	"go-snake/internal/geom"
	"go-snake/internal/snake"
	"sync"
)

type playState struct {
	movementOffset     float64
	framesFromLastMove float64
	dirQueue           []snake.Direction
	mu                 sync.Mutex

	MainGameLayer *ebiten.Image
	SnakeLayer    *ebiten.Image
	FruitLayer    *ebiten.Image

	FruitImage    *ebiten.Image
	ObstacleImage *ebiten.Image
	HeadImage     *ebiten.Image
	BlockImage    *ebiten.Image

	snake snake.Snake
	fruit geom.Point
}

var Play *playState
