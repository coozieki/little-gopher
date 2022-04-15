package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"go-snake/internal/config"
	"go-snake/internal/snake"
	"golang.org/x/image/colornames"
	"image/color"
	"math"
	"math/rand"
	"sync"
)

const (
	maxDirQueueSize   = 3
	drawFrameInterval = 4
	framesForMove     = drawFrameInterval * 3
)

type fruit struct {
	X float64
	Y float64
}

// Game implements ebiten.Game interface.
type game struct {
	snake snake.Snake
	fruit fruit

	framesFromLastMove int
	frames             int
	dirQueue           []snake.Direction
	mu                 sync.Mutex

	fruitImage *ebiten.Image
	blockImage *ebiten.Image
	headImage  *ebiten.Image
}

func NewGame(s snake.Snake) *game {
	g := &game{snake: s, fruit: fruit{}}
	g.initImages()
	g.moveFruit()
	return g
}

func (g *game) initImages() {
	fruitInner := ebiten.NewImage(config.BlockWidth-1, config.BlockWidth-1)
	fruitInner.Fill(colornames.Yellow)
	geoM := ebiten.GeoM{}
	geoM.Translate(1, 1)
	fruit := ebiten.NewImage(config.BlockWidth+1, config.BlockWidth+1)
	fruit.Fill(colornames.Black)
	fruit.DrawImage(fruitInner, &ebiten.DrawImageOptions{GeoM: geoM})

	headInner := ebiten.NewImage(config.BlockWidth-1, config.BlockWidth-1)
	headInner.Fill(colornames.Green)
	blockInner := ebiten.NewImage(config.BlockWidth-1, config.BlockWidth-1)
	blockInner.Fill(color.White)

	geoM = ebiten.GeoM{}
	geoM.Translate(1, 1)

	head := ebiten.NewImage(config.BlockWidth+1, config.BlockWidth+1)
	head.Fill(color.Black)
	block := ebiten.NewImage(config.BlockWidth+1, config.BlockWidth+1)
	block.Fill(color.Black)

	head.DrawImage(headInner, &ebiten.DrawImageOptions{GeoM: geoM})
	block.DrawImage(blockInner, &ebiten.DrawImageOptions{GeoM: geoM})

	g.fruitImage = fruit
	g.headImage = head
	g.blockImage = block
}

func (g *game) moveFruit() {
	var newX, newY float64
	for {
		newX, newY = math.Floor(rand.Float64()*16), math.Floor(rand.Float64()*16)
		if !g.snake.BlockExistsAt(newX, newY) {
			break
		}
	}
	g.fruit.X = newX
	g.fruit.Y = newY
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *game) Update() error {
	go func() {
		for _, key := range []ebiten.Key{ebiten.KeyS, ebiten.KeyD, ebiten.KeyW, ebiten.KeyA, ebiten.KeyF} {
			if inpututil.IsKeyJustPressed(key) {
				newDir := new(snake.Direction)
				switch key {
				case ebiten.KeyS:
					*newDir = snake.DirectionDown
				case ebiten.KeyD:
					*newDir = snake.DirectionRight
				case ebiten.KeyW:
					*newDir = snake.DirectionUp
				case ebiten.KeyA:
					*newDir = snake.DirectionLeft
				default:
					newDir = nil
				}
				if newDir != nil && len(g.dirQueue) <= maxDirQueueSize {
					g.mu.Lock()
					g.dirQueue = append(g.dirQueue, *newDir)
					g.mu.Unlock()
				}
			}
		}
	}()
	g.framesFromLastMove++
	if g.framesFromLastMove > framesForMove {
		if len(g.dirQueue) > 0 {
			g.snake.ChangeDir(g.dirQueue[0])
			g.mu.Lock()
			g.dirQueue = g.dirQueue[1:]
			g.mu.Unlock()
		}

		headX, headY := g.snake.GetNextHeadCoords()
		if g.fruit.X == headX && g.fruit.Y == headY {
			g.snake.PushBlock()
			g.moveFruit()
		} else {
			g.snake.Move()
		}
		g.framesFromLastMove = 0
	}
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *game) Draw(screen *ebiten.Image) {
	g.frames++
	if g.frames%drawFrameInterval != 0 {
		return
	}
	var wg sync.WaitGroup

	screen.Clear()
	screen.Fill(colornames.Aquamarine)

	blocks := g.snake.GetBlocks()
	for i, block := range blocks {
		wg.Add(1)
		block := block
		i := i
		go func() {
			defer wg.Done()
			g.drawBlock(screen, block, i == len(blocks)-1)
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()

		geoM := ebiten.GeoM{}
		geoM.Translate(g.fruit.X*float64(config.BlockWidth), g.fruit.Y*float64(config.BlockWidth))
		screen.DrawImage(g.fruitImage, &ebiten.DrawImageOptions{GeoM: geoM})
	}()

	wg.Wait()
}

func (g *game) drawBlock(screen *ebiten.Image, block snake.Block, isHead bool) {
	geoM := ebiten.GeoM{}
	geoM.Translate(block.X*float64(config.BlockWidth), block.Y*float64(config.BlockWidth))

	var img *ebiten.Image
	if isHead {
		img = g.headImage
	} else {
		img = g.blockImage
	}
	screen.DrawImage(img, &ebiten.DrawImageOptions{GeoM: geoM})
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.BlockWidth*config.FieldWidthInBlocks + 1, config.BlockWidth*config.FieldWidthInBlocks + 1
}
