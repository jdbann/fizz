package scene_test

import (
	"testing"

	"github.com/jdbann/fizz/geo"
	"github.com/jdbann/fizz/scene"
	"gotest.tools/v3/assert"
	"gotest.tools/v3/assert/cmp"
)

func TestNewSprite(t *testing.T) {
	t.Run("texture contains runes wider than one cell", func(t *testing.T) {
		assert.Assert(t, cmp.Panics(func() {
			scene.NewSprite(geo.Point{}, "😳")
		}))
	})

	t.Run("texture has inconsistent line widths", func(t *testing.T) {
		assert.Assert(t, cmp.Panics(func() {
			scene.NewSprite(geo.Point{}, "XX\nX")
		}))
	})
}
