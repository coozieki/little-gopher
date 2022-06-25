package events

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type ButtonPressHandler struct {
	Key     ebiten.Key
	Pressed func(ctx Context)
}

func (ev *eventListener) ButtonPress(handler *ButtonPressHandler) interface{} {
	return ev.registerEvent(buttonPress, nil, handler)
}

func processButtonPressEvent(ctx Context, ev *event) {
	handler := ev.Handler.(*ButtonPressHandler)

	if inpututil.IsKeyJustPressed(handler.Key) {
		handler.Pressed(ctx)
	}
}
