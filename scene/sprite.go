package scene

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/jdbann/fizz/geo"
	"github.com/mattn/go-runewidth"
)

type Sprite struct {
	position geo.Point
	rect     geo.Rect
	texture  string
}

func NewSprite(position geo.Point, texture string) *Sprite {
	size, err := textureSize(texture)
	if err != nil {
		panic(err)
	}

	return &Sprite{
		position: position,
		rect:     geo.RectAroundPoint(position, size),
		texture:  texture,
	}
}

func (s Sprite) Position() geo.Point {
	return s.position
}

func (s Sprite) RuneAt(p geo.Point) rune {
	rows := strings.Split(s.texture, "\n")
	return []rune(rows[p.Y-s.rect.Position.Y])[p.X-s.rect.Position.X]
}

func (s Sprite) Texture() string {
	return string(s.texture)
}

var (
	errTextureBlank        = errors.New("texture is blank")
	errTextureRuneWidth    = errors.New("texture contains characters wider than one cell")
	errTextureInconsistent = errors.New("texture has inconsistent line widths")
)

func textureSize(t string) (geo.Size, error) {
	if utf8.RuneCountInString(t) == 0 {
		return geo.Size{}, errTextureBlank
	}

	rows := strings.Split(t, "\n")

	s := geo.Size{
		Width:  runewidth.StringWidth(rows[0]),
		Height: len(rows),
	}

	for _, row := range rows {
		if w := runewidth.StringWidth(row); w != s.Width {
			return geo.Size{}, errTextureInconsistent
		}

		for _, char := range row {
			if char == '\n' {
				continue
			}

			if rw := runewidth.RuneWidth(char); rw != 1 {
				return geo.Size{}, fmt.Errorf("%w: %q is width %d", errTextureRuneWidth, char, rw)
			}
		}
	}

	return s, nil
}
