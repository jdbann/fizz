package geo_test

import (
	"testing"

	"github.com/jdbann/fizz/geo"
	"gotest.tools/v3/assert"
)

func TestRectAroundPoint(t *testing.T) {
	type testCase struct {
		name     string
		point    geo.Point
		size     geo.Size
		wantRect geo.Rect
	}

	run := func(t *testing.T, tc testCase) {
		assert.DeepEqual(t, geo.RectAroundPoint(tc.point, tc.size), tc.wantRect)
	}

	testCases := []testCase{
		{
			name:  "point around origin",
			point: geo.Point{0, 0},
			size:  geo.Size{5, 5},
			wantRect: geo.Rect{
				Position: geo.Point{-2, -2},
				Size:     geo.Size{5, 5},
			},
		},
		{
			name:  "point away from origin",
			point: geo.Point{10, 10},
			size:  geo.Size{4, 4},
			wantRect: geo.Rect{
				Position: geo.Point{8, 8},
				Size:     geo.Size{4, 4},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			run(t, tc)
		})
	}
}
