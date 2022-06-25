package events

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Clickable interface {
	IsHovered(ctx Context) bool
}

type ClickHandler struct {
	Pressed  func(ctx Context)
	Released func(ctx Context)
}

func (ev *eventListener) Click(component Clickable, handler *ClickHandler) interface{} {
	return ev.registerEvent(click, component, handler)
}

func processClickEvent(ctx Context, ev *event, comp interface{}) {
	convComp := comp.(Clickable)
	handler := ev.Handler.(*ClickHandler)

	if convComp.IsHovered(ctx) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		handler.Pressed(ctx)
	}
}
