package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"go-snake/internal/game"
	"log"
)

func main() {
	g := game.NewGame()
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(800, 800)
	ebiten.SetWindowTitle("Snake")
	ebiten.SetScreenClearedEveryFrame(false)
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
