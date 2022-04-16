package game

import (
	"fmt"
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
	maxDirQueueSize      = 3
	drawIntervalInFrames = 4
	moveIntervalInFrames = drawIntervalInFrames * 1
)

type fruit struct {
	X float64
	Y float64
}

type images struct {
	fruitImage    *ebiten.Image
	blockImage    *ebiten.Image
	headImage     *ebiten.Image
	obstacleImage *ebiten.Image
}

type layers struct {
	mainLayer  *ebiten.Image
	snakeLayer *ebiten.Image
	fruitLayer *ebiten.Image
}

type game struct {
	snake snake.Snake
	fruit fruit

	framesFromLastMove int
	frames             int
	dirQueue           []snake.Direction
	mu                 sync.Mutex

	images images
	layers layers
}

func NewGame(s snake.Snake) *game {
	g := &game{snake: s, fruit: fruit{}}
	g.initImages()
	g.initLayers()
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

	obstacleInner := ebiten.NewImage(config.BlockWidth-1, config.BlockWidth-1)
	obstacleInner.Fill(colornames.Red)
	geoM = ebiten.GeoM{}
	geoM.Translate(1, 1)
	obstacle := ebiten.NewImage(config.BlockWidth+1, config.BlockWidth+1)
	obstacle.Fill(colornames.Black)
	obstacle.DrawImage(obstacleInner, &ebiten.DrawImageOptions{GeoM: geoM})

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

	g.images = images{
		fruitImage:    fruit,
		headImage:     head,
		blockImage:    block,
		obstacleImage: obstacle,
	}
}

func (g *game) initLayers() {
	var wg sync.WaitGroup

	mainLayer := ebiten.NewImage(g.Layout(1, 1))
	fruitLayer := ebiten.NewImage(g.Layout(1, 1))
	snakeLayer := ebiten.NewImage(g.Layout(1, 1))

	for x, v := range config.Map {
		for y := range v {
			wg.Add(1)
			y := y
			x := x

			go func() {
				defer wg.Done()
				g.drawObstacle(mainLayer, x, y)
			}()
		}
	}

	wg.Wait()

	g.layers = layers{mainLayer: mainLayer, fruitLayer: fruitLayer, snakeLayer: snakeLayer}
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

	g.layers.fruitLayer.Clear()
	geoM := ebiten.GeoM{}
	geoM.Translate(g.fruit.X*float64(config.BlockWidth), g.fruit.Y*float64(config.BlockWidth))
	g.layers.fruitLayer.DrawImage(g.images.fruitImage, &ebiten.DrawImageOptions{GeoM: geoM})
}

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
	if g.framesFromLastMove > moveIntervalInFrames {
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

func (g *game) Draw(screen *ebiten.Image) {
	g.frames++
	var wg sync.WaitGroup

	screen.Clear()
	screen.Fill(colornames.Aquamarine)

	g.layers.snakeLayer.Clear()
	blocks := g.snake.GetBlocks()
	for i, block := range blocks {
		wg.Add(1)
		block := block
		i := i
		go func() {
			defer wg.Done()
			g.drawBlock(g.layers.snakeLayer, block, i == len(blocks)-1)
		}()
	}

	wg.Wait()

	wg.Add(3)
	for _, layer := range []*ebiten.Image{g.layers.mainLayer, g.layers.fruitLayer, g.layers.snakeLayer} {
		layer := layer
		go func() {
			defer wg.Done()
			screen.DrawImage(layer, &ebiten.DrawImageOptions{})
		}()
	}

	wg.Wait()

	fmt.Println(ebiten.CurrentFPS())
}

func (g *game) drawBlock(screen *ebiten.Image, block snake.Block, isHead bool) {
	geoM := ebiten.GeoM{}
	geoM.Translate(block.X*float64(config.BlockWidth), block.Y*float64(config.BlockWidth))

	var img *ebiten.Image
	if isHead {
		img = g.images.headImage
	} else {
		img = g.images.blockImage
	}
	screen.DrawImage(img, &ebiten.DrawImageOptions{GeoM: geoM})
}

func (g *game) drawObstacle(target *ebiten.Image, x, y float64) {
	geoM := ebiten.GeoM{}
	geoM.Translate(x*float64(config.BlockWidth), y*float64(config.BlockWidth))
	target.DrawImage(g.images.obstacleImage, &ebiten.DrawImageOptions{GeoM: geoM})
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *game) Layout(int, int) (screenWidth int, screenHeight int) {
	return config.BlockWidth*config.FieldWidthInBlocks + 1, config.BlockWidth*config.FieldWidthInBlocks + 1
}
