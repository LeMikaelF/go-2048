package engine

import (
	"math/rand"
	"testing"

	"github.com/LeMikaelF/2048/src/grid"
)

func Test(t *testing.T) {
	type test struct {
		name         string
		grid         grid.Grid
		randomSource rand.Source
		move         Direction
		assertions   func(t *testing.T, actualGrid grid.Grid, err error)
	}
	tests := []test{
		{
			name: "given a grid with one tile and no slides, it generates a number on an empty square (random seed 1)",
			grid: grid.Grid{
				[4]int{0, 0, 0, 0},
				[4]int{2, 0, 0, 0},
				[4]int{0, 0, 0, 0},
				[4]int{0, 0, 0, 0},
			},
			randomSource: rand.NewSource(1),
			move:         Left,
			assertions: func(t *testing.T, actualGrid grid.Grid, err error) {
				assertNil(t, err)
				assertEquals(t, actualGrid, grid.Grid{
					[4]int{0, 0, 0, 0},
					[4]int{2, 0, 0, 0},
					[4]int{0, 0, 0, 0},
					[4]int{0, 0, 2, 0},
				})
			},
		},
		{
			name: "given a grid with one tile and no slides, it generates a number on an empty square (random seed 2)",
			grid: grid.Grid{
				[4]int{0, 0, 0, 0},
				[4]int{2, 0, 0, 0},
				[4]int{0, 0, 0, 0},
				[4]int{0, 0, 0, 0},
			},
			randomSource: rand.NewSource(2),
			move:         Left,
			assertions: func(t *testing.T, actualGrid grid.Grid, err error) {
				assertNil(t, err)
				assertEquals(t, actualGrid, grid.Grid{
					[4]int{0, 0, 0, 0},
					[4]int{2, 2, 0, 0},
					[4]int{0, 0, 0, 0},
					[4]int{0, 0, 0, 0},
				})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			engine := NewFromLiteral(tt.grid, withRandomSource(tt.randomSource))
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

type Equalable[T any] interface {
	Equals(T) bool
}

func assertEquals[LEFT Equalable[RIGHT], RIGHT any](t *testing.T, left LEFT, right RIGHT) {
	equals := left.Equals(right)
	if !equals {
		t.Errorf("wanted equality, but left was\n%v\nand right was\n%v\n", left, right)
	}
}
