package scene

import (
	"github.com/jdbann/fizz/geo"
	"github.com/mattn/go-runewidth"
)

type Sprite struct {
	position geo.Point
	texture  string
}

func NewSprite(position geo.Point, texture string) *Sprite {
	if runewidth.StringWidth(texture) > 1 {
		panic("cannot handle textures wider than one cell")
	}

	return &Sprite{
		position: position,
		texture:  texture,
	}
}

func (s Sprite) Position() geo.Point {
	return s.position
}

func (s Sprite) Texture() string {
	return string(s.texture)
}
