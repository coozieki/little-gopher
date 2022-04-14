package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"go-snake/internal/game"
	"go-snake/internal/snake"
	"log"
)

func main() {
	g := game.NewGame(snake.NewSnake())
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Snake")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
