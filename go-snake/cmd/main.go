package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"go-snake/internal/game"
	"log"
)

func main() {
	g := game.NewGame()

	ebiten.SetWindowSize(800, 800)
	ebiten.SetWindowTitle("Snake")
	ebiten.SetScreenClearedEveryFrame(false)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
