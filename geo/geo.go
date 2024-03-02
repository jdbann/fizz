package geo

type Point struct{ X, Y int }

func (p Point) Is(x, y int) bool {
	return p.X == x && p.Y == y
}

type Size struct{ Width, Height int }

type Rect struct {
	Position Point
	Size     Size
}

func RectAroundPoint(point Point, size Size) Rect {
	return Rect{
		Position: Point{
			X: point.X - (size.Width / 2),
			Y: point.Y - (size.Height / 2),
		},
		Size: size,
	}
}

func (r Rect) Bounds() (min, max Point) {
	return r.Position,
		Point{r.Position.X + r.Size.Width, r.Position.Y + r.Size.Height}
}

func (r Rect) Contains(p Point) bool {
	return p.X >= r.Position.X && p.X < r.Position.X+r.Size.Width &&
		p.Y >= r.Position.Y && p.Y < r.Position.Y+r.Size.Height
}

func (r Rect) WithOrigin(origin Point) Rect {
	return Rect{
		Position: Point{
			X: r.Position.X + origin.X,
			Y: r.Position.Y + origin.Y,
		},
		Size: r.Size,
	}
}
