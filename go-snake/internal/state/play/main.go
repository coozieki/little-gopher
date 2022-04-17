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

	mainGameLayer *ebiten.Image
	snakeLayer    *ebiten.Image
	fruitLayer    *ebiten.Image

	fruitImage    *ebiten.Image
	obstacleImage *ebiten.Image
	headImage     *ebiten.Image
	blockImage    *ebiten.Image

	snake snake.Snake
	fruit geom.Point
}

var Play *playState
