package geo

type Point struct{ X, Y int }

func (p Point) Is(x, y int) bool {
	return p.X == x && p.Y == y
}

type Size struct{ Width, Height int }

type Rect struct {
	Origin Point
	Size   Size
}

func RectAroundPoint(point Point, size Size) Rect {
	return Rect{
		Origin: Point{
			X: point.X - (size.Width / 2),
			Y: point.Y - (size.Height / 2),
		},
		Size: size,
	}
}

func (r Rect) Bounds() (min, max Point) {
	return r.Origin,
		Point{r.Origin.X + r.Size.Width, r.Origin.Y + r.Size.Height}
}
