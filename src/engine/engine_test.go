package engine

import (
	"testing"

	"github.com/LeMikaelF/2048/src/grid"
)

func Test(t *testing.T) {
	type test struct {
		name       string
		grid       grid.Grid
		move       Direction
		assertions func(t *testing.T, actualGrid grid.Grid, err error)
	}
	tests := []test{
		{
			name: "given a grid with one tile and no slides, it generates a number on an empty square",
			grid: grid.Grid{
				[4]int{0, 0, 0, 0},
				[4]int{2, 0, 0, 0},
				[4]int{0, 0, 0, 0},
				[4]int{0, 0, 0, 0},
			},
			move: Left,
			assertions: func(t *testing.T, actualGrid grid.Grid, err error) {
				assertNil(t, err)
				assertTrue(t, actualGrid.Equals(grid.Grid{
					[4]int{2, 0, 0, 0},
					[4]int{2, 0, 0, 0},
					[4]int{0, 0, 0, 0},
					[4]int{0, 0, 0, 0},
				}))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			engine := NewFromLiteral(tt.grid)
			err := engine.Next(Right)

			tt.assertions(t, engine.Grid, err)
		})
	}
}

func assertNil(t *testing.T, err error) {
	if err != nil {
		t.Errorf("expected nil, was %v", err)
	}
}

func assertTrue(t *testing.T, b bool) {
	t.Helper()
	if !b {
		t.Errorf("expected true, was false")
	}
}
