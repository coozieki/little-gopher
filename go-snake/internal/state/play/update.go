package play

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"go-snake/internal/config"
	"go-snake/internal/snake"
	"go-snake/internal/state"
	"math"
	"math/rand"
	"sync"
)

func (p *playState) Update(data *state.Data) {
	p.processInputs(data)

	p.framesFromLastMove++
	if p.framesFromLastMove > config.MoveIntervalInFrames {
		if len(p.dirQueue) > 0 {
			p.snake.ChangeDir(p.dirQueue[0])
			p.mu.Lock()
			p.dirQueue = p.dirQueue[1:]
			p.mu.Unlock()
		}

		headX, headY := p.snake.GetNextHeadCoords()
		if p.fruit.X == int(headX) && p.fruit.Y == int(headY) {
			p.snake.PushBlock()
			p.moveFruit()
		} else {
			err := p.snake.Move()
			if err != nil {
				data.CurrentState = data.States.Gameover
				return
			}
		}
		p.framesFromLastMove = 0
		p.movementOffset = 0
	} else {
		p.movementOffset = p.framesFromLastMove / config.MoveIntervalInFrames
	}
}

func (p *playState) GameOver() {
	p.snake = snake.NewSnake()
	p.snakeLayer = nil
	p.mainGameLayer = nil
	p.fruitLayer = nil
	p.dirQueue = []snake.Direction{}
	p.initLayers()
	p.moveFruit()
}

func (p *playState) processInputs(data *state.Data) {
	var wg sync.WaitGroup
	for _, key := range []ebiten.Key{ebiten.KeyS, ebiten.KeyD, ebiten.KeyW, ebiten.KeyA, config.PauseKey, config.MenuKey} {
		wg.Add(1)
		key := key
		go func() {
			defer wg.Done()

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
				case config.PauseKey:
					p.mu.Lock()
					data.CurrentState = data.States.Pause
					p.mu.Unlock()
					return
				default:
					return
				}
				if len(p.dirQueue) <= config.MaxDirQueueSize {
					p.mu.Lock()
					p.dirQueue = append(p.dirQueue, *newDir)
					p.mu.Unlock()
				}
			}
		}()
	}
	wg.Wait()
}

func (p *playState) moveFruit() {
	var newX, newY float64
	for {
		newX, newY = math.Floor(rand.Float64()*config.FieldWidthInBlocks), math.Floor(rand.Float64()*config.FieldWidthInBlocks)
		if !p.snake.BlockExistsAt(newX, newY) {
			break
		}
	}
	p.fruit.X = int(newX)
	p.fruit.Y = int(newY)

	p.fruitLayer.Clear()
	p.drawFruit(p.fruitLayer, p.fruit.X, p.fruit.Y)
}
