package events

import (
	"github.com/hajimehoshi/ebiten/v2"
	"go-snake/internal/geom"
)

type eventType int

const (
	hover       eventType = 0
	click       eventType = 1
	buttonPress eventType = 2
)

type Context struct {
	Cursor geom.Point
	Event  interface{}
}

type event struct {
	Type    eventType
	Handler interface{}
}

type EventListener interface {
	ProcessEvents()

	Hover(component Hoverable, handler *HoverHandler) interface{}
	Click(component Clickable, handler *ClickHandler) interface{}
	ButtonPress(handler *ButtonPressHandler) interface{}

	ClearEvent(ev interface{})
	ClearEvents()
}

type eventListener struct {
	events map[*event]interface{}
}

func NewEventListener() EventListener {
	return &eventListener{
		events: map[*event]interface{}{},
	}
}

func (ev *eventListener) ProcessEvents() {
	x, y := ebiten.CursorPosition()

	for event, comp := range ev.events {
		ctx := Context{
			Cursor: geom.Point{X: x, Y: y},
			Event:  event,
		}
		switch event.Type {
		case hover:
			processHoverEvent(ctx, event, comp)
			break
		case click:
			processClickEvent(ctx, event, comp)
			break
		case buttonPress:
			processButtonPressEvent(ctx, event)
			break
		}
	}
}

func (ev *eventListener) ClearEvent(evnt interface{}) {
	delete(ev.events, evnt.(*event))
}

func (ev *eventListener) ClearEvents() {
	ev.events = map[*event]interface{}{}
}

func (ev *eventListener) registerEvent(t eventType, component interface{}, handler interface{}) *event {
	newEv := &event{
		Type:    t,
		Handler: handler,
	}
	ev.events[newEv] = component

	return newEv
}
