package components

import (
	"go-snake/internal/events"
	"go-snake/internal/geom"
)

type rectangleShape struct {
	Rect geom.Rect
}

func (r *rectangleShape) IsHovered(ctx events.Context) bool {
	return r.Rect.Contains(geom.Point{X: ctx.Cursor.X, Y: ctx.Cursor.Y})
}
