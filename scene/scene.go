package scene

import (
	"strings"

	"github.com/jdbann/fizz/geo"
)

type Scene struct {
	sprites []*Sprite
}

func (s *Scene) AddSprite(sprite *Sprite) {
	s.sprites = append(s.sprites, sprite)
}

func (s Scene) RenderString(size geo.Size) (string, error) {
	min, max := geo.RectAroundPoint(geo.Point{}, size).Bounds()

	var b strings.Builder

	for y := min.Y; y < max.Y; y++ {
	XLoop:
		for x := min.X; x < max.X; x++ {
			for _, sprite := range s.sprites {
				if !sprite.rect.Contains(geo.Point{X: x, Y: y}) {
					continue
				}

				if _, err := b.WriteRune(sprite.RuneAt(geo.Point{X: x, Y: y})); err != nil {
					return "", err
				}

				continue XLoop
			}

			if _, err := b.WriteRune(' '); err != nil {
				return "", err
			}
		}

		if y < max.Y-1 {
			if _, err := b.WriteRune('\n'); err != nil {
				return "", err
			}
		}
	}

	return b.String(), nil
}
