package menu

import (
	"github.com/hajimehoshi/ebiten/v2"
	"go-snake/internal/components"
	"go-snake/internal/config"
	"image/color"
	"os"
)

func init() {
	Menu = &menuState{}

	menuLayerImg := ebiten.NewImage(config.ScreenWidth, config.ScreenHeight)
	menuLayerImg.Fill(color.White)
	Menu.menuLayer = menuLayer{
		image: menuLayerImg,
		startButton: &components.Button{
			Active: false,
			Action: func(ctx interface{}) {
				buttonCtx := ctx.(buttonContext)
				buttonCtx.data.CurrentState = buttonCtx.data.States.Play
			},
		},
		optionsButton: &components.Button{
			Active: false,
			Action: func(ctx interface{}) {
				buttonCtx := ctx.(buttonContext)
				buttonCtx.data.CurrentState = buttonCtx.data.States.Options
			},
		},
		exitButton: &components.Button{
			Active: false,
			Action: func(ctx interface{}) {
				os.Exit(1)
			},
		},
	}
	Menu.menuLayer.render()
}
