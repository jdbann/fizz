package scene

import "github.com/jdbann/fizz/geo"

type Sprite struct {
	position geo.Point
	texture  rune
}

func NewSprite(position geo.Point, texture rune) *Sprite {
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
