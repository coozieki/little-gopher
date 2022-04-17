package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"go-snake/internal/config"
	"go-snake/internal/state"
	"go-snake/internal/state/gameover"
	"go-snake/internal/state/menu"
	"go-snake/internal/state/options"
	"go-snake/internal/state/pause"
	"go-snake/internal/state/play"
)

type Game struct {
	state state.State
}

func NewGame() *Game {
	g := &Game{}
	g.state = menu.Menu
	return g
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()

	data := g.formStateData()
	g.state.Draw(screen, &data)
}

func (g *Game) Update() error {
	data := g.formStateData()
	g.state.Update(&data)
	g.state = data.CurrentState
	return nil
}

func (g *Game) formStateData() state.Data {
	return state.Data{
		CurrentState: g.state,
		States: struct {
			Menu     state.State
			Play     state.State
			Pause    state.State
			Options  state.State
			Gameover state.State
		}{Menu: menu.Menu, Play: play.Play, Pause: pause.Pause, Options: options.Options, Gameover: gameover.Gameover},
		Actions: struct {
			Gameover func()
		}{
			Gameover: play.Play.GameOver,
		},
	}
}

func (g *Game) Layout(int, int) (screenWidth int, screenHeight int) {
	return config.ScreenWidth, config.ScreenHeight
}
