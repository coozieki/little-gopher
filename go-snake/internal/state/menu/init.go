package menu

import (
	"github.com/hajimehoshi/ebiten/v2"
	"go-snake/internal/components"
	"go-snake/internal/config"
	"go-snake/internal/events"
	"image/color"
	"os"
)

func init() {
	Menu = &menuState{}
	Menu.EventListener = events.NewEventListener()

	menuLayerImg := ebiten.NewImage(config.ScreenWidth, config.ScreenHeight)
	menuLayerImg.Fill(color.White)

	menuLayer := &menuLayer{image: menuLayerImg}

	menuLayer.startButton = &components.Button{Parent: menuLayer}
	Menu.Click(menuLayer.startButton, &events.ClickHandler{
		Pressed: func(ctx events.Context) {
			Menu.stateData.CurrentState = Menu.stateData.States.Play
		},
	})

	menuLayer.optionsButton = &components.Button{Parent: menuLayer}
	Menu.Click(menuLayer.optionsButton, &events.ClickHandler{
		Pressed: func(ctx events.Context) {
			Menu.stateData.CurrentState = Menu.stateData.States.Options
		},
	})

	menuLayer.exitButton = &components.Button{Parent: menuLayer}
	Menu.Click(menuLayer.exitButton, &events.ClickHandler{
		Pressed: func(ctx events.Context) {
			os.Exit(1)
		},
	})

	Menu.menuLayer = menuLayer

	buttons := []*components.Button{Menu.menuLayer.startButton, Menu.menuLayer.optionsButton, Menu.menuLayer.exitButton}
	for _, button := range buttons {
		Menu.Hover(button, &events.HoverHandler{})
	}

	Menu.ButtonPress(&events.ButtonPressHandler{
		Key: ebiten.KeyEnter,
		Pressed: func(ctx events.Context) {
			Menu.stateData.CurrentState = Menu.stateData.States.Play
		},
	})

	Menu.menuLayer.Render()
}
