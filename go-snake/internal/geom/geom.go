package geom

type Point struct {
	X int
	Y int
}

type Rect struct {
	X      int
	Y      int
	Width  int
	Height int
}

func (r Rect) Contains(point Point) bool {
	return (point.X > r.X && point.X < (r.X+r.Width)) && (point.Y > r.Y && point.Y < (r.Y+r.Height))
}
