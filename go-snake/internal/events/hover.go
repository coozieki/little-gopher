package events

type Hoverable interface {
	IsHovered(ctx Context) bool
	OnHover(ctx Context)
	OnUnhover(ctx Context)
}

type HoverHandler struct {
	isHovered bool
	Start     func(ctx Context)
	Stop      func(ctx Context)
}

func (ev *eventListener) Hover(component Hoverable, handler *HoverHandler) interface{} {
	return ev.registerEvent(hover, component, handler)
}

func processHoverEvent(ctx Context, ev *event, comp interface{}) {
	convComp := comp.(Hoverable)
	handler := ev.Handler.(*HoverHandler)
	if !handler.isHovered && convComp.IsHovered(ctx) {
		convComp.OnHover(ctx)
		if handler.Start != nil {
			handler.Start(ctx)
		}
		handler.isHovered = true
	} else if handler.isHovered && !convComp.IsHovered(ctx) {
		convComp.OnUnhover(ctx)
		if handler.Stop != nil {
			handler.Stop(ctx)
		}
		handler.isHovered = false
	}
}
