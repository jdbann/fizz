package scene_test

import (
	"testing"

	"github.com/jdbann/fizz/geo"
	"github.com/jdbann/fizz/scene"
	"gotest.tools/v3/assert"
	"gotest.tools/v3/golden"
)

func TestScene_RenderString(t *testing.T) {
	type testCase struct {
		name       string
		sprites    []*scene.Sprite
		renderSize geo.Size
	}

	run := func(t *testing.T, tc testCase) {
		s := &scene.Scene{}

		for _, sprite := range tc.sprites {
			s.AddSprite(sprite)
		}

		view, err := s.RenderString(tc.renderSize)

		assert.NilError(t, err)
		golden.Assert(t, view, t.Name())
	}

	testCases := []testCase{
		{
			name:       "empty scene",
			renderSize: geo.Size{Width: 5, Height: 5},
		},
		{
			name: "central sprite",
			sprites: []*scene.Sprite{
				scene.NewSprite(geo.Point{X: 0, Y: 0}, "X"),
			},
			renderSize: geo.Size{Width: 5, Height: 5},
		},
		{
			name: "ring of sprites",
			sprites: []*scene.Sprite{
				scene.NewSprite(geo.Point{X: -1, Y: -1}, "╭"),
				scene.NewSprite(geo.Point{X: 0, Y: -1}, "─"),
				scene.NewSprite(geo.Point{X: 1, Y: -1}, "╮"),
				scene.NewSprite(geo.Point{X: 1, Y: 0}, "│"),
				scene.NewSprite(geo.Point{X: 1, Y: 1}, "╯"),
				scene.NewSprite(geo.Point{X: 0, Y: 1}, "─"),
				scene.NewSprite(geo.Point{X: -1, Y: 1}, "╰"),
				scene.NewSprite(geo.Point{X: -1, Y: 0}, "│"),
			},
			renderSize: geo.Size{Width: 7, Height: 7},
		},
		{
			name: "large sprite",
			sprites: []*scene.Sprite{
				scene.NewSprite(geo.Point{X: 0, Y: 0}, "XXX\nXXX\nXXX"),
			},
			renderSize: geo.Size{Width: 5, Height: 5},
		},
		{
			name: "ring as single sprite",
			sprites: []*scene.Sprite{
				scene.NewSprite(geo.Point{X: 0, Y: 0}, "╭─╮\n│ │\n╰─╯"),
			},
			renderSize: geo.Size{Width: 7, Height: 7},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			run(t, tc)
		})
	}
}
