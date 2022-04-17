package play

import (
	"github.com/hajimehoshi/ebiten/v2"
	"go-snake/internal/config"
	"go-snake/internal/snake"
	"go-snake/internal/state"
	"sync"
)

func (p *playState) Draw(screen *ebiten.Image, data *state.Data) {
	var wg sync.WaitGroup

	screen.Fill(config.BGColor)

	p.SnakeLayer.Clear()
	blocks := p.snake.GetBlocks()
	for i := 0; i < len(blocks); i++ {
		block := blocks[i]
		switch block.Direction {
		case snake.DirectionRight:
			block.X += -1 + p.movementOffset
		case snake.DirectionUp:
			block.Y -= -1 + p.movementOffset
		case snake.DirectionLeft:
			block.X -= -1 + p.movementOffset
		case snake.DirectionDown:
			block.Y += -1 + p.movementOffset
		}
		i := i
		p.drawBlock(p.SnakeLayer, block, i == len(blocks)-1)
	}

	wg.Add(3)
	for _, layer := range []*ebiten.Image{p.MainGameLayer, p.FruitLayer, p.SnakeLayer} {
		layer := layer
		go func() {
			defer wg.Done()
			screen.DrawImage(layer, &ebiten.DrawImageOptions{})
		}()
	}

	wg.Wait()
}

func (p *playState) drawBlock(target *ebiten.Image, block snake.Block, isHead bool) {
	geoM := ebiten.GeoM{}
	geoM.Translate(block.X*float64(config.BlockWidth), block.Y*float64(config.BlockWidth))

	var img *ebiten.Image
	if isHead {
		img = p.HeadImage
	} else {
		img = p.BlockImage
	}
	target.DrawImage(img, &ebiten.DrawImageOptions{GeoM: geoM})
}

func (p *playState) drawFruit(target *ebiten.Image, x, y int) {
	geoM := ebiten.GeoM{}
	geoM.Translate(float64(x*config.BlockWidth), float64(y*config.BlockWidth))
	target.DrawImage(p.FruitImage, &ebiten.DrawImageOptions{GeoM: geoM})
}

func (p *playState) drawObstacle(target *ebiten.Image, x, y float64) {
	geoM := ebiten.GeoM{}
	geoM.Translate(x*float64(config.BlockWidth), y*float64(config.BlockWidth))
	target.DrawImage(p.ObstacleImage, &ebiten.DrawImageOptions{GeoM: geoM})
}
