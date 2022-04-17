package play

import (
	"github.com/hajimehoshi/ebiten/v2"
	"go-snake/internal/config"
	"go-snake/internal/snake"
	"golang.org/x/image/colornames"
	"image/color"
	"sync"
)

func init() {
	Play = &playState{}
	Play.snake = snake.NewSnake()
	Play.initImages()
	Play.initLayers()
	Play.moveFruit()
}

func (p *playState) initImages() {
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

	p.fruitImage = fruit
	p.headImage = head
	p.blockImage = block
	p.obstacleImage = obstacle
}

func (p *playState) initLayers() {
	var wg sync.WaitGroup

	var screenWidth, screenHeight = config.ScreenWidth, config.ScreenHeight

	mainLayer := ebiten.NewImage(screenWidth, screenHeight)
	fruitLayer := ebiten.NewImage(screenWidth, screenHeight)
	snakeLayer := ebiten.NewImage(screenWidth, screenHeight)

	for x, v := range config.Map {
		for y := range v {
			wg.Add(1)
			y := y
			x := x

			go func() {
				defer wg.Done()
				p.drawObstacle(mainLayer, x, y)
			}()
		}
	}

	wg.Wait()

	geoM := ebiten.GeoM{}
	geoM.Translate(float64(p.fruit.X*config.BlockWidth), float64(p.fruit.Y*config.BlockWidth))
	fruitLayer.DrawImage(p.fruitImage, &ebiten.DrawImageOptions{GeoM: geoM})

	Play.fruitLayer = fruitLayer
	Play.snakeLayer = snakeLayer
	Play.mainGameLayer = mainLayer
}
